package netmux

import (
	"context"
	"net"

	"github.com/xtaci/smux"
)

// NewSMUX 创建一个 smux 多路复用连接，请注意区分服务端/客户端侧。
func NewSMUX(conn net.Conn, cfg *smux.Config, serverSide bool) (Muxer, error) {
	var err error
	mux := &smuxMUX{traffic: new(trafficStat)}
	if serverSide {
		mux.sess, err = smux.Server(conn, cfg)
	} else {
		mux.sess, err = smux.Client(conn, cfg)
	}
	if err != nil {
		return nil, err
	}

	return mux, nil
}

type smuxMUX struct {
	sess    *smux.Session
	traffic *trafficStat
}

func (m *smuxMUX) Accept() (net.Conn, error) {
	stm, err := m.sess.AcceptStream()
	if err != nil {
		return nil, err
	}
	conn := m.newConn(stm)

	return conn, nil
}

func (m *smuxMUX) Open(context.Context) (net.Conn, error) {
	stm, err := m.sess.OpenStream()
	if err != nil {
		return nil, err
	}
	conn := m.newConn(stm)

	return conn, nil
}

func (m *smuxMUX) Close() error             { return m.sess.Close() }
func (m *smuxMUX) Addr() net.Addr           { return m.sess.LocalAddr() }
func (m *smuxMUX) RemoteAddr() net.Addr     { return m.sess.RemoteAddr() }
func (m *smuxMUX) IsClosed() bool           { return m.sess.IsClosed() }
func (m *smuxMUX) Protocol() string         { return "smux" }
func (m *smuxMUX) Traffic() (rx, tx uint64) { return m.traffic.load() }

func (m *smuxMUX) newConn(stm *smux.Stream) *smuxConn {
	return &smuxConn{stm: stm, mst: m}
}
