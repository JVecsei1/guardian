// Code generated by counterfeiter. DO NOT EDIT.
package gardenerfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/lager/v3"
)

type FakeNetworker struct {
	BulkNetOutStub        func(lager.Logger, string, []garden.NetOutRule) error
	bulkNetOutMutex       sync.RWMutex
	bulkNetOutArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 []garden.NetOutRule
	}
	bulkNetOutReturns struct {
		result1 error
	}
	bulkNetOutReturnsOnCall map[int]struct {
		result1 error
	}
	CapacityStub        func() uint64
	capacityMutex       sync.RWMutex
	capacityArgsForCall []struct {
	}
	capacityReturns struct {
		result1 uint64
	}
	capacityReturnsOnCall map[int]struct {
		result1 uint64
	}
	DestroyStub        func(lager.Logger, string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	NetInStub        func(lager.Logger, string, uint32, uint32) (uint32, uint32, error)
	netInMutex       sync.RWMutex
	netInArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 uint32
		arg4 uint32
	}
	netInReturns struct {
		result1 uint32
		result2 uint32
		result3 error
	}
	netInReturnsOnCall map[int]struct {
		result1 uint32
		result2 uint32
		result3 error
	}
	NetOutStub        func(lager.Logger, string, garden.NetOutRule) error
	netOutMutex       sync.RWMutex
	netOutArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 garden.NetOutRule
	}
	netOutReturns struct {
		result1 error
	}
	netOutReturnsOnCall map[int]struct {
		result1 error
	}
	NetworkStub        func(lager.Logger, garden.ContainerSpec, int) error
	networkMutex       sync.RWMutex
	networkArgsForCall []struct {
		arg1 lager.Logger
		arg2 garden.ContainerSpec
		arg3 int
	}
	networkReturns struct {
		result1 error
	}
	networkReturnsOnCall map[int]struct {
		result1 error
	}
	RestoreStub        func(lager.Logger, string) error
	restoreMutex       sync.RWMutex
	restoreArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	restoreReturns struct {
		result1 error
	}
	restoreReturnsOnCall map[int]struct {
		result1 error
	}
	SetupBindMountsStub        func(lager.Logger, string, bool, string) ([]garden.BindMount, error)
	setupBindMountsMutex       sync.RWMutex
	setupBindMountsArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 bool
		arg4 string
	}
	setupBindMountsReturns struct {
		result1 []garden.BindMount
		result2 error
	}
	setupBindMountsReturnsOnCall map[int]struct {
		result1 []garden.BindMount
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeNetworker) BulkNetOut(arg1 lager.Logger, arg2 string, arg3 []garden.NetOutRule) error {
	var arg3Copy []garden.NetOutRule
	if arg3 != nil {
		arg3Copy = make([]garden.NetOutRule, len(arg3))
		copy(arg3Copy, arg3)
	}
	fake.bulkNetOutMutex.Lock()
	ret, specificReturn := fake.bulkNetOutReturnsOnCall[len(fake.bulkNetOutArgsForCall)]
	fake.bulkNetOutArgsForCall = append(fake.bulkNetOutArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 []garden.NetOutRule
	}{arg1, arg2, arg3Copy})
	fake.recordInvocation("BulkNetOut", []interface{}{arg1, arg2, arg3Copy})
	fake.bulkNetOutMutex.Unlock()
	if fake.BulkNetOutStub != nil {
		return fake.BulkNetOutStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.bulkNetOutReturns
	return fakeReturns.result1
}

func (fake *FakeNetworker) BulkNetOutCallCount() int {
	fake.bulkNetOutMutex.RLock()
	defer fake.bulkNetOutMutex.RUnlock()
	return len(fake.bulkNetOutArgsForCall)
}

func (fake *FakeNetworker) BulkNetOutCalls(stub func(lager.Logger, string, []garden.NetOutRule) error) {
	fake.bulkNetOutMutex.Lock()
	defer fake.bulkNetOutMutex.Unlock()
	fake.BulkNetOutStub = stub
}

