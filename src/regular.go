package retry4go

import (
	"math"
	"math/rand"
	"time"
)

type Regular struct {
	interval          time.Duration
	maxInterval       time.Duration
	maxJitterInterval time.Duration
	regularInterval   time.Duration
	nextInterval      time.Duration
}

func NewRegular(p *Policy) *Regular {
	return &Regular{
		interval:          p.interval,
		maxInterval:       p.maxInterval,
		maxJitterInterval: p.maxJitterInterval,
		regularInterval:   p.regularInterval,
		nextInterval:      p.nextInterval,
	}
}

func (r *Regular) Next() time.Duration {
	r.nextInterval = r.getRandomizedInterval(r.nextInterval)
	return r.nextInterval
}

func (r *Regular) getRandomizedInterval(i time.Duration) time.Duration {
	f := rand.New(rand.NewSource(time.Now().UnixNano())).Float64()
	return time.Duration(math.Min(float64(i)+float64(r.regularInterval)+float64(r.maxJitterInterval)*f, float64(r.maxInterval)))
}
