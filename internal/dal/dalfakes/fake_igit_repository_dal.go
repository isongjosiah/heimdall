// Code generated by counterfeiter. DO NOT EDIT.
package dalfakes

import (
	"context"
	"heimdall/internal/dal"
	"heimdall/internal/dal/model"
	"sync"
)

type FakeIGitRepositoryDAL struct {
	AddRepoStub        func(context.Context, model.GitRepository) error
	addRepoMutex       sync.RWMutex
	addRepoArgsForCall []struct {
		arg1 context.Context
		arg2 model.GitRepository
	}
	addRepoReturns struct {
		result1 error
	}
	addRepoReturnsOnCall map[int]struct {
		result1 error
	}
	ListRepoCursorStub        func(context.Context, string, int) ([]model.GitRepository, error)
	listRepoCursorMutex       sync.RWMutex
	listRepoCursorArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 int
	}
	listRepoCursorReturns struct {
		result1 []model.GitRepository
		result2 error
	}
	listRepoCursorReturnsOnCall map[int]struct {
		result1 []model.GitRepository
		result2 error
	}
	RepoByNameStub        func(context.Context, string) (model.GitRepository, string, error)
	repoByNameMutex       sync.RWMutex
	repoByNameArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	repoByNameReturns struct {
		result1 model.GitRepository
		result2 string
		result3 error
	}
	repoByNameReturnsOnCall map[int]struct {
		result1 model.GitRepository
		result2 string
		result3 error
	}
	RepoExistsStub        func(context.Context, string) (bool, error)
	repoExistsMutex       sync.RWMutex
	repoExistsArgsForCall []struct {
		arg1 context.Context
		arg2 string
	}
	repoExistsReturns struct {
		result1 bool
		result2 error
	}
	repoExistsReturnsOnCall map[int]struct {
		result1 bool
		result2 error
	}
	UpdateRepoStub        func(context.Context, string, map[string]any) error
	updateRepoMutex       sync.RWMutex
	updateRepoArgsForCall []struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]any
	}
	updateRepoReturns struct {
		result1 error
	}
	updateRepoReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeIGitRepositoryDAL) AddRepo(arg1 context.Context, arg2 model.GitRepository) error {
	fake.addRepoMutex.Lock()
	ret, specificReturn := fake.addRepoReturnsOnCall[len(fake.addRepoArgsForCall)]
	fake.addRepoArgsForCall = append(fake.addRepoArgsForCall, struct {
		arg1 context.Context
		arg2 model.GitRepository
	}{arg1, arg2})
	stub := fake.AddRepoStub
	fakeReturns := fake.addRepoReturns
	fake.recordInvocation("AddRepo", []interface{}{arg1, arg2})
	fake.addRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIGitRepositoryDAL) AddRepoCallCount() int {
	fake.addRepoMutex.RLock()
	defer fake.addRepoMutex.RUnlock()
	return len(fake.addRepoArgsForCall)
}

func (fake *FakeIGitRepositoryDAL) AddRepoCalls(stub func(context.Context, model.GitRepository) error) {
	fake.addRepoMutex.Lock()
	defer fake.addRepoMutex.Unlock()
	fake.AddRepoStub = stub
}

func (fake *FakeIGitRepositoryDAL) AddRepoArgsForCall(i int) (context.Context, model.GitRepository) {
	fake.addRepoMutex.RLock()
	defer fake.addRepoMutex.RUnlock()
	argsForCall := fake.addRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIGitRepositoryDAL) AddRepoReturns(result1 error) {
	fake.addRepoMutex.Lock()
	defer fake.addRepoMutex.Unlock()
	fake.AddRepoStub = nil
	fake.addRepoReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIGitRepositoryDAL) AddRepoReturnsOnCall(i int, result1 error) {
	fake.addRepoMutex.Lock()
	defer fake.addRepoMutex.Unlock()
	fake.AddRepoStub = nil
	if fake.addRepoReturnsOnCall == nil {
		fake.addRepoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.addRepoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIGitRepositoryDAL) ListRepoCursor(arg1 context.Context, arg2 string, arg3 int) ([]model.GitRepository, error) {
	fake.listRepoCursorMutex.Lock()
	ret, specificReturn := fake.listRepoCursorReturnsOnCall[len(fake.listRepoCursorArgsForCall)]
	fake.listRepoCursorArgsForCall = append(fake.listRepoCursorArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 int
	}{arg1, arg2, arg3})
	stub := fake.ListRepoCursorStub
	fakeReturns := fake.listRepoCursorReturns
	fake.recordInvocation("ListRepoCursor", []interface{}{arg1, arg2, arg3})
	fake.listRepoCursorMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIGitRepositoryDAL) ListRepoCursorCallCount() int {
	fake.listRepoCursorMutex.RLock()
	defer fake.listRepoCursorMutex.RUnlock()
	return len(fake.listRepoCursorArgsForCall)
}

