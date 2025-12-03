package preadtls

import (
	"bufio"
	"context"
	"io"
	"net"
	"sync/atomic"
	"time"
)

func NewListener(lis net.Listener, readTimeout time.Duration) *Listener {
	addr := lis.Addr()
	ln := &Listener{
		raw: lis,
		tcp: &chanListener{
			conn: make(chan net.Conn),
			addr: addr,
			stop: make(chan struct{}),
		},
		tls: &chanListener{
			conn: make(chan net.Conn),
			addr: addr,
			stop: make(chan struct{}),
		},
		timeout: readTimeout,
	}

	go func() { _ = ln.Accept() }()

	return ln
}

type Listener struct {
	raw     net.Listener
	tcp     *chanListener
	tls     *chanListener
	timeout time.Duration
	closed  atomic.Bool
}

func (pl *Listener) TCPListener() net.Listener {
	return pl.tcp
}

func (pl *Listener) TLSListener() net.Listener {
	return pl.tls
}

func (pl *Listener) Close() error {
	if pl.closed.CompareAndSwap(false, true) {
		err := pl.raw.Close()
		_ = pl.tcp.Close()
		_ = pl.tls.Close()

		return err
	}

	return net.ErrClosed
}

func (pl *Listener) Accept() error {
	var tempDelay time.Duration // how long to sleep on accept failure
	for {
		conn, err := pl.raw.Accept()
		if err != nil {
			if pl.closed.Load() {
				return net.ErrClosed
			}

			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				if tempDelay == 0 {
					tempDelay = 5 * time.Millisecond
				} else {
					tempDelay *= 2
				}
				if delay := 1 * time.Second; tempDelay > delay {
					tempDelay = delay
				}
				time.Sleep(tempDelay)
				continue
			}
			return err
		}

		pl.enqueue(conn)
	}
}

func (pl *Listener) enqueue(conn net.Conn) {
	if du := pl.timeout; du > 0 {
		now := time.Now()
		_ = conn.SetReadDeadline(now.Add(du))
	}
	buf := bufio.NewReader(conn)
	peek, err := buf.Peek(1)
	if err != nil {
		_ = conn.Close()
		return
	}

	pc := &peekConn{conn: conn, read: io.MultiReader(buf, conn)}
	if peek[0] == 0x16 { // TLS 首字符特征
		err = pl.tls.enqueue(pc, pl.timeout)
	} else {
		err = pl.tcp.enqueue(pc, pl.timeout)
	}
	if err != nil {
		_ = conn.Close()
	}
}

type chanListener struct {
	conn   chan net.Conn
	addr   net.Addr
	stop   chan struct{}
	closed atomic.Bool
}

func (cl *chanListener) Accept() (net.Conn, error) {
	select {
	case conn, ok := <-cl.conn:
		if ok {
			return conn, nil
		}
		return nil, net.ErrClosed
	case <-cl.stop:
		return nil, net.ErrClosed
	}
}

func (cl *chanListener) Close() error {
	if cl.closed.CompareAndSwap(false, true) {
		close(cl.stop)
		return nil
	}

	return net.ErrClosed
}

func (cl *chanListener) Addr() net.Addr {
	return cl.addr
}

func (cl *chanListener) enqueue(conn net.Conn, timeout time.Duration) error {
	if cl.closed.Load() {
		return net.ErrClosed
	}

	select {
	case <-cl.stop:
		return net.ErrClosed
	default:
	}

	var done <-chan time.Time
	if timeout > 0 {
		timer := time.NewTimer(timeout)
		defer timer.Stop()
		done = timer.C
	}

	select {
	case <-cl.stop:
		return net.ErrClosed
	case <-done:
		return context.DeadlineExceeded
	case cl.conn <- conn:
		return nil
	}
}
