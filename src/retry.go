package retry

import (
	"context"
	"time"
)

// The function will be retried when if it return an error.
type RetryableFunc func() error

// Using to check if error from RetryableFunc can retry or not.
type checkRetryable func(err error) bool

// Do creates a new Backff for the retry with the passed options,
// And will retry RetryableFunc when if it return an error.
func Do(fn RetryableFunc, ops ...Option) error {
	return DoWithContext(nil, fn, ops...)
}

// DoWithContext using goroutine to sleep and monitoring context.Done().
// Will stop the retry once receive it from Context.Done().
func DoWithContext(ctx context.Context, fn RetryableFunc, ops ...Option) error {
	eb := DefaultExponentialBackoff()
	for _, op := range ops {
		op(eb)
	}

	if ctx == nil {
		ctx = context.Background()
	}
	be := withContext(ctx, eb)

	t := &defaultTimer{}
	defer func() {
		t.Stop()
	}()

	var err error
	var next time.Duration
	for {
		if err = fn(); err == nil {
			return nil
		}

		if !be.checkRetryable(err) {
			return err
		}

		if next = be.Next(); next == Stop {
			return err
		}

		t.Start(next)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-t.C():
		}
	}
}