func (fake *FakeIGitRepositoryDAL) ListRepoCursorCalls(stub func(context.Context, string, int) ([]model.GitRepository, error)) {
	fake.listRepoCursorMutex.Lock()
	defer fake.listRepoCursorMutex.Unlock()
	fake.ListRepoCursorStub = stub
}

func (fake *FakeIGitRepositoryDAL) ListRepoCursorArgsForCall(i int) (context.Context, string, int) {
	fake.listRepoCursorMutex.RLock()
	defer fake.listRepoCursorMutex.RUnlock()
	argsForCall := fake.listRepoCursorArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIGitRepositoryDAL) ListRepoCursorReturns(result1 []model.GitRepository, result2 error) {
	fake.listRepoCursorMutex.Lock()
	defer fake.listRepoCursorMutex.Unlock()
	fake.ListRepoCursorStub = nil
	fake.listRepoCursorReturns = struct {
		result1 []model.GitRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeIGitRepositoryDAL) ListRepoCursorReturnsOnCall(i int, result1 []model.GitRepository, result2 error) {
	fake.listRepoCursorMutex.Lock()
	defer fake.listRepoCursorMutex.Unlock()
	fake.ListRepoCursorStub = nil
	if fake.listRepoCursorReturnsOnCall == nil {
		fake.listRepoCursorReturnsOnCall = make(map[int]struct {
			result1 []model.GitRepository
			result2 error
		})
	}
	fake.listRepoCursorReturnsOnCall[i] = struct {
		result1 []model.GitRepository
		result2 error
	}{result1, result2}
}

func (fake *FakeIGitRepositoryDAL) RepoByName(arg1 context.Context, arg2 string) (model.GitRepository, string, error) {
	fake.repoByNameMutex.Lock()
	ret, specificReturn := fake.repoByNameReturnsOnCall[len(fake.repoByNameArgsForCall)]
	fake.repoByNameArgsForCall = append(fake.repoByNameArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.RepoByNameStub
	fakeReturns := fake.repoByNameReturns
	fake.recordInvocation("RepoByName", []interface{}{arg1, arg2})
	fake.repoByNameMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2, ret.result3
	}
	return fakeReturns.result1, fakeReturns.result2, fakeReturns.result3
}

func (fake *FakeIGitRepositoryDAL) RepoByNameCallCount() int {
	fake.repoByNameMutex.RLock()
	defer fake.repoByNameMutex.RUnlock()
	return len(fake.repoByNameArgsForCall)
}

func (fake *FakeIGitRepositoryDAL) RepoByNameCalls(stub func(context.Context, string) (model.GitRepository, string, error)) {
	fake.repoByNameMutex.Lock()
	defer fake.repoByNameMutex.Unlock()
	fake.RepoByNameStub = stub
}

func (fake *FakeIGitRepositoryDAL) RepoByNameArgsForCall(i int) (context.Context, string) {
	fake.repoByNameMutex.RLock()
	defer fake.repoByNameMutex.RUnlock()
	argsForCall := fake.repoByNameArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIGitRepositoryDAL) RepoByNameReturns(result1 model.GitRepository, result2 string, result3 error) {
	fake.repoByNameMutex.Lock()
	defer fake.repoByNameMutex.Unlock()
	fake.RepoByNameStub = nil
	fake.repoByNameReturns = struct {
		result1 model.GitRepository
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIGitRepositoryDAL) RepoByNameReturnsOnCall(i int, result1 model.GitRepository, result2 string, result3 error) {
	fake.repoByNameMutex.Lock()
	defer fake.repoByNameMutex.Unlock()
	fake.RepoByNameStub = nil
	if fake.repoByNameReturnsOnCall == nil {
		fake.repoByNameReturnsOnCall = make(map[int]struct {
			result1 model.GitRepository
			result2 string
			result3 error
		})
	}
	fake.repoByNameReturnsOnCall[i] = struct {
		result1 model.GitRepository
		result2 string
		result3 error
	}{result1, result2, result3}
}

func (fake *FakeIGitRepositoryDAL) RepoExists(arg1 context.Context, arg2 string) (bool, error) {
	fake.repoExistsMutex.Lock()
	ret, specificReturn := fake.repoExistsReturnsOnCall[len(fake.repoExistsArgsForCall)]
	fake.repoExistsArgsForCall = append(fake.repoExistsArgsForCall, struct {
		arg1 context.Context
		arg2 string
	}{arg1, arg2})
	stub := fake.RepoExistsStub
	fakeReturns := fake.repoExistsReturns
	fake.recordInvocation("RepoExists", []interface{}{arg1, arg2})
	fake.repoExistsMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeIGitRepositoryDAL) RepoExistsCallCount() int {
	fake.repoExistsMutex.RLock()
	defer fake.repoExistsMutex.RUnlock()
	return len(fake.repoExistsArgsForCall)
}

func (fake *FakeIGitRepositoryDAL) RepoExistsCalls(stub func(context.Context, string) (bool, error)) {
	fake.repoExistsMutex.Lock()
	defer fake.repoExistsMutex.Unlock()
	fake.RepoExistsStub = stub
}

func (fake *FakeIGitRepositoryDAL) RepoExistsArgsForCall(i int) (context.Context, string) {
	fake.repoExistsMutex.RLock()
	defer fake.repoExistsMutex.RUnlock()
	argsForCall := fake.repoExistsArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeIGitRepositoryDAL) RepoExistsReturns(result1 bool, result2 error) {
	fake.repoExistsMutex.Lock()
	defer fake.repoExistsMutex.Unlock()
	fake.RepoExistsStub = nil
	fake.repoExistsReturns = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeIGitRepositoryDAL) RepoExistsReturnsOnCall(i int, result1 bool, result2 error) {
	fake.repoExistsMutex.Lock()
	defer fake.repoExistsMutex.Unlock()
	fake.RepoExistsStub = nil
	if fake.repoExistsReturnsOnCall == nil {
		fake.repoExistsReturnsOnCall = make(map[int]struct {
			result1 bool
			result2 error
		})
	}
	fake.repoExistsReturnsOnCall[i] = struct {
		result1 bool
		result2 error
	}{result1, result2}
}

func (fake *FakeIGitRepositoryDAL) UpdateRepo(arg1 context.Context, arg2 string, arg3 map[string]any) error {
	fake.updateRepoMutex.Lock()
	ret, specificReturn := fake.updateRepoReturnsOnCall[len(fake.updateRepoArgsForCall)]
	fake.updateRepoArgsForCall = append(fake.updateRepoArgsForCall, struct {
		arg1 context.Context
		arg2 string
		arg3 map[string]any
	}{arg1, arg2, arg3})
	stub := fake.UpdateRepoStub
	fakeReturns := fake.updateRepoReturns
	fake.recordInvocation("UpdateRepo", []interface{}{arg1, arg2, arg3})
	fake.updateRepoMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeIGitRepositoryDAL) UpdateRepoCallCount() int {
	fake.updateRepoMutex.RLock()
	defer fake.updateRepoMutex.RUnlock()
	return len(fake.updateRepoArgsForCall)
}

func (fake *FakeIGitRepositoryDAL) UpdateRepoCalls(stub func(context.Context, string, map[string]any) error) {
	fake.updateRepoMutex.Lock()
	defer fake.updateRepoMutex.Unlock()
	fake.UpdateRepoStub = stub
}

func (fake *FakeIGitRepositoryDAL) UpdateRepoArgsForCall(i int) (context.Context, string, map[string]any) {
	fake.updateRepoMutex.RLock()
	defer fake.updateRepoMutex.RUnlock()
	argsForCall := fake.updateRepoArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeIGitRepositoryDAL) UpdateRepoReturns(result1 error) {
	fake.updateRepoMutex.Lock()
	defer fake.updateRepoMutex.Unlock()
	fake.UpdateRepoStub = nil
	fake.updateRepoReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeIGitRepositoryDAL) UpdateRepoReturnsOnCall(i int, result1 error) {
	fake.updateRepoMutex.Lock()
	defer fake.updateRepoMutex.Unlock()
	fake.UpdateRepoStub = nil
	if fake.updateRepoReturnsOnCall == nil {
		fake.updateRepoReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.updateRepoReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeIGitRepositoryDAL) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.addRepoMutex.RLock()
	defer fake.addRepoMutex.RUnlock()
	fake.listRepoCursorMutex.RLock()
	defer fake.listRepoCursorMutex.RUnlock()
	fake.repoByNameMutex.RLock()
	defer fake.repoByNameMutex.RUnlock()
	fake.repoExistsMutex.RLock()
	defer fake.repoExistsMutex.RUnlock()
	fake.updateRepoMutex.RLock()
	defer fake.updateRepoMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeIGitRepositoryDAL) recordInvocation(key string, args []interface{}) {
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

var _ dal.IGitRepositoryDAL = new(FakeIGitRepositoryDAL)
