package retry

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRetryDo(t *testing.T) {
	tErr := errors.New("test")

	tests := []struct {
		name string
		ctx  context.Context
		rt   uint
		cr   checkRetryable
		i    time.Duration
		mi   time.Duration
		ji   time.Duration
		m    float64
		et   time.Duration

		// result
		resultRrr error
		doErr     error
		c         int
	}{
		{
			name: "retryable and retry one time",
			rt:   1,
			cr: func(err error) bool {
				if err == tErr {
					return true
				}
				return false
			},
			i:         100 * time.Millisecond,
			mi:        1000 * time.Millisecond,
			ji:        30 * time.Millisecond,
			m:         2,
			et:        5 * time.Minute,
			doErr:     tErr,
			resultRrr: tErr,
			c:         2,
		},
		{
			name: "not retryable",
			rt:   3,
			cr: func(err error) bool {
				if err == tErr {
					return false
				}
				return true
			},
			i:         100 * time.Millisecond,
			mi:        1000 * time.Millisecond,
			ji:        30 * time.Millisecond,
			m:         2,
			et:        5 * time.Minute,
			doErr:     tErr,
			resultRrr: tErr,
			c:         1,
		},
		{
			name: "retryable and retry max times",
			rt:   3,
			cr: func(err error) bool {
				if err == tErr {
					return true
				}
				return false
			},
			i:         100 * time.Millisecond,
			mi:        1000 * time.Millisecond,
			ji:        30 * time.Millisecond,
			m:         2,
			et:        5 * time.Minute,
			doErr:     tErr,
			resultRrr: tErr,
			c:         4,
		},
		{
			name: "no need retry",
			rt:   1,
			cr: func(err error) bool {
				if err == tErr {
					return true
				}
				return false
			},
			i:         100 * time.Millisecond,
			mi:        1000 * time.Millisecond,
			ji:        30 * time.Millisecond,
			m:         2,
			et:        5 * time.Minute,
			doErr:     nil,
			resultRrr: nil,
			c:         1,
		},
		{
			name: "max elapsed time",
			rt:   10,
			cr: func(err error) bool {
				if err == tErr {
					return true
				}
				return false
			},
			i:         100 * time.Millisecond,
			mi:        1000 * time.Millisecond,
			ji:        30 * time.Millisecond,
			m:         2,
			et:        400 * time.Millisecond,
			doErr:     tErr,
			resultRrr: tErr,
			c:         3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := 0
			err := Do(
				func() error {
					c++
					return tt.doErr
				},
				Interval(tt.i),
				MaxRetryTimes(tt.rt),
				MaxJitterInterval(tt.ji),
				CheckRetryable(tt.cr),
				MaxElapsedTime(tt.et),
				MaxInterval(tt.mi),
				Multiplier(tt.m),
			)
			assert.Equal(t, tt.resultRrr, err)
			assert.EqualValues(t, tt.c, c)
		})
	}
}

func TestDoWithContextMaxElapsedTime(t *testing.T) {
	t.Parallel()
	c := 0
	tErr := errors.New("test")
	ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
	defer cancel()

	start := time.Now()
	err := DoWithContext(
		ctx,
		func() error {
			c++
			return tErr
		},
		CheckRetryable(func(err error) bool {
			return true
		}),
	)
	end := time.Now()

	assert.EqualValues(t, tErr, err)
	assert.EqualValues(t, 3, c)
	assert.True(t, end.Sub(start).Milliseconds() >= 100)
	assert.True(t, end.Sub(start).Milliseconds() <= 600)
}

func TestDoWithContextDeadline(t *testing.T) {
	t.Parallel()
	c := 0
	ctx, cancel := context.WithTimeout(context.Background(), 1000*time.Millisecond)
	defer cancel()

	var err error
	var got string
	start := time.Now()
	err = DoWithContext(
		ctx,
		func() error {
			c++
			got, err = getStringForTest(ctx)
			return err
		},
		CheckRetryable(func(err error) bool {
			return true
		}),
	)
	end := time.Now()

	assert.EqualValues(t, context.DeadlineExceeded, err)
	assert.EqualValues(t, 1, c)
	assert.Equal(t, got, "")
	assert.True(t, end.Sub(start).Milliseconds() >= 100)
	assert.True(t, end.Sub(start).Milliseconds() <= 1100)
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

func TestDoWithContextCanceled(t *testing.T) {
	t.Parallel()
	c := 0
	tErr := errors.New("test")
	ctx, cancel := context.WithTimeout(context.Background(), 2000*time.Millisecond)

	go func() {
		time.Sleep(400 * time.Millisecond)
		cancel()
	}()
	start := time.Now()
	err := DoWithContext(
		ctx,
		func() error {
			c++
			time.Sleep(100 * time.Millisecond)
			return tErr
		},
		CheckRetryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
	)
	end := time.Now()

	assert.EqualValues(t, context.Canceled, err)
	assert.EqualValues(t, 2, c)
	assert.True(t, end.Sub(start).Milliseconds() >= 400)
	assert.True(t, end.Sub(start).Milliseconds() <= 800)
}
