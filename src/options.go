package retry

import (
	"time"
)

// Option
type Option func(*exponentialBackoff)

// Maximum retry number of times you can retry.
// Default is 10.
func MaxRetryTimes(mt uint) Option {
	return func(eb *exponentialBackoff) {
		eb.maxRetryTimes = mt
	}
}

// Check if the error returned by retryableFunc is a retryable error.
// Default will return false for everything.
func CheckRetryable(cr checkRetryable) Option {
	return func(eb *exponentialBackoff) {
		eb.checkRetryable = cr
	}
}

// Standard　waiting Time, will use to calculate the nextInterval
// and it will be first retry's waiting time.
// Default is 100.0 * time.Millisecond.
func Interval(i time.Duration) Option {
	return func(eb *exponentialBackoff) {
		eb.interval = i
		eb.nextInterval = i
	}
}

// Maximum possible waiting time.
// Default is 1000.0 * time.Millisecond.
func MaxInterval(mi time.Duration) Option {
	return func(eb *exponentialBackoff) {
		eb.maxInterval = mi
	}
}

// Used to delay the waiting time.
// Default is 30.0 * time.Millisecond.
func MaxJitterInterval(ji time.Duration) Option {
	return func(eb *exponentialBackoff) {
		eb.maxJitterInterval = ji
	}
}

// Used to calculate the nextInterval of Exponential Backoff.
// Default is 2.
func Multiplier(m float64) Option {
	return func(eb *exponentialBackoff) {
		eb.multiplier = m
	}
}

// Time limit including retries.
// Default is 5 * time.Minute
func MaxElapsedTime(et time.Duration) Option {
	return func(eb *exponentialBackoff) {
		eb.maxElapsedTime = et
	}
}
