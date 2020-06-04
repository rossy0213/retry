package retry4go

import (
	"time"
)

const (
	DefaultInterval        = 100.0 * time.Millisecond
	DefaultMaxInterval     = 1000.0 * time.Millisecond
	DefaultMaxRetryTimes   = 3
	DefaultJitterInterval  = DefaultInterval
	DefaultRegularInterval = DefaultInterval
	DefaultRetryType       = BackOffRetry
	DefaultMultiplier      = 2.0
	DefaultRandomFactor    = 0.5
)

type RetryableFunc func() error

func Do(fn RetryableFunc, cfs ...Config) error {
	var c uint
	var err error

	p := NewDefaultPolicy()
	for _, cf := range cfs {
		cf(p)
	}

	var rt Retry
	switch p.retryType {
	case RegularRetry:
		rt = NewRegular(p)
	default:
		rt = NewExponentialBackoff(p)
	}

	for c < p.maxRetryTimes {
		err = fn()

		if err != nil {
			if !p.isRetryableError(err, p.retryable) {
				return err
			}

			if c >= p.maxRetryTimes-1 {
				break
			}

			d := rt.Next()
			time.Sleep(d)
		} else {
			return nil
		}

		c++
	}
	return err
}
