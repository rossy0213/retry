package retry4go

import (
	"context"
	"time"
)

const (
	DefaultMaxRetryTimes  = 5
	DefaultInterval       = 100.0 * time.Millisecond
	DefaultMaxInterval    = 1000.0 * time.Millisecond
	DefaultJitterInterval = 30.0 * time.Millisecond
	DefaultMultiplier     = 2.0
	DefaultElapsedTime    = 5 * time.Second
)

type RetryableFunc func() error

type checkRetryable func(err error) bool

func canRetry(cr checkRetryable, err error) bool {
	return cr(err)
}

func Do(fn RetryableFunc, cfs ...Config) error {
	return DoWithContext(nil, fn, cfs...)
}

func DoWithContext(ctx context.Context, fn RetryableFunc, cfs ...Config) error {
	eb := DefaultExponentialBackoff()
	for _, cf := range cfs {
		cf(eb)
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

		if !canRetry(be.checkRetryable, err) {
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
