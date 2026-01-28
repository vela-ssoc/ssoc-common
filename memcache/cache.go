package memcache

import (
	"context"
	"sync"
)

type Cache[V, E any] interface {
	Load(context.Context) (V, E)
	Forget() (V, E)
}

func NewCache[V, E any](fn func(context.Context) (V, E)) Cache[V, E] {
	return &cache[V, E]{
		fn: fn,
	}
}

type cache[V, E any] struct {
	fn  func(context.Context) (V, E)
	mu  sync.RWMutex
	ent *entry2[V, E]
}

func (ch *cache[V, E]) Load(ctx context.Context) (V, E) {
	ch.mu.RLock()
	ent := ch.ent
	ch.mu.RUnlock()
	if ent != nil {
		return ent.load()
	}

	return ch.slowLoad(ctx)
}

func (ch *cache[V, E]) Forget() (v V, e E) {
	ch.mu.Lock()
	if ent := ch.ent; ent != nil {
		v, e = ent.load()
		ch.ent = nil
	}
	ch.mu.Unlock()

	return
}

func (ch *cache[V, E]) slowLoad(ctx context.Context) (V, E) {
	ch.mu.Lock()
	defer ch.mu.Unlock()

	if ent := ch.ent; ent != nil {
		return ent.load()
	}

	v, e := ch.fn(ctx)
	ch.ent = &entry2[V, E]{v: v, e: e}

	return v, e
}
