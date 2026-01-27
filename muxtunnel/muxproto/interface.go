package muxproto

import (
	"context"
	"time"

	"github.com/vela-ssoc/ssoc-common/muxtunnel/muxstream"
)

type MUXAccepter interface {
	AcceptMUX(muxstream.Muxer)
}

type ConfigLoader[T any] interface {
	LoadConfig(ctx context.Context) (*T, error)
}

type ClientHooker interface {
	// Disconnected 通道掉线。
	Disconnected(mux muxstream.Muxer, err error)

	// Reconnected 通道掉线后重连成功。
	Reconnected(mux muxstream.Muxer)

	// OnExit 连接通道遇到不可重试的错误无法继续保持连接，
	// 通常原因是 context 取消。
	OnExit(err error)
}

type HubHooker interface {
	// OnAuthFailed 当节点认证失败回调接口。
	OnAuthFailed(ctx context.Context, mux muxstream.Muxer, connAt time.Time, err error)

	// OnConnected 认证连接成功后回调接口。
	OnConnected(ctx context.Context, p Peer, connAt time.Time)

	// OnDisconnected 认证通过后断开连接时的回调接口。
	OnDisconnected(ctx context.Context, p Peer, connectAt, disconnectAt time.Time)
}
