package memcache

import (
	"context"
	"sync"
	"sync/atomic"
)

type Cache[T any] struct {
	fun func(context.Context) (T, error)
	mtx sync.Mutex
	ptr atomic.Pointer[resultEntry[T]]
}

func NewCache[T any](fun func(context.Context) (T, error)) *Cache[T] {
	return &Cache[T]{
		fun: fun,
	}
}

func (c *Cache[T]) Load(ctx context.Context) (T, error) {
	if r := c.ptr.Load(); r != nil {
		return r.t, r.e
	}

	return c.slowLoad(ctx)
}

func (c *Cache[T]) Forget() (T, error) {
	if r := c.ptr.Swap(nil); r != nil {
		return r.t, r.e
	}

	var t T

	return t, nil
}

func (c *Cache[T]) slowLoad(ctx context.Context) (T, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if d := c.ptr.Load(); d != nil {
		return d.t, d.e
	}

	t, e := c.fun(ctx)
	if isTempError(e) {
		return t, e
	}

	d := &resultEntry[T]{t: t, e: e}
	c.ptr.Store(d)

	return t, e
}
