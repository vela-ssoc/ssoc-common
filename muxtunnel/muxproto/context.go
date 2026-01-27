package muxproto

import "context"

type contextKey struct {
	name string
}

var peerContextKey = &contextKey{name: "http-peer-context"}

func WithContext(parent context.Context, peer Peer) context.Context {
	if parent == nil {
		parent = context.Background()
	}

	return context.WithValue(parent, peerContextKey, peer)
}

func FromContext(ctx context.Context) Peer {
	if ctx == nil {
		return nil
	}

	if peer, ok := ctx.Value(peerContextKey).(Peer); ok {
		return peer
	}

	return nil
}
