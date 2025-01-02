package raas

import (
	"context"
	"errors"
)

func IsContextCanceled(ctx context.Context) bool {
	return errors.Is(ctx.Err(), context.Canceled)
}
