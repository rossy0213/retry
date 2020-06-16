package retry

import (
	"time"
)

const Stop time.Duration = -1

// Backoff is to accommodate a variety of backoff.
type Backoff interface {
	Next() time.Duration
}
