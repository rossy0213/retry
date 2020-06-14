package retry4go

import (
	"math"
	"math/rand"
	"time"
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

	maxElapsedTime time.Duration
	startTime      time.Time
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

		// need move to policy
		checkRetryable: defaultCheckRetryable(),
		maxElapsedTime: DefaultElapsedTime,
	}
}

func (eb *exponentialBackoff) Next() time.Duration {
	if eb.maxRetryTimes == 0 || eb.retryTimes >= eb.maxRetryTimes {
		return Stop
	}
	return eb.NextBackoff()
}

func (eb *exponentialBackoff) NextBackoff() time.Duration {
	next := eb.getRandomizedInterval(eb.retryTimes, eb.nextInterval)
	eb.retryTimes++

	// calculate and assign the correct nextInterval
	eb.nextInterval = time.Duration(math.Min(float64(eb.nextInterval)*eb.multiplier, float64(eb.maxInterval)))

	elapsed := eb.getElapsedTime()
	if eb.maxElapsedTime != 0 && elapsed+next > eb.maxElapsedTime {
		return Stop
	}
	return next
}

func (eb *exponentialBackoff) getElapsedTime() time.Duration {
	return time.Now().Sub(eb.startTime)
}

// formula: [ nextInterval - maxJitter, nextInterval - maxJitter]
func (eb *exponentialBackoff) getRandomizedInterval(t uint, i time.Duration) time.Duration {
	if t == 0 {
		return i
	}

	s := rand.New(rand.NewSource(time.Now().UnixNano()))
	min := float64(i) - float64(eb.maxJitterInterval)
	max := float64(i) + float64(eb.maxJitterInterval)

	return time.Duration(min + ((max - min) * s.Float64()))
}
