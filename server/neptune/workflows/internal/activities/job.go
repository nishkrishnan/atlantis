package activities

import (
	"context"

	"github.com/pkg/errors"
	"github.com/runatlantis/atlantis/server/neptune/logger"
)

type streamCloser interface {
	Close(ctx context.Context, jobID string) error
}

type jobActivities struct {
	StreamCloser streamCloser
}

type CloseJobRequest struct {
	JobID string
}

func (t *jobActivities) CloseJob(ctx context.Context, request CloseJobRequest) error {
	err := t.StreamCloser.Close(ctx, request.JobID)
	if err != nil {
		logger.Error(ctx, errors.Wrapf(err, "closing job").Error())
	}
	return err
}