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
	err := Do(
		func() error {
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
}

func TestDoNoRetry(t *testing.T) {
	tErr := errors.New("test")
	err := Do(
		func() error {
			return nil
		},
		Retryable(func(err error) bool {
			if err == tErr {
				return true
			}
			return false
		}),
	)
	assert.Nil(t, err)
}

func TestDoRegular(t *testing.T) {
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
		RetryType(RegularRetry),
		RegularInterval(300*time.Millisecond),
	)
	assert.Equal(t, err, tErr)
}

func TestDoMaxRetry(t *testing.T) {
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
		RetryType(RegularRetry),
		RegularInterval(300*time.Millisecond),
	)
	assert.Equal(t, err, tErr)
}
