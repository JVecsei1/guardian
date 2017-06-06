// Code generated by counterfeiter. DO NOT EDIT.
package gardenerfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/gardener"
)

type FakeSysInfoProvider struct {
	TotalMemoryStub        func() (uint64, error)
	totalMemoryMutex       sync.RWMutex
	totalMemoryArgsForCall []struct{}
	totalMemoryReturns     struct {
		result1 uint64
		result2 error
	}
	totalMemoryReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	TotalDiskStub        func() (uint64, error)
	totalDiskMutex       sync.RWMutex
	totalDiskArgsForCall []struct{}
	totalDiskReturns     struct {
		result1 uint64
		result2 error
	}
	totalDiskReturnsOnCall map[int]struct {
		result1 uint64
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeSysInfoProvider) TotalMemory() (uint64, error) {
	fake.totalMemoryMutex.Lock()
	ret, specificReturn := fake.totalMemoryReturnsOnCall[len(fake.totalMemoryArgsForCall)]
	fake.totalMemoryArgsForCall = append(fake.totalMemoryArgsForCall, struct{}{})
	fake.recordInvocation("TotalMemory", []interface{}{})
	fake.totalMemoryMutex.Unlock()
	if fake.TotalMemoryStub != nil {
		return fake.TotalMemoryStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.totalMemoryReturns.result1, fake.totalMemoryReturns.result2
}

func (fake *FakeSysInfoProvider) TotalMemoryCallCount() int {
	fake.totalMemoryMutex.RLock()
	defer fake.totalMemoryMutex.RUnlock()
	return len(fake.totalMemoryArgsForCall)
}

func (fake *FakeSysInfoProvider) TotalMemoryReturns(result1 uint64, result2 error) {
	fake.TotalMemoryStub = nil
	fake.totalMemoryReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeSysInfoProvider) TotalMemoryReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.TotalMemoryStub = nil
	if fake.totalMemoryReturnsOnCall == nil {
		fake.totalMemoryReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.totalMemoryReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeSysInfoProvider) TotalDisk() (uint64, error) {
	fake.totalDiskMutex.Lock()
	ret, specificReturn := fake.totalDiskReturnsOnCall[len(fake.totalDiskArgsForCall)]
	fake.totalDiskArgsForCall = append(fake.totalDiskArgsForCall, struct{}{})
	fake.recordInvocation("TotalDisk", []interface{}{})
	fake.totalDiskMutex.Unlock()
	if fake.TotalDiskStub != nil {
		return fake.TotalDiskStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.totalDiskReturns.result1, fake.totalDiskReturns.result2
}

func (fake *FakeSysInfoProvider) TotalDiskCallCount() int {
	fake.totalDiskMutex.RLock()
	defer fake.totalDiskMutex.RUnlock()
	return len(fake.totalDiskArgsForCall)
}

func (fake *FakeSysInfoProvider) TotalDiskReturns(result1 uint64, result2 error) {
	fake.TotalDiskStub = nil
	fake.totalDiskReturns = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeSysInfoProvider) TotalDiskReturnsOnCall(i int, result1 uint64, result2 error) {
	fake.TotalDiskStub = nil
	if fake.totalDiskReturnsOnCall == nil {
		fake.totalDiskReturnsOnCall = make(map[int]struct {
			result1 uint64
			result2 error
		})
	}
	fake.totalDiskReturnsOnCall[i] = struct {
		result1 uint64
		result2 error
	}{result1, result2}
}

func (fake *FakeSysInfoProvider) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.totalMemoryMutex.RLock()
	defer fake.totalMemoryMutex.RUnlock()
	fake.totalDiskMutex.RLock()
	defer fake.totalDiskMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeSysInfoProvider) recordInvocation(key string, args []interface{}) {
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

var _ gardener.SysInfoProvider = new(FakeSysInfoProvider)
