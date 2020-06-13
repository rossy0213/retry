package retry4go

import (
	"time"
)

const Stop time.Duration = -1

type Backoff interface {
	Next() time.Duration
}
