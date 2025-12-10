package linkhub

import (
	"context"
	"net"
	"strings"
)

type Dialer interface {
	DialContext(ctx context.Context, network, addr string) (net.Conn, error)
}

func NewSelectDialer(fallback Dialer, matches ...Dialer) Dialer {
	return &selectDialer{
		matches:  matches,
		fallback: fallback,
	}
}

type selectDialer struct {
	matches  []Dialer
	fallback Dialer
}

func (md *selectDialer) DialContext(ctx context.Context, network, addr string) (net.Conn, error) {
	for _, d := range md.matches {
		// 当 conn 和 err 都为 nil 时，表示模式不匹配，尝试下一个 dialer
		conn, err := d.DialContext(ctx, network, addr)
		if conn != nil {
			return conn, nil
		} else if err != nil {
			return nil, err
		}
	}

	if d := md.fallback; d != nil {
		return d.DialContext(ctx, network, addr)
	}

	return nil, &net.DNSError{
		Err:        "没有找到合适的 dialer",
		Name:       addr,
		Server:     "matches-dialer",
		IsNotFound: true,
	}
}

func NewSuffixDialer(suffix string, hub Huber) Dialer {
	return &suffixDialer{
		suffix: suffix,
		hub:    hub,
	}
}

type suffixDialer struct {
	suffix string
	hub    Huber
}

func (sd *suffixDialer) DialContext(_ context.Context, _, addr string) (net.Conn, error) {
	host, _, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, err
	}
	peerHost, found := strings.CutSuffix(host, sd.suffix)
	if !found {
		return nil, nil
	}

	peer := sd.hub.Get(peerHost)
	if peer == nil {
		return nil, &net.AddrError{
			Err:  "节点未上线",
			Addr: addr,
		}
	}

	mux := peer.Muxer()
	stm, err := mux.OpenStream()
	if err != nil {
		return nil, err
	}

	return stm, nil
}
