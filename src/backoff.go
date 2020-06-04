package retry4go

import (
	"math"
	"math/rand"
	"time"
)

// Inspired by https://googleapis.dev/java/google-http-client/latest/com/google/api/client/util/ExponentialBackOff.html
type ExponentialBackoff struct {
	interval     time.Duration
	maxInterval  time.Duration
	multiplier   float64
	nextInterval time.Duration
	randomFactor float64 // [0 ~ 1]
}

func NewExponentialBackoff(p *Policy) *ExponentialBackoff {
	return &ExponentialBackoff{
		interval:     p.interval,
		maxInterval:  p.maxInterval,
		multiplier:   p.multiplier,
		nextInterval: p.nextInterval,
		randomFactor: p.randomFactor,
	}
}

func (eb *ExponentialBackoff) Next() time.Duration {
	// get random interval from calculated/initial next interval
	d := eb.getRandomizedInterval(eb.nextInterval)

	// calculate and update next interval
	eb.nextInterval = time.Duration(math.Min(float64(eb.nextInterval)*eb.multiplier, float64(eb.maxInterval)))

	return d
}

// formula: [ (1 - randomFactor) * duration, (1 + randomFactor) * duration]
func (eb *ExponentialBackoff) getRandomizedInterval(i time.Duration) time.Duration {
	f := rand.New(rand.NewSource(time.Now().UnixNano()))
	delta := eb.randomFactor * float64(i)
	min := float64(i) - delta
	max := float64(i) + delta

	return time.Duration(min + (f.Float64() * (max - min + 1)))
}