func (fake *FakeNetworker) BulkNetOutArgsForCall(i int) (lager.Logger, string, []garden.NetOutRule) {
	fake.bulkNetOutMutex.RLock()
	defer fake.bulkNetOutMutex.RUnlock()
	argsForCall := fake.bulkNetOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNetworker) BulkNetOutReturns(result1 error) {
	fake.bulkNetOutMutex.Lock()
	defer fake.bulkNetOutMutex.Unlock()
	fake.BulkNetOutStub = nil
	fake.bulkNetOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) BulkNetOutReturnsOnCall(i int, result1 error) {
	fake.bulkNetOutMutex.Lock()
	defer fake.bulkNetOutMutex.Unlock()
	fake.BulkNetOutStub = nil
	if fake.bulkNetOutReturnsOnCall == nil {
		fake.bulkNetOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.bulkNetOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) Capacity() uint64 {
	fake.capacityMutex.Lock()
	ret, specificReturn := fake.capacityReturnsOnCall[len(fake.capacityArgsForCall)]
	fake.capacityArgsForCall = append(fake.capacityArgsForCall, struct {
	}{})
	fake.recordInvocation("Capacity", []interface{}{})
	fake.capacityMutex.Unlock()
	if fake.CapacityStub != nil {
		return fake.CapacityStub()
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.capacityReturns
	return fakeReturns.result1
}

func (fake *FakeNetworker) CapacityCallCount() int {
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	return len(fake.capacityArgsForCall)
}

func (fake *FakeNetworker) CapacityCalls(stub func() uint64) {
	fake.capacityMutex.Lock()
	defer fake.capacityMutex.Unlock()
	fake.CapacityStub = stub
}

func (fake *FakeNetworker) CapacityReturns(result1 uint64) {
	fake.capacityMutex.Lock()
	defer fake.capacityMutex.Unlock()
	fake.CapacityStub = nil
	fake.capacityReturns = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeNetworker) CapacityReturnsOnCall(i int, result1 uint64) {
	fake.capacityMutex.Lock()
	defer fake.capacityMutex.Unlock()
	fake.CapacityStub = nil
	if fake.capacityReturnsOnCall == nil {
		fake.capacityReturnsOnCall = make(map[int]struct {
			result1 uint64
		})
	}
	fake.capacityReturnsOnCall[i] = struct {
		result1 uint64
	}{result1}
}

func (fake *FakeNetworker) Destroy(arg1 lager.Logger, arg2 string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Destroy", []interface{}{arg1, arg2})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.destroyReturns
	return fakeReturns.result1
}

func (fake *FakeNetworker) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeNetworker) DestroyCalls(stub func(lager.Logger, string) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeNetworker) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetworker) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) DestroyReturnsOnCall(i int, result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	if fake.destroyReturnsOnCall == nil {
		fake.destroyReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.destroyReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) NetIn(arg1 lager.Logger, arg2 string, arg3 uint32, arg4 uint32) (uint32, uint32, error) {
	fake.netInMutex.Lock()
	ret, specificReturn := fake.netInReturnsOnCall[len(fake.netInArgsForCall)]
	fake.netInArgsForCall = append(fake.netInArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 uint32
		arg4 uint32
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("NetIn", []interface{}{arg1, arg2, arg3, arg4})
	fake.netInMutex.Unlock()
	if fake.NetInStub != nil {
		return fake.NetInStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	fakeReturns := fake.netInReturns
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeNetworker) NetInCallCount() int {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	return len(fake.netInArgsForCall)
}

func (fake *FakeNetworker) NetInCalls(stub func(lager.Logger, string, uint32, uint32) (uint32, uint32, error)) {
	fake.netInMutex.Lock()
	defer fake.netInMutex.Unlock()
	fake.NetInStub = stub
}

func (fake *FakeNetworker) NetInArgsForCall(i int) (lager.Logger, string, uint32, uint32) {
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	argsForCall := fake.netInArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeNetworker) NetInReturns(result1 uint32, result2 uint32, result3 error) {
	fake.netInMutex.Lock()
	defer fake.netInMutex.Unlock()
	fake.NetInStub = nil
	fake.netInReturns = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworker) NetInReturnsOnCall(i int, result1 uint32, result2 uint32, result3 error) {
	fake.netInMutex.Lock()
	defer fake.netInMutex.Unlock()
	fake.NetInStub = nil
	if fake.netInReturnsOnCall == nil {
		fake.netInReturnsOnCall = make(map[int]struct {
			result1 uint32
			result2 uint32
			result3 error
		})
	}
	fake.netInReturnsOnCall[i] = struct {
		result1 uint32
		result2 uint32
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeNetworker) NetOut(arg1 lager.Logger, arg2 string, arg3 garden.NetOutRule) error {
	fake.netOutMutex.Lock()
	ret, specificReturn := fake.netOutReturnsOnCall[len(fake.netOutArgsForCall)]
	fake.netOutArgsForCall = append(fake.netOutArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 garden.NetOutRule
	}{arg1, arg2, arg3})
	fake.recordInvocation("NetOut", []interface{}{arg1, arg2, arg3})
	fake.netOutMutex.Unlock()
	if fake.NetOutStub != nil {
		return fake.NetOutStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.netOutReturns
	return fakeReturns.result1
}

func (fake *FakeNetworker) NetOutCallCount() int {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	return len(fake.netOutArgsForCall)
}

func (fake *FakeNetworker) NetOutCalls(stub func(lager.Logger, string, garden.NetOutRule) error) {
	fake.netOutMutex.Lock()
	defer fake.netOutMutex.Unlock()
	fake.NetOutStub = stub
}

func (fake *FakeNetworker) NetOutArgsForCall(i int) (lager.Logger, string, garden.NetOutRule) {
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	argsForCall := fake.netOutArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNetworker) NetOutReturns(result1 error) {
	fake.netOutMutex.Lock()
	defer fake.netOutMutex.Unlock()
	fake.NetOutStub = nil
	fake.netOutReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) NetOutReturnsOnCall(i int, result1 error) {
	fake.netOutMutex.Lock()
	defer fake.netOutMutex.Unlock()
	fake.NetOutStub = nil
	if fake.netOutReturnsOnCall == nil {
		fake.netOutReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.netOutReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) Network(arg1 lager.Logger, arg2 garden.ContainerSpec, arg3 int) error {
	fake.networkMutex.Lock()
	ret, specificReturn := fake.networkReturnsOnCall[len(fake.networkArgsForCall)]
	fake.networkArgsForCall = append(fake.networkArgsForCall, struct {
		arg1 lager.Logger
		arg2 garden.ContainerSpec
		arg3 int
	}{arg1, arg2, arg3})
	fake.recordInvocation("Network", []interface{}{arg1, arg2, arg3})
	fake.networkMutex.Unlock()
	if fake.NetworkStub != nil {
		return fake.NetworkStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.networkReturns
	return fakeReturns.result1
}

func (fake *FakeNetworker) NetworkCallCount() int {
	fake.networkMutex.RLock()
	defer fake.networkMutex.RUnlock()
	return len(fake.networkArgsForCall)
}

func (fake *FakeNetworker) NetworkCalls(stub func(lager.Logger, garden.ContainerSpec, int) error) {
	fake.networkMutex.Lock()
	defer fake.networkMutex.Unlock()
	fake.NetworkStub = stub
}

func (fake *FakeNetworker) NetworkArgsForCall(i int) (lager.Logger, garden.ContainerSpec, int) {
	fake.networkMutex.RLock()
	defer fake.networkMutex.RUnlock()
	argsForCall := fake.networkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeNetworker) NetworkReturns(result1 error) {
	fake.networkMutex.Lock()
	defer fake.networkMutex.Unlock()
	fake.NetworkStub = nil
	fake.networkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) NetworkReturnsOnCall(i int, result1 error) {
	fake.networkMutex.Lock()
	defer fake.networkMutex.Unlock()
	fake.NetworkStub = nil
	if fake.networkReturnsOnCall == nil {
		fake.networkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.networkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) Restore(arg1 lager.Logger, arg2 string) error {
	fake.restoreMutex.Lock()
	ret, specificReturn := fake.restoreReturnsOnCall[len(fake.restoreArgsForCall)]
	fake.restoreArgsForCall = append(fake.restoreArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("Restore", []interface{}{arg1, arg2})
	fake.restoreMutex.Unlock()
	if fake.RestoreStub != nil {
		return fake.RestoreStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	fakeReturns := fake.restoreReturns
	return fakeReturns.result1
}

func (fake *FakeNetworker) RestoreCallCount() int {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	return len(fake.restoreArgsForCall)
}

func (fake *FakeNetworker) RestoreCalls(stub func(lager.Logger, string) error) {
	fake.restoreMutex.Lock()
	defer fake.restoreMutex.Unlock()
	fake.RestoreStub = stub
}

func (fake *FakeNetworker) RestoreArgsForCall(i int) (lager.Logger, string) {
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	argsForCall := fake.restoreArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetworker) RestoreReturns(result1 error) {
	fake.restoreMutex.Lock()
	defer fake.restoreMutex.Unlock()
	fake.RestoreStub = nil
	fake.restoreReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) RestoreReturnsOnCall(i int, result1 error) {
	fake.restoreMutex.Lock()
	defer fake.restoreMutex.Unlock()
	fake.RestoreStub = nil
	if fake.restoreReturnsOnCall == nil {
		fake.restoreReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.restoreReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworker) SetupBindMounts(arg1 lager.Logger, arg2 string, arg3 bool, arg4 string) ([]garden.BindMount, error) {
	fake.setupBindMountsMutex.Lock()
	ret, specificReturn := fake.setupBindMountsReturnsOnCall[len(fake.setupBindMountsArgsForCall)]
	fake.setupBindMountsArgsForCall = append(fake.setupBindMountsArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 bool
		arg4 string
	}{arg1, arg2, arg3, arg4})
	fake.recordInvocation("SetupBindMounts", []interface{}{arg1, arg2, arg3, arg4})
	fake.setupBindMountsMutex.Unlock()
	if fake.SetupBindMountsStub != nil {
		return fake.SetupBindMountsStub(arg1, arg2, arg3, arg4)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.setupBindMountsReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeNetworker) SetupBindMountsCallCount() int {
	fake.setupBindMountsMutex.RLock()
	defer fake.setupBindMountsMutex.RUnlock()
	return len(fake.setupBindMountsArgsForCall)
}

func (fake *FakeNetworker) SetupBindMountsCalls(stub func(lager.Logger, string, bool, string) ([]garden.BindMount, error)) {
	fake.setupBindMountsMutex.Lock()
	defer fake.setupBindMountsMutex.Unlock()
	fake.SetupBindMountsStub = stub
}

func (fake *FakeNetworker) SetupBindMountsArgsForCall(i int) (lager.Logger, string, bool, string) {
	fake.setupBindMountsMutex.RLock()
	defer fake.setupBindMountsMutex.RUnlock()
	argsForCall := fake.setupBindMountsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeNetworker) SetupBindMountsReturns(result1 []garden.BindMount, result2 error) {
	fake.setupBindMountsMutex.Lock()
	defer fake.setupBindMountsMutex.Unlock()
	fake.SetupBindMountsStub = nil
	fake.setupBindMountsReturns = struct {
		result1 []garden.BindMount
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworker) SetupBindMountsReturnsOnCall(i int, result1 []garden.BindMount, result2 error) {
	fake.setupBindMountsMutex.Lock()
	defer fake.setupBindMountsMutex.Unlock()
	fake.SetupBindMountsStub = nil
	if fake.setupBindMountsReturnsOnCall == nil {
		fake.setupBindMountsReturnsOnCall = make(map[int]struct {
			result1 []garden.BindMount
			result2 error
		})
	}
	fake.setupBindMountsReturnsOnCall[i] = struct {
		result1 []garden.BindMount
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworker) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.bulkNetOutMutex.RLock()
	defer fake.bulkNetOutMutex.RUnlock()
	fake.capacityMutex.RLock()
	defer fake.capacityMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.netInMutex.RLock()
	defer fake.netInMutex.RUnlock()
	fake.netOutMutex.RLock()
	defer fake.netOutMutex.RUnlock()
	fake.networkMutex.RLock()
	defer fake.networkMutex.RUnlock()
	fake.restoreMutex.RLock()
	defer fake.restoreMutex.RUnlock()
	fake.setupBindMountsMutex.RLock()
	defer fake.setupBindMountsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetworker) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ gardener.Networker = new(FakeNetworker)
