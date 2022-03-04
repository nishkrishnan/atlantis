// Code generated by pegomock. DO NOT EDIT.
// Source: github.com/runatlantis/atlantis/server/events (interfaces: StaleCommandChecker)

package mocks

import (
	"reflect"
	"time"

	pegomock "github.com/petergtz/pegomock"
	"github.com/runatlantis/atlantis/server/events/command"
)

type MockStaleCommandChecker struct {
	fail func(message string, callerSkip ...int)
}

func NewMockStaleCommandChecker(options ...pegomock.Option) *MockStaleCommandChecker {
	mock := &MockStaleCommandChecker{}
	for _, option := range options {
		option.Apply(mock)
	}
	return mock
}

func (mock *MockStaleCommandChecker) SetFailHandler(fh pegomock.FailHandler) { mock.fail = fh }
func (mock *MockStaleCommandChecker) FailHandler() pegomock.FailHandler      { return mock.fail }

func (mock *MockStaleCommandChecker) CommandIsStale(ctx *command.Context) bool {
	if mock == nil {
		panic("mock must not be nil. Use myMock := NewMockStaleCommandChecker().")
	}
	params := []pegomock.Param{ctx}
	result := pegomock.GetGenericMockFrom(mock).Invoke("CommandIsStale", params, []reflect.Type{reflect.TypeOf((*bool)(nil)).Elem()})
	var ret0 bool
	if len(result) != 0 {
		if result[0] != nil {
			ret0 = result[0].(bool)
		}
	}
	return ret0
}

func (mock *MockStaleCommandChecker) VerifyWasCalledOnce() *VerifierMockStaleCommandChecker {
	return &VerifierMockStaleCommandChecker{
		mock:                   mock,
		invocationCountMatcher: pegomock.Times(1),
	}
}

func (mock *MockStaleCommandChecker) VerifyWasCalled(invocationCountMatcher pegomock.InvocationCountMatcher) *VerifierMockStaleCommandChecker {
	return &VerifierMockStaleCommandChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
	}
}

func (mock *MockStaleCommandChecker) VerifyWasCalledInOrder(invocationCountMatcher pegomock.InvocationCountMatcher, inOrderContext *pegomock.InOrderContext) *VerifierMockStaleCommandChecker {
	return &VerifierMockStaleCommandChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		inOrderContext:         inOrderContext,
	}
}

func (mock *MockStaleCommandChecker) VerifyWasCalledEventually(invocationCountMatcher pegomock.InvocationCountMatcher, timeout time.Duration) *VerifierMockStaleCommandChecker {
	return &VerifierMockStaleCommandChecker{
		mock:                   mock,
		invocationCountMatcher: invocationCountMatcher,
		timeout:                timeout,
	}
}

type VerifierMockStaleCommandChecker struct {
	mock                   *MockStaleCommandChecker
	invocationCountMatcher pegomock.InvocationCountMatcher
	inOrderContext         *pegomock.InOrderContext
	timeout                time.Duration
}

func (verifier *VerifierMockStaleCommandChecker) CommandIsStale(ctx *command.Context) *MockStaleCommandChecker_CommandIsStale_OngoingVerification {
	params := []pegomock.Param{ctx}
	methodInvocations := pegomock.GetGenericMockFrom(verifier.mock).Verify(verifier.inOrderContext, verifier.invocationCountMatcher, "CommandIsStale", params, verifier.timeout)
	return &MockStaleCommandChecker_CommandIsStale_OngoingVerification{mock: verifier.mock, methodInvocations: methodInvocations}
}

type MockStaleCommandChecker_CommandIsStale_OngoingVerification struct {
	mock              *MockStaleCommandChecker
	methodInvocations []pegomock.MethodInvocation
}

func (c *MockStaleCommandChecker_CommandIsStale_OngoingVerification) GetCapturedArguments() *command.Context {
	ctx := c.GetAllCapturedArguments()
	return ctx[len(ctx)-1]
}

func (c *MockStaleCommandChecker_CommandIsStale_OngoingVerification) GetAllCapturedArguments() (_param0 []*command.Context) {
	params := pegomock.GetGenericMockFrom(c.mock).GetInvocationParams(c.methodInvocations)
	if len(params) > 0 {
		_param0 = make([]*command.Context, len(c.methodInvocations))
		for u, param := range params[0] {
			_param0[u] = param.(*command.Context)
		}
	}
	return
}