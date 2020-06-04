package retry4go

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDefaultRegular_Next(t *testing.T) {
	p := NewDefaultPolicy()
	eb := NewRegular(p)

	var d time.Duration
	for {
		interval := eb.nextInterval
		d = eb.Next()
		if eb.nextInterval >= eb.maxInterval {
			break
		}
		assert.True(t, time.Duration(float64(interval)+float64(eb.regularInterval)-float64(eb.maxJitterInterval)) <= d)
		assert.True(t, time.Duration(float64(interval)+float64(eb.regularInterval)+float64(eb.maxJitterInterval)) >= d)
	}
}
