package netmux

import (
	"context"
	"errors"
	"net"
	"sync/atomic"

	"golang.org/x/net/quic"
)

// NewQUIC 创建一个基于 quic 的多路复用流。
// 服务端侧 endpoint 可以为 nil。
//
//	FIXME: QUIC 模式在 windows 下存在 reset stream 的问题(2026-01-14)，原因未知，生产环境不建议使用。
//		生产环境若想使用 quic，建议使用 https://github.com/quic-go/quic-go.
func NewQUIC(parent context.Context, conn *quic.Conn, endpoint *quic.Endpoint) Muxer {
	if parent == nil {
		parent = context.Background()
	}

	return &quicMUX{
		conn:    conn,
		end:     endpoint,
		traffic: new(trafficStat),
		parent:  parent,
	}
}

type quicMUX struct {
	conn    *quic.Conn
	end     *quic.Endpoint
	traffic *trafficStat
	closed  atomic.Bool
	parent  context.Context
}

func (m *quicMUX) Accept() (net.Conn, error) {
	return m.newConn(m.conn.AcceptStream(m.parent))
}

func (m *quicMUX) Open(ctx context.Context) (net.Conn, error) {
	return m.newConn(m.conn.NewStream(ctx))
}

func (m *quicMUX) Close() error {
	if !m.closed.CompareAndSwap(false, true) {
		return net.ErrClosed
	}

	var errs []error
	if err := m.conn.Close(); err != nil {
		errs = append(errs, err)
	}
	if end := m.end; end != nil {
		err := end.Close(m.parent)
		errs = append(errs, err)
	}

	return errors.Join(errs...)
}

func (m *quicMUX) Addr() net.Addr            { return net.UDPAddrFromAddrPort(m.conn.LocalAddr()) }
func (m *quicMUX) RemoteAddr() net.Addr      { return net.UDPAddrFromAddrPort(m.conn.RemoteAddr()) }
func (m *quicMUX) IsClosed() bool            { return m.closed.Load() }
func (m *quicMUX) Protocol() string          { return "quic" }
func (m *quicMUX) Traffic() (uint64, uint64) { return m.traffic.load() }

func (m *quicMUX) newConn(stm *quic.Stream, err error) (net.Conn, error) {
	if err != nil {
		return nil, err
	}
	conn := &quicConn{stm: stm, mst: m}

	return conn, nil
}
