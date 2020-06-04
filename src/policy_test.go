package retry4go

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetPolicy(t *testing.T) {
	p := NewDefaultPolicy()

	maxRetryTimes := uint(10)
	tErr := errors.New("test")
	retryable := func(err error) bool {
		if err == tErr {
			return true
		}
		return false
	}
	interval := 200 * time.Millisecond
	maxInterval := 10000 * time.Microsecond
	maxJitter := 300 * time.Microsecond
	regularInterval := 500 * time.Microsecond
	multiplier := 5.0
	retryType := RegularRetry
	randomFactor := 0.8

	var cfs []Config
	cfs = append(
		cfs,
		MaxRetryTimes(maxRetryTimes),
		Retryable(retryable),
		Interval(interval),
		MaxInterval(maxInterval),
		MaxJitterInterval(maxJitter),
		RegularInterval(regularInterval),
		Multiplier(multiplier),
		RetryType(retryType),
		RandomFactor(randomFactor),
	)

	for _, cf := range cfs {
		cf(p)
	}

	assert.EqualValues(t, maxRetryTimes, p.maxRetryTimes)
	r := p.isRetryableError(tErr, p.retryable)
	assert.True(t, r)
	assert.EqualValues(t, interval, p.interval)
	assert.EqualValues(t, maxInterval, p.maxInterval)
	assert.EqualValues(t, maxJitter, p.maxJitterInterval)
	assert.EqualValues(t, regularInterval, p.regularInterval)
	assert.EqualValues(t, multiplier, p.multiplier)
	assert.EqualValues(t, retryType, p.retryType)
	assert.EqualValues(t, randomFactor, p.randomFactor)
}

func TestDefaultRetryable(t *testing.T) {
	retryable := defaultRetryable()
	err := errors.New("test")
	assert.False(t, retryable(err))
}
