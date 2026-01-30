package muxserver

import (
	"context"
	"errors"
	"net"
	"strings"

	"github.com/vela-ssoc/ssoc-proto/muxproto"
)

func NewMixedDialer(mux muxproto.MUXOpener, hub Huber, sys muxproto.Dialer) muxproto.Dialer {
	return &mixedDialer{
		mux: mux,
		hub: hub,
		sys: sys,
	}
}

type mixedDialer struct {
	mux muxproto.MUXOpener
	hub Huber
	sys muxproto.Dialer
}

func (m *mixedDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	host, _, err := net.SplitHostPort(address)
	if err != nil {
		return nil, err
	}

	if m.mux != nil && host == m.mux.Host() {
		return m.mux.Open(ctx)
	}

	if m.hub != nil {
		_, domain, found := strings.Cut(host, ".")
		if found && domain == m.hub.Domain() {
			peer := m.hub.Get(host)
			if peer == nil {
				return nil, &net.OpError{
					Op:   "dial",
					Net:  network,
					Addr: &net.UnixAddr{Net: network, Name: address},
					Err:  errors.New("节点不存在或已经离线"),
				}
			}

			mux := peer.MUX()

			return mux.Open(ctx)
		}
	}

	if m.sys != nil {
		return m.sys.DialContext(ctx, network, address)
	}

	return nil, &net.OpError{
		Op:     "dial",
		Net:    address,
		Source: &net.UnixAddr{Net: network, Name: address},
		Err:    errors.New("没有找到合适的拨号器"),
	}
}
