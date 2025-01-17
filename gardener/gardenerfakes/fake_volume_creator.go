// Code generated by counterfeiter. DO NOT EDIT.
package gardenerfakes

import (
	"sync"

	"code.cloudfoundry.org/guardian/gardener"
	"code.cloudfoundry.org/lager/v3"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

type FakeVolumeCreator struct {
	CreateStub        func(lager.Logger, string, gardener.RootfsSpec) (specs.Spec, error)
	createMutex       sync.RWMutex
	createArgsForCall []struct {
		arg1 lager.Logger
		arg2 string
		arg3 gardener.RootfsSpec
	}
	createReturns struct {
		result1 specs.Spec
		result2 error
	}
	createReturnsOnCall map[int]struct {
		result1 specs.Spec
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeVolumeCreator) Create(arg1 lager.Logger, arg2 string, arg3 gardener.RootfsSpec) (specs.Spec, error) {
	fake.createMutex.Lock()
	ret, specificReturn := fake.createReturnsOnCall[len(fake.createArgsForCall)]
	fake.createArgsForCall = append(fake.createArgsForCall, struct {
		arg1 lager.Logger
		arg2 string
		arg3 gardener.RootfsSpec
	}{arg1, arg2, arg3})
	fake.recordInvocation("Create", []interface{}{arg1, arg2, arg3})
	fake.createMutex.Unlock()
	if fake.CreateStub != nil {
		return fake.CreateStub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	fakeReturns := fake.createReturns
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeVolumeCreator) CreateCallCount() int {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	return len(fake.createArgsForCall)
}

func (fake *FakeVolumeCreator) CreateCalls(stub func(lager.Logger, string, gardener.RootfsSpec) (specs.Spec, error)) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = stub
}

func (fake *FakeVolumeCreator) CreateArgsForCall(i int) (lager.Logger, string, gardener.RootfsSpec) {
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	argsForCall := fake.createArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeVolumeCreator) CreateReturns(result1 specs.Spec, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	fake.createReturns = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeCreator) CreateReturnsOnCall(i int, result1 specs.Spec, result2 error) {
	fake.createMutex.Lock()
	defer fake.createMutex.Unlock()
	fake.CreateStub = nil
	if fake.createReturnsOnCall == nil {
		fake.createReturnsOnCall = make(map[int]struct {
			result1 specs.Spec
			result2 error
		})
	}
	fake.createReturnsOnCall[i] = struct {
		result1 specs.Spec
		result2 error
	}{result1, result2}
}

func (fake *FakeVolumeCreator) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.createMutex.RLock()
	defer fake.createMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeVolumeCreator) recordInvocation(key string, args []interface{}) {
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

var _ gardener.VolumeCreator = new(FakeVolumeCreator)
