package muxserver

import (
	"context"
	"net"

	"github.com/vela-ssoc/ssoc-proto/muxproto"
)

func NewMixedDialer(mux muxproto.MUXOpener) muxproto.Dialer {
	return &mixedDialer{mux: mux}
}

type mixedDialer struct {
	mux muxproto.MUXOpener
}

func (m *mixedDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if host == m.mux.Host() {
		return m.mux.Open(ctx)
	}

	return nil, net.UnknownNetworkError("未知的主机")
}
