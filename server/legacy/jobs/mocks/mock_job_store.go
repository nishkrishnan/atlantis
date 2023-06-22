// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/legacy/jobs (interfaces: JobStore)

package mocks

import (
	context "context"
	pegomock "github.com/petergtz/pegomock"
	jobs "github.com/runatlantis/atlantis/server/legacy/jobs"
	"reflect"
	"time"
)

type MockJobStore struct {
	fail func(message string, callerSkip ...int)
}

func NewMockJobStore(options ...pegomock.Option) *MockJobStore {
	mock := &MockJobStore{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockJobStore) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockJobStore) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockJobStore) AppendOutput(_param0 string, _param1 string) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockJobStore().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("AppendOutput", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockJobStore) Get(_param0 context.Context, _param1 string) (*jobs.Job, error) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockJobStore().")
	}
	params := []pegomock.Param{_param0, _param1}
	result := pegomock.GetGenericMockFrom(mock).Invoke("Get", params, []reflect.Type{reflect.TypeOf((**jobs.Job)(nil)).Elem(), reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 *jobs.Job
	var ret1 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(*jobs.Job)
		}
		if result[1] != nil {
			ret1 = result[1].(error)
		}
	}
	return ret0, ret1
}

func (mock *MockJobStore) RemoveJob(_param0 string) {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockJobStore().")
	}
	params := []pegomock.Param{_param0}
	pegomock.GetGenericMockFrom(mock).Invoke("RemoveJob", params, []reflect.Type{})
}

func (mock *MockJobStore) SetJobCompleteStatus(_param0 context.Context, _param1 string, _param2 jobs.JobStatus) error {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockJobStore().")
	}
	params := []pegomock.Param{_param0, _param1, _param2}
	result := pegomock.GetGenericMockFrom(mock).Invoke("SetJobCompleteStatus", params, []reflect.Type{reflect.TypeOf((*error)(nil)).Elem()})
	var ret0 error
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(error)
		}
	}
	return ret0
}

func (mock *MockJobStore) VerifyWasCalledOnce() *VerifierMockJobStore {
	return &VerifierMockJobStore{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockJobStore) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockJobStore {
	return &VerifierMockJobStore{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockJobStore) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockJobStore {
	return &VerifierMockJobStore{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockJobStore) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockJobStore {
	return &VerifierMockJobStore{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockJobStore struct {
	mock                   *MockJobStore
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockJobStore) AppendOutput(_param0 string, _param1 string) *MockJobStore_AppendOutput_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "AppendOutput", params, verifier.timeout)
	return &MockJobStore_AppendOutput_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockJobStore_AppendOutput_OngoingVerification struct {
	mock              *MockJobStore
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockJobStore_AppendOutput_OngoingVerification) GetCapturedArguments() (string, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockJobStore_AppendOutput_OngoingVerification) GetAllCapturedArguments() (_param0 []string, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockJobStore) Get(_param0 context.Context, _param1 string) *MockJobStore_Get_OngoingVerification {
	params := []pegomock.Param{_param0, _param1}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "Get", params, verifier.timeout)
	return &MockJobStore_Get_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockJobStore_Get_OngoingVerification struct {
	mock              *MockJobStore
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockJobStore_Get_OngoingVerification) GetCapturedArguments() (context.Context, string) {
	_param0, _param1 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1]
}

func (c *MockJobStore_Get_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockJobStore) RemoveJob(_param0 string) *MockJobStore_RemoveJob_OngoingVerification {
	params := []pegomock.Param{_param0}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "RemoveJob", params, verifier.timeout)
	return &MockJobStore_RemoveJob_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockJobStore_RemoveJob_OngoingVerification struct {
	mock              *MockJobStore
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockJobStore_RemoveJob_OngoingVerification) GetCapturedArguments() string {
	_param0 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1]
}

func (c *MockJobStore_RemoveJob_OngoingVerification) GetAllCapturedArguments() (_param0 []string) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]string, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(string)
		}
	}
	return
}

func (verifier *VerifierMockJobStore) SetJobCompleteStatus(_param0 context.Context, _param1 string, _param2 jobs.JobStatus) *MockJobStore_SetJobCompleteStatus_OngoingVerification {
	params := []pegomock.Param{_param0, _param1, _param2}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "SetJobCompleteStatus", params, verifier.timeout)
	return &MockJobStore_SetJobCompleteStatus_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockJobStore_SetJobCompleteStatus_OngoingVerification struct {
	mock              *MockJobStore
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockJobStore_SetJobCompleteStatus_OngoingVerification) GetCapturedArguments() (context.Context, string, jobs.JobStatus) {
	_param0, _param1, _param2 := c.GetAllCapturedArguments()
	return _param0[len(_param0)-1], _param1[len(_param1)-1], _param2[len(_param2)-1]
}

func (c *MockJobStore_SetJobCompleteStatus_OngoingVerification) GetAllCapturedArguments() (_param0 []context.Context, _param1 []string, _param2 []jobs.JobStatus) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]context.Context, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(context.Context)
		}
		_param1 = make([]string, len(c.methodInvocations))
		for u, param := range params[1] {
			_param1[u] = param.(string)
		}
		_param2 = make([]jobs.JobStatus, len(c.methodInvocations))
		for u, param := range params[2] {
			_param2[u] = param.(jobs.JobStatus)
		}
	}
	return
}