package gardener

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/lager/v3"
)

type container struct {
	logger lager.Logger

	handle          string
	containerizer   Containerizer
	volumizer       Volumizer
	networker       Networker
	propertyManager PropertyManager
}

func (c *container) Handle() string {
	return c.handle
}

func (c *container) Run(spec garden.ProcessSpec, io garden.ProcessIO) (garden.Process, error) {
	return c.containerizer.Run(c.logger, c.handle, spec, io)
}

func (c *container) Attach(processID string, io garden.ProcessIO) (garden.Process, error) {
	return c.containerizer.Attach(c.logger, c.handle, processID, io)
}

func (c *container) Stop(kill bool) error {
	return c.containerizer.Stop(c.logger, c.handle, kill)
}

func (c *container) Info() (garden.ContainerInfo, error) {
	log := c.logger.Session("info", lager.Data{"handle": c.handle})

	log.Debug("starting")
	defer log.Debug("finished")

	containerIP, _ := c.propertyManager.Get(c.handle, ContainerIPKey)
	hostIP, _ := c.propertyManager.Get(c.handle, BridgeIPKey)
	externalIP, _ := c.propertyManager.Get(c.handle, ExternalIPKey)

	actualContainerSpec, err := c.containerizer.Info(c.logger, c.handle)
	if err != nil {
		return garden.ContainerInfo{}, err
	}

	properties, err := c.propertyManager.All(c.handle)
	if err != nil {
		return garden.ContainerInfo{}, err
	}

	mappedPorts := []garden.PortMapping{}
	mappedPortsCfg, _ := c.propertyManager.Get(c.handle, MappedPortsKey)

	state := "active"
	if actualContainerSpec.Stopped {
		state = "stopped"
	}

	json.Unmarshal([]byte(mappedPortsCfg), &mappedPorts)
	return garden.ContainerInfo{
		State:         state,
		ContainerIP:   containerIP,
		HostIP:        hostIP,
		ExternalIP:    externalIP,
		ContainerPath: actualContainerSpec.BundlePath,
		Events:        actualContainerSpec.Events,
		Properties:    properties,
		MappedPorts:   mappedPorts,
	}, nil
}

func (c *container) StreamIn(spec garden.StreamInSpec) error {
	return c.containerizer.StreamIn(c.logger, c.handle, spec)
}

func (c *container) StreamOut(spec garden.StreamOutSpec) (io.ReadCloser, error) {
	return c.containerizer.StreamOut(c.logger, c.handle, spec)
}

func (c *container) LimitBandwidth(limits garden.BandwidthLimits) error {
	return nil
}

func (c *container) CurrentBandwidthLimits() (garden.BandwidthLimits, error) {
	return garden.BandwidthLimits{}, nil
}

func (c *container) LimitCPU(limits garden.CPULimits) error {
	return nil
}

func (c *container) CurrentCPULimits() (garden.CPULimits, error) {
	info, err := c.containerizer.Info(c.logger, c.handle)
	return info.Limits.CPU, err
}

func (c *container) LimitDisk(limits garden.DiskLimits) error {
	return nil
}

func (c *container) CurrentDiskLimits() (garden.DiskLimits, error) {
	return garden.DiskLimits{}, nil
}

func (c *container) LimitMemory(limits garden.MemoryLimits) error {
	return nil
}

func (c *container) CurrentMemoryLimits() (garden.MemoryLimits, error) {
	info, err := c.containerizer.Info(c.logger, c.handle)
	return info.Limits.Memory, err
}

func (c *container) NetIn(hostPort, containerPort uint32) (uint32, uint32, error) {
	return c.networker.NetIn(c.logger, c.handle, hostPort, containerPort)
}

func (c *container) NetOut(netOutRule garden.NetOutRule) error {
	return c.networker.NetOut(c.logger, c.handle, netOutRule)
}

func (c *container) BulkNetOut(netOutRules []garden.NetOutRule) error {
	return c.networker.BulkNetOut(c.logger, c.handle, netOutRules)
}

