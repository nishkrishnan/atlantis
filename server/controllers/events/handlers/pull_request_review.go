package handlers

import (
	"context"
	"github.com/runatlantis/atlantis/server/events"
	"github.com/runatlantis/atlantis/server/http"
	"github.com/runatlantis/atlantis/server/logging"
	"github.com/runatlantis/atlantis/server/neptune/gateway/event"
)

type PullRequestReviewEventHandler struct {
	PRReviewCommandRunner events.CommandRunner
}

func (p PullRequestReviewEventHandler) Handle(ctx context.Context, event event.PullRequestReview, _ *http.BufferedRequest) error {
	p.PRReviewCommandRunner.RunPRReviewCommand(
		ctx,
		event.Repo,
		event.Pull,
		event.User,
		event.Timestamp,
		event.InstallationToken,
	)
	return nil
}

type AsyncPullRequestReviewEvent struct {
	handler *PullRequestReviewEventHandler
	logger  logging.Logger
}

func NewPullRequestReviewEvent(prReviewCommandRunner events.CommandRunner, logger logging.Logger) *AsyncPullRequestReviewEvent {
	return &AsyncPullRequestReviewEvent{
		handler: &PullRequestReviewEventHandler{
			PRReviewCommandRunner: prReviewCommandRunner,
		},
		logger: logger,
	}
}

func (a AsyncPullRequestReviewEvent) Handle(_ context.Context, event event.PullRequestReview, req *http.BufferedRequest) error {
	go func() {
		// Passing background context to avoid context cancellation since the parent goroutine does not wait for this goroutine to finish execution.
		err := a.handler.Handle(context.Background(), event, req)
		if err != nil {
			a.logger.ErrorContext(context.Background(), err.Error())
		}
	}()
	return nil
}