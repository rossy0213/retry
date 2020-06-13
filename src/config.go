package retry4go

import "time"

type Config func(*exponentialBackoff)

func defaultCheckRetryable() checkRetryable {
	return func(error) bool {
		return false
	}
}

func MaxRetryTimes(mt uint) Config {
	return func(eb *exponentialBackoff) {
		eb.maxRetryTimes = mt
	}
}

func CheckRetryable(cr checkRetryable) Config {
	return func(eb *exponentialBackoff) {
		eb.checkRetryable = cr
	}
}

func Interval(i time.Duration) Config {
	return func(eb *exponentialBackoff) {
		eb.interval = i
		eb.nextInterval = i
	}
}

func MaxInterval(mi time.Duration) Config {
	return func(eb *exponentialBackoff) {
		eb.maxInterval = mi
	}
}

func MaxJitterInterval(ji time.Duration) Config {
	return func(eb *exponentialBackoff) {
		eb.maxJitterInterval = ji
	}
}

func Multiplier(m float64) Config {
	return func(eb *exponentialBackoff) {
		eb.multiplier = m
	}
}

func MaxElapsedTime(et time.Duration) Config {
	return func(eb *exponentialBackoff) {
		eb.maxElapsedTime = et
	}
}
