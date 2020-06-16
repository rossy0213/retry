package retry

import (
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBackoff_Next(t *testing.T) {
	t.Parallel()
	eb := DefaultExponentialBackoff()

	var d time.Duration
	for {
		interval := eb.nextInterval
		d = eb.Next()
		if eb.nextInterval >= eb.maxInterval {
			break
		}

		assert.True(t, interval-eb.maxJitterInterval <= d)
		assert.True(t, interval+eb.maxJitterInterval >= d)
	}
	err := errors.New("test")
	assert.False(t, eb.checkRetryable(err))
}
