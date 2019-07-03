// Code generated by counterfeiter. DO NOT EDIT.
package peasfakes

import (
	"sync"

	"code.cloudfoundry.org/garden"
	"code.cloudfoundry.org/guardian/rundmc/peas"
	"code.cloudfoundry.org/lager"
)

type FakeNetworkDepot struct {
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

func (fake *FakeNetworkDepot) Destroy(arg1 lager.Logger, arg2 string) error {
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

func (fake *FakeNetworkDepot) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeNetworkDepot) DestroyCalls(stub func(lager.Logger, string) error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = stub
}

func (fake *FakeNetworkDepot) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	argsForCall := fake.destroyArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeNetworkDepot) DestroyReturns(result1 error) {
	fake.destroyMutex.Lock()
	defer fake.destroyMutex.Unlock()
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeNetworkDepot) DestroyReturnsOnCall(i int, result1 error) {
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

func (fake *FakeNetworkDepot) SetupBindMounts(arg1 lager.Logger, arg2 string, arg3 bool, arg4 string) ([]garden.BindMount, error) {
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

func (fake *FakeNetworkDepot) SetupBindMountsCallCount() int {
	fake.setupBindMountsMutex.RLock()
	defer fake.setupBindMountsMutex.RUnlock()
	return len(fake.setupBindMountsArgsForCall)
}

func (fake *FakeNetworkDepot) SetupBindMountsCalls(stub func(lager.Logger, string, bool, string) ([]garden.BindMount, error)) {
	fake.setupBindMountsMutex.Lock()
	defer fake.setupBindMountsMutex.Unlock()
	fake.SetupBindMountsStub = stub
}

func (fake *FakeNetworkDepot) SetupBindMountsArgsForCall(i int) (lager.Logger, string, bool, string) {
	fake.setupBindMountsMutex.RLock()
	defer fake.setupBindMountsMutex.RUnlock()
	argsForCall := fake.setupBindMountsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3, argsForCall.arg4
}

func (fake *FakeNetworkDepot) SetupBindMountsReturns(result1 []garden.BindMount, result2 error) {
	fake.setupBindMountsMutex.Lock()
	defer fake.setupBindMountsMutex.Unlock()
	fake.SetupBindMountsStub = nil
	fake.setupBindMountsReturns = struct {
		result1 []garden.BindMount
		result2 error
	}{result1, result2}
}

func (fake *FakeNetworkDepot) SetupBindMountsReturnsOnCall(i int, result1 []garden.BindMount, result2 error) {
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

func (fake *FakeNetworkDepot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.setupBindMountsMutex.RLock()
	defer fake.setupBindMountsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeNetworkDepot) recordInvocation(key string, args []interface{}) {
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

var _ peas.NetworkDepot = new(FakeNetworkDepot)