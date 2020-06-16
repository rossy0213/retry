package retry

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWithTest(t *testing.T) {
	t.Parallel()
	eb := DefaultExponentialBackoff()
	ctx := context.Background()
	be := withContext(ctx, eb)

	assert.EqualValues(t, ctx, be.Context())
}
