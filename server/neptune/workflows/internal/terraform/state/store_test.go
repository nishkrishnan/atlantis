package state_test

import (
	"bytes"
	"net/url"
	"testing"

	"github.com/gorilla/mux"
	"github.com/runatlantis/atlantis/server/neptune/workflows/internal/terraform/state"
	"github.com/stretchr/testify/assert"
)

type testNotifier struct {
	expectedState *state.Workflow
	t             *testing.T
	called        bool
}

func (n *testNotifier) notify(state *state.Workflow) error {
	n.called = true
	assert.Equal(n.t, n.expectedState, state)
	return nil
}

func TestInitPlanJob(t *testing.T) {

	exoectedURL, err := url.Parse("www.test.com/jobs/1234")
	assert.NoError(t, err)

	notifier := &testNotifier{
		expectedState: &state.Workflow{
			Plan: &state.Job{
				Status: state.InProgressJobStatus,
				Output: &state.JobOutput{
					URL: exoectedURL,
				},
			},
		},
		t: t,
	}

	subject := state.NewWorkflowStore(
		notifier.notify,
	)

	jobID := bytes.NewBufferString("1234")
	baseURL := bytes.NewBufferString("www.test.com")

	err = subject.InitPlanJob(jobID, baseURL)
	assert.NoError(t, err)
	assert.True(t, notifier.called)
}

func TestInitApplyJob(t *testing.T) {
	route := &mux.Route{}
	route.Path("/jobs/{job-id}")

	exoectedURL, err := url.Parse("www.test.com/jobs/1234")
	assert.NoError(t, err)

	notifier := &testNotifier{
		expectedState: &state.Workflow{
			Apply: &state.Job{
				Status: state.InProgressJobStatus,
				Output: &state.JobOutput{
					URL: exoectedURL,
				},
			},
		},
		t: t,
	}

	subject := state.NewWorkflowStore(
		notifier.notify,
	)

	jobID := bytes.NewBufferString("1234")
	baseURL := bytes.NewBufferString("www.test.com")

	err = subject.InitApplyJob(jobID, baseURL)
	assert.NoError(t, err)
	assert.True(t, notifier.called)
}

func TestUpdateApplyJob(t *testing.T) {
	route := &mux.Route{}
	route.Path("/jobs/{job-id}")

	exoectedURL, err := url.Parse("www.test.com/jobs/1234")
	assert.NoError(t, err)
	notifier := &testNotifier{
		expectedState: &state.Workflow{
			Apply: &state.Job{
				Status: state.InProgressJobStatus,
				Output: &state.JobOutput{
					URL: exoectedURL,
				},
			},
		},
		t: t,
	}

	subject := state.NewWorkflowStore(
		notifier.notify,
	)

	jobID := bytes.NewBufferString("1234")
	baseURL := bytes.NewBufferString("www.test.com")

	// init and then update
	err = subject.InitApplyJob(jobID, baseURL)
	assert.NoError(t, err)

	notifier.expectedState.Apply.Status = state.FailedJobStatus

	err = subject.UpdateApplyJobWithStatus(state.FailedJobStatus)
	assert.NoError(t, err)
	assert.True(t, notifier.called)
}

func TestUpdatePlanJob(t *testing.T) {
	route := &mux.Route{}
	route.Path("/jobs/{job-id}")

	exoectedURL, err := url.Parse("www.test.com/jobs/1234")
	assert.NoError(t, err)
	notifier := &testNotifier{
		expectedState: &state.Workflow{
			Plan: &state.Job{
				Status: state.InProgressJobStatus,
				Output: &state.JobOutput{
					URL: exoectedURL,
				},
			},
		},
		t: t,
	}

	subject := state.NewWorkflowStore(
		notifier.notify,
	)

	jobID := bytes.NewBufferString("1234")
	baseURL := bytes.NewBufferString("www.test.com")

	// init and then update
	err = subject.InitPlanJob(jobID, baseURL)
	assert.NoError(t, err)

	notifier.expectedState.Plan.Status = state.FailedJobStatus

	err = subject.UpdatePlanJobWithStatus(state.FailedJobStatus)
	assert.NoError(t, err)
	assert.True(t, notifier.called)
}