package retry4go

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetBackoff(t *testing.T) {
	eb := DefaultExponentialBackoff()

	maxRetryTimes := uint(10)
	tErr := errors.New("test")
	checkRetryable := func(err error) bool {
		if err == tErr {
			return true
		}
		return false
	}
	interval := 200 * time.Millisecond
	maxInterval := 10000 * time.Microsecond
	maxJitter := 300 * time.Microsecond
	multiplier := 5.0
	maxElapsedTime := 100 * time.Second

	var cfs []Config
	cfs = append(
		cfs,
		MaxRetryTimes(maxRetryTimes),
		CheckRetryable(checkRetryable),
		Interval(interval),
		MaxInterval(maxInterval),
		MaxJitterInterval(maxJitter),
		Multiplier(multiplier),
		MaxElapsedTime(maxElapsedTime),
	)

	for _, cf := range cfs {
		cf(eb)
	}

	assert.EqualValues(t, maxRetryTimes, eb.maxRetryTimes)
	r := canRetry(eb.checkRetryable, tErr)
	assert.True(t, r)
	assert.EqualValues(t, interval, eb.interval)
	assert.EqualValues(t, maxInterval, eb.maxInterval)
	assert.EqualValues(t, maxJitter, eb.maxJitterInterval)
	assert.EqualValues(t, multiplier, eb.multiplier)
	assert.EqualValues(t, maxElapsedTime, eb.maxElapsedTime)
}

func TestDefaultCheckRetryable(t *testing.T) {
	retryable := defaultCheckRetryable()
	err := errors.New("test")
	assert.False(t, retryable(err))
}
