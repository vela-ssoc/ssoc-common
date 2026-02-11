package memcache

import (
	"context"
	"errors"
	"time"
)

type resultEntry[T any] struct {
	t       T
	e       error
	expires time.Time
}

func (e *resultEntry[T]) alive(at time.Time) bool {
	return !e.expires.Before(at)
}

func isTempError(err error) bool {
	return err != nil && (errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded))
}
