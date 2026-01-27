package muxproto

import (
	"context"
	"net"

	"github.com/vela-ssoc/ssoc-common/muxtunnel/muxstream"
)

type Dialer interface {
	DialContext(ctx context.Context, network, address string) (net.Conn, error)
}

type MUXOpener interface {
	Host() string

	Open(context.Context) (net.Conn, error)
}

func NewMUXOpener(mux muxstream.Muxer, host string) MUXOpener {
	return &muxOpen{
		mux:  mux,
		host: host,
	}
}

type muxOpen struct {
	mux  muxstream.Muxer
	host string
}

func (m *muxOpen) Host() string                               { return m.host }
func (m *muxOpen) Open(ctx context.Context) (net.Conn, error) { return m.mux.Open(ctx) }
