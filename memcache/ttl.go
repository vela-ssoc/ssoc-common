package memcache

import (
	"context"
	"sync"
	"sync/atomic"
	"time"
)

type TTLCache[T any] struct {
	ttl time.Duration
	fun func(context.Context) (T, error)
	mtx sync.Mutex
	ptr atomic.Pointer[resultEntry[T]]
}

func NewTTLCache[T any](ttl time.Duration, fun func(context.Context) (T, error)) *TTLCache[T] {
	return &TTLCache[T]{
		ttl: ttl,
		fun: fun,
	}
}

func (c *TTLCache[T]) Load(ctx context.Context) (T, error) {
	now := time.Now()
	if r := c.ptr.Load(); r != nil && r.alive(now) {
		return r.t, r.e
	}

	return c.slowLoad(ctx, now)
}

func (c *TTLCache[T]) Forget() {
	c.ptr.Store(nil)
}

func (c *TTLCache[T]) slowLoad(ctx context.Context, now time.Time) (T, error) {
	c.mtx.Lock()
	defer c.mtx.Unlock()

	if r := c.ptr.Load(); r != nil && r.alive(now) {
		return r.t, r.e
	}

	t, e := c.fun(ctx)
	if c.ttl <= 0 || isTempError(e) { // TTL <= 0 无需缓存，每次都缓存穿透。
		return t, e
	}

	now = time.Now()
	d := &resultEntry[T]{t: t, e: e, expires: now.Add(c.ttl)}
	c.ptr.Store(d)

	return t, e
}
