// Code generated by counterfeiter. DO NOT EDIT.
package rundmcfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/rundmc"
	"code.cloudfoundry.org/guardian/rundmc/goci"
	"code.cloudfoundry.org/lager"
)

type FakeDepot struct {
	CreateStub        func(log lager.Logger, handle string, bundle goci.Bndl) error
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		log    lager.Logger
		handle string
		bundle goci.Bndl
	}
	createReturns struct {
		result1 error
	}
	createReturnsOnCall map[int]struct {
		result1 error
	}
	LookupStub        func(log lager.Logger, handle string) (path string, err error)
	lookupMutex       sync.RWMutex
	lookupArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	lookupReturns struct {
		result1 string
		result2 error
	}
	lookupReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	DestroyStub        func(log lager.Logger, handle string) error
	destroyMutex       sync.RWMutex
	destroyArgsForCall []struct {
		log    lager.Logger
		handle string
	}
	destroyReturns struct {
		result1 error
	}
	destroyReturnsOnCall map[int]struct {
		result1 error
	}
	HandlesStub        func() ([]string, error)
	handlesMutex       sync.RWMutex
	handlesArgsForCall []struct{}
	handlesReturns     struct {
		result1 []string
		result2 error
	}
	handlesReturnsOnCall map[int]struct {
		result1 []string
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeDepot) Create(log lager.Logger, handle string, bundle goci.Bndl) error {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		log    lager.Logger
		handle string
		bundle goci.Bndl
	}{log, handle, bundle})
	fake.recordInvocation("Create", []interface{}{log, handle, bundle})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(log, handle, bundle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.createReturns.result1
}

func (fake *FakeDepot) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeDepot) CreateArgsForCall(i int) (lager.Logger, string, goci.Bndl) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return fake.createArgsForCall[i].log, fake.createArgsForCall[i].handle, fake.createArgsForCall[i].bundle
}

func (fake *FakeDepot) CreateReturns(result1 error) {
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDepot) CreateReturnsOnCall(i int, result1 error) {
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeDepot) Lookup(log lager.Logger, handle string) (path string, err error) {
	fake.lookupMutex.Lock()
	ret, specificReturn := fake.lookupReturnsOnCall[len(fake.lookupArgsForCall)]
	fake.lookupArgsForCall = append(fake.lookupArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Lookup", []interface{}{log, handle})
	fake.lookupMutex.Unlock()
	if fake.LookupStub != nil {
		return fake.LookupStub(log, handle)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.lookupReturns.result1, fake.lookupReturns.result2
}

func (fake *FakeDepot) LookupCallCount() int {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return len(fake.lookupArgsForCall)
}

func (fake *FakeDepot) LookupArgsForCall(i int) (lager.Logger, string) {
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	return fake.lookupArgsForCall[i].log, fake.lookupArgsForCall[i].handle
}

func (fake *FakeDepot) LookupReturns(result1 string, result2 error) {
	fake.LookupStub = nil
	fake.lookupReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDepot) LookupReturnsOnCall(i int, result1 string, result2 error) {
	fake.LookupStub = nil
	if fake.lookupReturnsOnCall == nil {
		fake.lookupReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.lookupReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeDepot) Destroy(log lager.Logger, handle string) error {
	fake.destroyMutex.Lock()
	ret, specificReturn := fake.destroyReturnsOnCall[len(fake.destroyArgsForCall)]
	fake.destroyArgsForCall = append(fake.destroyArgsForCall, struct {
		log    lager.Logger
		handle string
	}{log, handle})
	fake.recordInvocation("Destroy", []interface{}{log, handle})
	fake.destroyMutex.Unlock()
	if fake.DestroyStub != nil {
		return fake.DestroyStub(log, handle)
	}
	if specificReturn {
		return ret.result1
	}
	return fake.destroyReturns.result1
}

func (fake *FakeDepot) DestroyCallCount() int {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return len(fake.destroyArgsForCall)
}

func (fake *FakeDepot) DestroyArgsForCall(i int) (lager.Logger, string) {
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	return fake.destroyArgsForCall[i].log, fake.destroyArgsForCall[i].handle
}

func (fake *FakeDepot) DestroyReturns(result1 error) {
	fake.DestroyStub = nil
	fake.destroyReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeDepot) DestroyReturnsOnCall(i int, result1 error) {
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

func (fake *FakeDepot) Handles() ([]string, error) {
	fake.handlesMutex.Lock()
	ret, specificReturn := fake.handlesReturnsOnCall[len(fake.handlesArgsForCall)]
	fake.handlesArgsForCall = append(fake.handlesArgsForCall, struct{}{})
	fake.recordInvocation("Handles", []interface{}{})
	fake.handlesMutex.Unlock()
	if fake.HandlesStub != nil {
		return fake.HandlesStub()
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.handlesReturns.result1, fake.handlesReturns.result2
}

func (fake *FakeDepot) HandlesCallCount() int {
	fake.handlesMutex.RLock()
	defer fake.handlesMutex.RUnlock()
	return len(fake.handlesArgsForCall)
}

func (fake *FakeDepot) HandlesReturns(result1 []string, result2 error) {
	fake.HandlesStub = nil
	fake.handlesReturns = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDepot) HandlesReturnsOnCall(i int, result1 []string, result2 error) {
	fake.HandlesStub = nil
	if fake.handlesReturnsOnCall == nil {
		fake.handlesReturnsOnCall = make(map[int]struct {
			result1 []string
			result2 error
		})
	}
	fake.handlesReturnsOnCall[i] = struct {
		result1 []string
		result2 error
	}{result1, result2}
}

func (fake *FakeDepot) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	fake.lookupMutex.RLock()
	defer fake.lookupMutex.RUnlock()
	fake.destroyMutex.RLock()
	defer fake.destroyMutex.RUnlock()
	fake.handlesMutex.RLock()
	defer fake.handlesMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeDepot) recordInvocation(key string, args []interface{}) {
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

var _ rundmc.Depot = new(FakeDepot)
