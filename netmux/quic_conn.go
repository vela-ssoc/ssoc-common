package netmux

import (
	"context"
	"net"
	"time"

	"golang.org/x/net/quic"
)

type quicConn struct {
	stm *quic.Stream
	mst *quicMUX
}

func (c *quicConn) Read(b []byte) (int, error) {
	n, err := c.stm.Read(b)
	c.mst.traffic.incrTX(n)

	return n, err
}

func (c *quicConn) Write(b []byte) (int, error) {
	n, err := c.stm.Write(b)
	c.mst.traffic.incrTX(n)

	return n, err
}

func (c *quicConn) Close() error {
	return c.stm.Close()
}

func (c *quicConn) LocalAddr() net.Addr {
	return c.mst.Addr()
}

func (c *quicConn) RemoteAddr() net.Addr {
	return c.mst.RemoteAddr()
}

func (c *quicConn) SetDeadline(t time.Time) error {
	ctx := c.withContext(t)
	c.stm.SetReadContext(ctx)
	c.stm.SetWriteContext(ctx)

	return nil
}

func (c *quicConn) SetReadDeadline(t time.Time) error {
	c.stm.SetReadContext(c.withContext(t))

	return nil
}

func (c *quicConn) SetWriteDeadline(t time.Time) error {
	c.stm.SetWriteContext(c.withContext(t))

	return nil
}

func (c *quicConn) withContext(t time.Time) context.Context {
	if t.IsZero() {
		return context.Background()
	}

	//goland:noinspection GoVetLostCancel
	ctx, _ := context.WithDeadline(context.Background(), t)

	return ctx
}
