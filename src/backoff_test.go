package retry4go

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultBackoff_Next(t *testing.T) {
	p := NewDefaultPolicy()
	eb := NewExponentialBackoff(p)

	var d time.Duration
	for {
		interval := eb.nextInterval
		d = eb.Next()
		if eb.nextInterval >= eb.maxInterval {
			break
		}

		assert.True(t, time.Duration(float64(interval)*(1-eb.randomFactor)) <= d)
		assert.True(t, time.Duration(float64(interval)*(1+eb.randomFactor)) >= d)
	}
}
