package retry4go

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDoRetryable(t *testing.T) {
	tErr := errors.New("test")
	err := Do(
		func() error {
			return tErr
		},
		Retryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
	)
	assert.Equal(t, err, tErr)
}

func TestDoNotRetryable(t *testing.T) {
	tErr := errors.New("test")
	c := 0
	err := Do(
		func() error {
			c++
			return errors.New("not retryable")
		},
		Retryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
	)
	assert.Equal(t, err, errors.New("not retryable"))
	assert.Equal(t, 1, c)
}

func TestDoNoRetry(t *testing.T) {
	c := 0
	err := Do(
		func() error {
			c++
			return nil
		},
	)
	assert.Nil(t, err)
	assert.Equal(t, 1, c)
}

func TestDoRegular(t *testing.T) {
	tErr := errors.New("test")
	c := 0
	err := Do(
		func() error {
			c++
			if c == 2 {
				return nil
			}
			return tErr
		},
		Retryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
		RetryType(RegularRetry),
		RegularInterval(400*time.Millisecond),
	)
	assert.Nil(t, err)
	assert.Equal(t, 2, c)
}

func TestDoMaxRegularRetry(t *testing.T) {
	tErr := errors.New("test")
	maxRetryTimes := 4

	c := 0
	start := time.Now()
	err := Do(
		func() error {
			c++
			return tErr
		},
		MaxRetryTimes(uint(maxRetryTimes)),
		Retryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
		RetryType(RegularRetry),
	)
	end := time.Now()

	assert.Equal(t, err, tErr)
	assert.Equal(t, maxRetryTimes, c)
	assert.True(t, end.Sub(start).Milliseconds() < 1200)
}

func TestDoMaxBackoffRetry(t *testing.T) {
	tErr := errors.New("test")
	maxRetryTimes := 4

	c := 0
	start := time.Now()
	err := Do(
		func() error {
			c++
			return tErr
		},
		MaxRetryTimes(uint(maxRetryTimes)),
		Retryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
	)
	end := time.Now()

	assert.Equal(t, err, tErr)
	assert.Equal(t, maxRetryTimes, c)
	assert.True(t, end.Sub(start).Milliseconds() < 1000)
}
