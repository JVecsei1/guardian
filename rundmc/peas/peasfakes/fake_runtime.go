// Code generated by counterfeiter. DO NOT EDIT.
package peasfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc/peas"
	"code.cloudfoundry.org/lager"
)

type FakeRuntime struct {
	ContainerHandlesStub        func() ([]string, error)
	containerHandlesMutex       sync.RWMutex
	containerHandlesArgsForCall []struct {
	}
	containerHandlesReturns struct {
		result1 []string
		result2 error
	}
	containerHandlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	ContainerPeaHandlesStub        func(lager.Logger, string) ([]string, error)
	containerPeaHandlesMutex       sync.RWMutex
	containerPeaHandlesArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
	}
	containerPeaHandlesReturns struct {
		result1 []string
		result2 error
	}
	containerPeaHandlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeRuntime) ContainerHandles() ([]string, error) {
	fake.containerHandlesMutex.Lock()
	ret, specificReturn := fake.containerHandlesReturnsOnCall[len(fake.containerHandlesArgsForCall)]
	fake.containerHandlesArgsForCall = append(fake.containerHandlesArgsForCall, struct {
	}{})
	fake.recordInvocation("ContainerHandles", []interface{}{})
	fake.containerHandlesMutex.Unlock()
	if fake.ContainerHandlesStub != nil {
		return fake.ContainerHandlesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.containerHandlesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRuntime) ContainerHandlesCallCount() int {
	fake.containerHandlesMutex.RLock()
	defer fake.containerHandlesMutex.RUnlock()
	return len(fake.containerHandlesArgsForCall)
}

func (fake *FakeRuntime) ContainerHandlesCalls(stub func() ([]string, error)) {
	fake.containerHandlesMutex.Lock()
	defer fake.containerHandlesMutex.Unlock()
	fake.ContainerHandlesStub = stub
}

func (fake *FakeRuntime) ContainerHandlesReturns(result1 []string, result2 error) {
	fake.containerHandlesMutex.Lock()
	defer fake.containerHandlesMutex.Unlock()
	fake.ContainerHandlesStub = nil
	fake.containerHandlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntime) ContainerHandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.containerHandlesMutex.Lock()
	defer fake.containerHandlesMutex.Unlock()
	fake.ContainerHandlesStub = nil
	if fake.containerHandlesReturnsOnCall == nil {
		fake.containerHandlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.containerHandlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntime) ContainerPeaHandles(arg1 lager.Logger, arg2 string) ([]string, error) {
	fake.containerPeaHandlesMutex.Lock()
	ret, specificReturn := fake.containerPeaHandlesReturnsOnCall[len(fake.containerPeaHandlesArgsForCall)]
	fake.containerPeaHandlesArgsForCall = append(fake.containerPeaHandlesArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
	}{arg1, arg2})
	fake.recordInvocation("ContainerPeaHandles", []interface{}{arg1, arg2})
	fake.containerPeaHandlesMutex.Unlock()
	if fake.ContainerPeaHandlesStub != nil {
		return fake.ContainerPeaHandlesStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.containerPeaHandlesReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeRuntime) ContainerPeaHandlesCallCount() int {
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	return len(fake.containerPeaHandlesArgsForCall)
}

func (fake *FakeRuntime) ContainerPeaHandlesCalls(stub func(lager.Logger, string) ([]string, error)) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = stub
}

func (fake *FakeRuntime) ContainerPeaHandlesArgsForCall(i int) (lager.Logger, string) {
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	argsForCall := fake.containerPeaHandlesArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeRuntime) ContainerPeaHandlesReturns(result1 []string, result2 error) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = nil
	fake.containerPeaHandlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntime) ContainerPeaHandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.containerPeaHandlesMutex.Lock()
	defer fake.containerPeaHandlesMutex.Unlock()
	fake.ContainerPeaHandlesStub = nil
	if fake.containerPeaHandlesReturnsOnCall == nil {
		fake.containerPeaHandlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.containerPeaHandlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeRuntime) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.containerHandlesMutex.RLock()
	defer fake.containerHandlesMutex.RUnlock()
	fake.containerPeaHandlesMutex.RLock()
	defer fake.containerPeaHandlesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeRuntime) recordInvocation(key string, args []interface{}) {
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

var _ peas.Runtime = new(FakeRuntime)