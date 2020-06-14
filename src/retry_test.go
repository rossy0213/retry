package retry4go

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// TODO: change to table driven tests

func TestDoRetryable(t *testing.T) {
	tErr := errors.New("test")
	err := Do(
		func() error {
			return tErr
		},
		CheckRetryable(func(err error) bool {
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
		CheckRetryable(func(err error) bool {
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

func TestDoMaxBackoffRetry(t *testing.T) {
	tErr := errors.New("test")
	maxRetryTimes := 4
	maxJitterInterval := 30 * time.Millisecond

	c := 0
	start := time.Now()
	err := Do(
		func() error {
			c++
			return tErr
		},
		MaxRetryTimes(uint(maxRetryTimes)),
		CheckRetryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
		MaxJitterInterval(maxJitterInterval),
	)
	end := time.Now()

	assert.Equal(t, err, tErr)
	assert.Equal(t, maxRetryTimes+1, c)
	assert.True(t, end.Sub(start).Milliseconds() >= 1500-int64(maxRetryTimes)*int64(maxJitterInterval))
	assert.True(t, end.Sub(start).Milliseconds() <= 1600+int64(maxRetryTimes)*int64(maxJitterInterval))
}

func TestDoMaxElapsedTime(t *testing.T) {
	tErr := errors.New("test")
	maxElapsedTime := 200 * time.Millisecond

	c := 0
	start := time.Now()
	err := Do(
		func() error {
			c++
			return tErr
		},
		CheckRetryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
		MaxElapsedTime(maxElapsedTime),
	)
	end := time.Now()

	assert.Equal(t, err, tErr)
	assert.Equal(t, 2, c)
	assert.True(t, end.Sub(start) >= DefaultInterval-DefaultJitterInterval)
}

func TestDoWithContext(t *testing.T) {
	tErr := errors.New("test")
	maxRetryTimes := 3
	maxJitterInterval := 0 * time.Millisecond

	c := 0
	start := time.Now()
	ctx := context.Background()
	reqCtx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()
	err := DoWithContext(
		reqCtx,
		func() error {
			c++
			return tErr
		},
		MaxRetryTimes(uint(maxRetryTimes)),
		CheckRetryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
		MaxJitterInterval(maxJitterInterval),
	)
	end := time.Now()

	assert.Equal(t, err, tErr)
	assert.Equal(t, 3, c)
	assert.True(t, end.Sub(start).Milliseconds() >= 300)
	assert.True(t, end.Sub(start).Milliseconds() <= 500)
}

func TestDoWithContextCancel(t *testing.T) {
	tErr := errors.New("test")
	cancelOn := 2
	maxRetryTimes := 3
	maxJitterInterval := 0 * time.Millisecond

	c := 0
	start := time.Now()
	ctx := context.Background()
	reqCtx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	err := DoWithContext(
		reqCtx,
		func() error {
			c++
			if c == cancelOn {
				cancel()
			}
			return tErr
		},
		MaxRetryTimes(uint(maxRetryTimes)),
		CheckRetryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
		MaxJitterInterval(maxJitterInterval),
	)
	end := time.Now()

	assert.Equal(t, err, tErr)
	assert.Equal(t, 2, c)
	assert.True(t, end.Sub(start).Milliseconds() >= 100)
	assert.True(t, end.Sub(start).Milliseconds() <= 300)
}

func getStringForTest(ctx context.Context) (got string, err error) {
	t := time.NewTimer(time.Minute)
	select {
	case <-ctx.Done():
		return got, ctx.Err()
	case <-t.C:
		return "test", nil
	}
}

func TestDoRetryFast(t *testing.T) {
	c := 0
	maxRetryTimes := uint(1)
	ctx := context.Background()
	interval := 100 * time.Millisecond

	var err error
	var got string
	timeout := 200 * time.Millisecond
	start := time.Now()
	err = Do(
		func() error {
			if c > 0 {
				timeout *= 2
			}
			reqCtx, cancel := context.WithTimeout(ctx, timeout)
			defer cancel()

			c++
			got, err = getStringForTest(reqCtx)
			return err
		},
		CheckRetryable(func(err error) bool {
			return true
		}),
		Interval(interval),
		MaxRetryTimes(maxRetryTimes),
	)
	end := time.Now()

	assert.Error(t, err)
	assert.EqualValues(t, maxRetryTimes+1, c)
	assert.Equal(t, got, "")
	assert.True(t, end.Sub(start).Milliseconds() >= 100)
	assert.True(t, end.Sub(start).Milliseconds() <= 800)
}
