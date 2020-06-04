package retry4go

import "time"

const (
	RegularRetry = retryType(1)
	BackOffRetry = retryType(2)
)

type Retry interface {
	Next() time.Duration
}
