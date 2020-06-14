package retry4go

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTest(t *testing.T) {
	eb := DefaultExponentialBackoff()
	ctx := context.Background()
	be := withContext(ctx, eb)

	assert.EqualValues(t, ctx, be.Context())
}
