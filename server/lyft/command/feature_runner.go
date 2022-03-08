package command

import (
	"github.com/runatlantis/atlantis/server/events"
	"github.com/runatlantis/atlantis/server/events/command"
	"github.com/runatlantis/atlantis/server/logging"
	"github.com/runatlantis/atlantis/server/lyft/feature"
)

func NewPlatformModeFeatureRunner(
	featureAllocator feature.Allocator,
	platformModeEnabled bool,
	logger logging.SimpleLogging,
	allocatedRunner events.CommentCommandRunner,
	unallocatedRunner events.CommentCommandRunner,
) *PlatformModeFeatureRunner {
	return &PlatformModeFeatureRunner{
		featureAllocator:    featureAllocator,
		platformModeEnabled: platformModeEnabled,
		logger:              logger,
		allocatedRunner:     allocatedRunner,
		unallocatedRunner:   unallocatedRunner,
	}
}

// PlatformModeFeatureRunner basic struct to that prepares a runner based on if
// the platform mode feature is enabled for the repo and if the server is
// running in platform mode
type PlatformModeFeatureRunner struct {
	featureAllocator    feature.Allocator
	platformModeEnabled bool
	logger              logging.SimpleLogging
	allocatedRunner     events.CommentCommandRunner
	unallocatedRunner   events.CommentCommandRunner
}

// Wrap returns CommentCommandRunner that encapsulates feature flags decision
// inside a CommentCommandRunner interface
func (r *PlatformModeFeatureRunner) Run(ctx *command.Context, cmd *events.CommentCommand) {
	// if platform mode is not enable run unallocatedRunner runner. No need
	// to invoke feature allocator
	if !r.platformModeEnabled {
		r.unallocatedRunner.Run(ctx, cmd)
		return
	}

	shouldAllocate, err := r.featureAllocator.ShouldAllocate(feature.PlatformMode, ctx.HeadRepo.FullName)
	if err != nil {
		r.logger.Err("unable to allocate for feature: %s, error: %s", feature.PlatformMode, err)
	}

	if !shouldAllocate {
		r.unallocatedRunner.Run(ctx, cmd)
		return
	}

	r.allocatedRunner.Run(ctx, cmd)
}