func (c *container) Metrics() (garden.Metrics, error) {
	actualContainerMetrics, err := c.containerizer.Metrics(c.logger, c.handle)
	if err != nil {
		return garden.Metrics{}, err
	}

	diskMetrics, err1 := c.volumizer.Metrics(c.logger, c.handle, true)
	if err1 != nil {
		diskMetrics, err = c.volumizer.Metrics(c.logger, c.handle, false)
		if err != nil {
			return garden.Metrics{}, fmt.Errorf("image plugin returned these errors:\nunprivileged: %s\nprivileged: %s", err1.Error(), err.Error())
		}
	}

	networkStat, err := c.containerNetworkStat()
	if err != nil {
		return garden.Metrics{}, fmt.Errorf("could not read container network statistics, %w", err)
	}

	return garden.Metrics{
		CPUStat:        actualContainerMetrics.CPU,
		MemoryStat:     actualContainerMetrics.Memory,
		DiskStat:       diskMetrics,
		PidStat:        actualContainerMetrics.Pid,
		Age:            actualContainerMetrics.Age,
		CPUEntitlement: actualContainerMetrics.CPUEntitlement,
		NetworkStat:    networkStat,
	}, nil
}

func (c *container) containerNetworkStat() (garden.ContainerNetworkStat, error) {
	ifName, err := c.Property(ContainerInterfaceKey)
	if err != nil {
		return garden.ContainerNetworkStat{}, err
	}

	stdout := new(bytes.Buffer)
	stderr := new(bytes.Buffer)

	process, err := c.Run(garden.ProcessSpec{
		Path: "cat",
		Args: []string{
			networkStatPath(ifName, "rx_bytes"),
			networkStatPath(ifName, "tx_bytes"),
		},
	}, garden.ProcessIO{
		Stdout: stdout,
		Stderr: stderr,
	})

	if err != nil {
		return garden.ContainerNetworkStat{}, fmt.Errorf("running process failed, %w", err)
	}

	exitStatus, err := process.Wait()
	if err != nil {
		return garden.ContainerNetworkStat{}, err
	}

	if exitStatus != 0 {
		return garden.ContainerNetworkStat{}, fmt.Errorf("running process failed with exit status %d, error %q", exitStatus, stderr.String())
	}

	stats := strings.Split(strings.TrimSpace(stdout.String()), "\n")
	if len(stats) != 2 {
		return garden.ContainerNetworkStat{}, fmt.Errorf("expected two values but got %q", stdout.String())
	}

	for idx, s := range stats {
		stats[idx] = strings.TrimSpace(s)
	}

	rxBytes, err := strconv.ParseUint(stats[0], 10, 64)
	if err != nil {
		return garden.ContainerNetworkStat{}, fmt.Errorf("could not parse rx_bytes value %q, %w", stats[0], err)
	}

	txBytes, err := strconv.ParseUint(stats[1], 10, 64)
	if err != nil {
		return garden.ContainerNetworkStat{}, fmt.Errorf("could not parse tx_bytes value %q, %w", stats[1], err)
	}

	return garden.ContainerNetworkStat{
		RxBytes: rxBytes,
		TxBytes: txBytes,
	}, nil
}

func networkStatPath(ifName, stat string) string {
	return filepath.Join("/sys/class/net", ifName, "statistics", stat)
}

func (c *container) Properties() (garden.Properties, error) {
	return c.propertyManager.All(c.handle)
}

func (c *container) Property(name string) (string, error) {
	if prop, ok := c.propertyManager.Get(c.handle, name); ok {
		return prop, nil
	}

	return "", fmt.Errorf("property does not exist: %s", name)
}

func (c *container) SetProperty(name string, value string) error {
	c.propertyManager.Set(c.handle, name, value)
	return nil
}

func (c *container) RemoveProperty(name string) error {
	c.propertyManager.Remove(c.handle, name)
	return nil
}

func (c *container) SetGraceTime(t time.Duration) error {
	c.propertyManager.Set(c.handle, GraceTimeKey, fmt.Sprintf("%d", t))
	return nil
}
