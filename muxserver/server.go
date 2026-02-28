package muxserver

import (
	"context"
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/vela-ssoc/ssoc-proto/muxconn"
)

type MUXAccepter interface {
	// AcceptMUX v2 接口。
	AcceptMUX(mux muxconn.Muxer) error

	// AcceptTCP 兼容 v1 agent 上线的接口。
	// 主要由 broker 方实现，manager 无需实现该接口。
	//
	// Deprecated: use AcceptMUX.
	AcceptTCP(conn net.Conn) error
}

type BootLoader[T any] interface {
	LoadBoot(ctx context.Context) (*T, error)
}

type ConnectNotifier interface {
	OnAuthFailed(ctx context.Context, mux muxconn.Muxer, connectAt time.Time, err error)

	OnConnected(ctx context.Context, inf PeerInfo, connAt time.Time)

	OnDisconnected(ctx context.Context, inf PeerInfo, connectAt, disconnectAt time.Time)
}

type OnceCloser interface {
	Close()
	Closed() bool
}

func NewOnceCloser(c io.Closer) OnceCloser {
	return &onceCloser{c: c}
}

type onceCloser struct {
	f atomic.Bool
	c io.Closer
}

func (o *onceCloser) Close() {
	if o.f.CompareAndSwap(false, true) {
		_ = o.c.Close()
	}
}

func (o *onceCloser) Closed() bool {
	return o.f.Load()
}
