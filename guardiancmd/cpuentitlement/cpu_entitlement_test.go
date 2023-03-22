package cpuentitlement_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"code.cloudfoundry.org/guardian/gardener/gardenerfakes"
	"code.cloudfoundry.org/guardian/guardiancmd/cpuentitlement"
)

var _ = Describe("CPU Entitlement Calculator", func() {
	var (
		sysInfoProvider *gardenerfakes.FakeSysInfoProvider
		calculator      cpuentitlement.Calculator
	)

	BeforeEach(func() {
		sysInfoProvider = new(gardenerfakes.FakeSysInfoProvider)
		sysInfoProvider.CPUCoresReturns(4, nil)
		sysInfoProvider.TotalMemoryReturns(uint64(1024*1024*1024), nil)

		calculator = cpuentitlement.Calculator{SysInfoProvider: sysInfoProvider}
	})

	Describe("CalculateDefaultEntitlementPerShare", func() {
		var (
			defaultEntitlement float64
			calculateErr       error
		)

		JustBeforeEach(func() {
			defaultEntitlement, calculateErr = calculator.CalculateDefaultEntitlementPerShare()
		})

		It("succeeds", func() {
			Expect(calculateErr).NotTo(HaveOccurred())
		})

		It("calculates the default cpu entitlement per share", func() {
			Expect(defaultEntitlement).To(BeNumerically("~", 0.39, 0.01))
		})

		Context("when getting CPU cores fails", func() {
			BeforeEach(func() {
				sysInfoProvider.CPUCoresReturns(0, errors.New("cpu-cores"))
			})

			It("returns the error", func() {
				Expect(calculateErr).To(MatchError("cpu-cores"))
			})
		})

		Context("when getting memory fails", func() {
			BeforeEach(func() {
				sysInfoProvider.TotalMemoryReturns(0, errors.New("total-memory"))
			})

			It("returns the error", func() {
				Expect(calculateErr).To(MatchError("total-memory"))
			})
		})
	})

	Describe("CalculateEntitlementMultiplier", func() {
		It("calculates the multiplier, given a custom CPU entitlement value", func() {
			Expect(calculator.CalculateEntitlementMultiplier(0.195)).To(BeNumerically("~", 0.5, 0.01))
		})
	})
})
