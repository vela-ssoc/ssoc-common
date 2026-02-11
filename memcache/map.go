package memcache

import (
	"context"
	"sync"
)

type MapCache[K comparable, T any] struct {
	fun func(context.Context, K) (T, error)
	mtx sync.RWMutex
	idx map[K]*resultEntry[T]
}

func (c *MapCache[K, T]) Load(ctx context.Context, k K) (T, error) {
	c.mtx.RLock()
	r := c.idx[k]
	c.mtx.RUnlock()
	if r != nil {
		return r.t, r.e
	}

	return c.slowLoad(ctx, k)
}

func (c *MapCache[K, T]) Forget() {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	c.idx = nil
}

func (c *MapCache[K, T]) slowLoad(ctx context.Context, k K) (T, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if r := c.idx[k]; r != nil {
		return r.t, r.e
	}

	t, e := c.fun(ctx, k)
	if isTempError(e) {
		return t, e
	}

	r := &resultEntry[T]{t: t, e: e}
	if c.idx == nil {
		c.idx = make(map[K]*resultEntry[T], 4)
	}
	c.idx[k] = r

	return t, e
}
