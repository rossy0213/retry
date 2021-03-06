package retry

import (
	"context"
	"time"
)

type BackoffContext interface {
	Backoff
	Context() context.Context
}

type backoffWithContext struct {
	*exponentialBackoff
	ctx context.Context
}

func withContext(ctx context.Context, eb *exponentialBackoff) *backoffWithContext {
	return &backoffWithContext{
		exponentialBackoff: eb,
		ctx:                ctx,
	}
}

// Get Context from Backoff
func (bc *backoffWithContext) Context() context.Context {
	return bc.ctx
}

// Return the waiting time.
// Will return Stop when if received message from context.Done().
// And if maxRetryTimes == 0 or smaller than retryTimes, will return Stop.
func (bc *backoffWithContext) Next() time.Duration {
	select {
	case <-bc.ctx.Done():
		return Stop
	default:
	}

	next := bc.exponentialBackoff.Next()
	if deadline, ok := bc.ctx.Deadline(); ok && deadline.Sub(time.Now()) < next {
		return Stop
	}

	return next
}
