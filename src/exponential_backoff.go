package retry

import (
	"math"
	"math/rand"
	"time"
)

const (
	DefaultMaxRetryTimes  = 10
	DefaultInterval       = 100.0 * time.Millisecond
	DefaultMaxInterval    = 1000.0 * time.Millisecond
	DefaultJitterInterval = 30.0 * time.Millisecond
	DefaultMultiplier     = 2.0
	DefaultMaxElapsedTime = 5 * time.Minute
)

type exponentialBackoff struct {
	maxRetryTimes     uint
	retryTimes        uint
	checkRetryable    checkRetryable
	interval          time.Duration
	maxJitterInterval time.Duration
	maxInterval       time.Duration
	nextInterval      time.Duration
	multiplier        float64

	startTime      time.Time
	maxElapsedTime time.Duration
}

func DefaultExponentialBackoff() *exponentialBackoff {
	return &exponentialBackoff{
		maxRetryTimes:     DefaultMaxRetryTimes,
		retryTimes:        0,
		interval:          DefaultInterval,
		maxInterval:       DefaultMaxInterval,
		maxJitterInterval: DefaultJitterInterval,
		nextInterval:      DefaultInterval,
		multiplier:        DefaultMultiplier,
		startTime:         time.Now(),
		maxElapsedTime:    DefaultMaxElapsedTime,
		checkRetryable: func(error) bool {
			return false
		},
	}
}

// Return the waiting time.
// If maxRetryTimes == 0 or smaller than retryTimes, will return Stop.
func (eb *exponentialBackoff) Next() time.Duration {
	if eb.maxRetryTimes == 0 || eb.retryTimes >= eb.maxRetryTimes {
		return Stop
	}
	return eb.NextBackoff()
}

// If waiting time bigger than time remaining, will return Stop.
func (eb *exponentialBackoff) NextBackoff() time.Duration {
	next := eb.getRandomizedInterval(eb.nextInterval)
	eb.retryTimes++

	// calculate and assign the correct nextInterval
	eb.nextInterval = time.Duration(math.Min(float64(eb.nextInterval)*eb.multiplier, float64(eb.maxInterval)))

	elapsed := eb.getElapsedTime()
	if eb.maxElapsedTime != 0 && elapsed+next > eb.maxElapsedTime {
		return Stop
	}
	return next
}

// Return time that has elapsed since the backoff instance was created.
func (eb *exponentialBackoff) getElapsedTime() time.Duration {
	return time.Now().Sub(eb.startTime)
}

// Will return : [ nextInterval - maxJitter, nextInterval - maxJitter]
func (eb *exponentialBackoff) getRandomizedInterval(i time.Duration) time.Duration {
	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := float64(i) - float64(eb.maxJitterInterval)
	max := float64(i) + float64(eb.maxJitterInterval)

	return time.Duration(min + ((max - min) * s.Float64()))
}
