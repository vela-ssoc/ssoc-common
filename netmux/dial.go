package netmux

import (
	"context"
	"crypto/tls"
	"errors"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

// Dialer 通道拨号。
//
// var d muxconn.Dialer
// ctx := context.Background()
// d.DialContext(ctx, []string{"server.example.com"})
type Dialer struct {
	// Websocket 拨号器，不填写默认宽容模式。
	Websocket *websocket.Dialer

	// Path 接入点地址，不填写默认：/api/v1/tunnel。
	Path string

	// Timeout 每次建连时的超时时间。
	Timeout time.Duration
}

func (d Dialer) DialContext(ctx context.Context, addresses []string) (Muxer, error) {
	var errs []error
	wd := d.websocketDialer()
	for _, addr := range addresses {
		if mux, err := d.dialContext(ctx, wd, addr); err != nil {
			errs = append(errs, err)
		} else {
			return mux, nil
		}
	}
	if err := errors.Join(errs...); err != nil {
		return nil, err
	}

	return nil, errors.New("empty addresses")
}

func (d Dialer) dialContext(parent context.Context, wd *websocket.Dialer, addr string) (Muxer, error) {
	wssURL := &url.URL{Scheme: "wss", Host: addr, Path: d.Path}
	if wssURL.Path == "" {
		wssURL.Path = "/api/v1/tunnel"
	}

	strURL := wssURL.String()
	ctx, cancel := d.perContext(parent)
	defer cancel()

	ws, _, err := wd.DialContext(ctx, strURL, nil)
	if err != nil {
		return nil, err
	}
	conn := ws.NetConn()
	mux, err1 := NewSMUX(conn, nil, false)
	if err1 != nil {
		_ = ws.Close()
		return nil, err1
	}

	return mux, nil
}

func (d Dialer) websocketDialer() *websocket.Dialer {
	if dd := d.Websocket; dd != nil {
		return dd
	}

	return &websocket.Dialer{
		TLSClientConfig: &tls.Config{
			InsecureSkipVerify: true, // 宽容模式。
		},
	}
}

func (d Dialer) perContext(parent context.Context) (context.Context, context.CancelFunc) {
	if du := d.Timeout; du > 0 {
		return context.WithTimeout(parent, du)
	}

	return parent, func() {}
}
