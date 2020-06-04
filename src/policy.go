package retry4go

import (
	"time"
)

type retryType int

type isRetryable func(error) bool

type Config func(*Policy)

type Policy struct {
	maxRetryTimes     uint
	retryable         func(err error) bool
	retryType         retryType
	interval          time.Duration
	maxInterval       time.Duration
	maxJitterInterval time.Duration
	nextInterval      time.Duration
	regularInterval   time.Duration
	multiplier        float64
	randomFactor      float64 // [0 ~ 1]
}

func NewDefaultPolicy() *Policy {
	return &Policy{
		maxRetryTimes:     DefaultMaxRetryTimes,
		retryable:         defaultRetryable(),
		interval:          DefaultInterval,
		maxInterval:       DefaultMaxInterval,
		maxJitterInterval: DefaultJitterInterval,
		nextInterval:      DefaultInterval,
		regularInterval:   DefaultRegularInterval,
		multiplier:        DefaultMultiplier,
		retryType:         DefaultRetryType,
		randomFactor:      DefaultRandomFactor,
	}
}

func defaultRetryable() isRetryable {
	return func(error) bool {
		return false
	}
}

func MaxRetryTimes(rt uint) Config {
	return func(p *Policy) {
		p.maxRetryTimes = rt
	}
}

func Retryable(ir isRetryable) Config {
	return func(p *Policy) {
		p.retryable = ir
	}
}

func Interval(i time.Duration) Config {
	return func(p *Policy) {
		p.interval = i
		p.nextInterval = i
	}
}

func MaxInterval(mi time.Duration) Config {
	return func(p *Policy) {
		p.maxInterval = mi
	}
}

func MaxJitterInterval(ji time.Duration) Config {
	return func(p *Policy) {
		p.maxJitterInterval = ji
	}
}

func RegularInterval(ri time.Duration) Config {
	return func(p *Policy) {
		p.regularInterval = ri
	}
}

func Multiplier(m float64) Config {
	return func(p *Policy) {
		p.multiplier = m
	}
}

func RandomFactor(rf float64) Config {
	return func(p *Policy) {
		p.randomFactor = rf
	}
}

func RetryType(rt retryType) Config {
	return func(p *Policy) {
		p.retryType = rt
	}
}

func (p *Policy) isRetryableError(resultError error, retryableErrors func(err error) bool) bool {
	return retryableErrors(resultError)
}
