package linkhub

import "context"

func WithContext(parent context.Context, p Peer) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, httpContextKey, p)
}

func FromContext(ctx context.Context) Peer {
	if ctx == nil {
		return nil
	}

	p, _ := ctx.Value(httpContextKey).(Peer)

	return p
}

var httpContextKey = contextKey{name: "tunnel-http-context-key"}

type contextKey struct {
	name string
}
