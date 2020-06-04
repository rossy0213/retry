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
	d := r.getRandomizedInterval(r.nextInterval)
	r.nextInterval = time.Duration(math.Min(float64(r.nextInterval)+float64(r.regularInterval), float64(r.maxInterval)))
	return d
}

func (r *Regular) getRandomizedInterval(i time.Duration) time.Duration {
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	if rd.Float64() > 0.5 {
		return time.Duration(float64(i) + float64(r.regularInterval) - float64(r.maxJitterInterval)*rd.Float64())
	}
	return time.Duration(float64(i) + float64(r.regularInterval) + float64(r.maxJitterInterval)*rd.Float64())
}
