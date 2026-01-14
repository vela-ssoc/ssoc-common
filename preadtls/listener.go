package preadtls

import (
	"bufio"
	"context"
	"io"
	"net"
	"sync"
	"sync/atomic"
	"time"
)

func ListenTCP(addr string, readTimeout time.Duration) (*Listener, error) {
	raw, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	master := &Listener{
		raw:     raw,
		tcp:     nil,
		tls:     nil,
		timeout: readTimeout,
	}
	laddr := raw.Addr()
	tcp := &chanListener{
		master: master,
		conn:   make(chan net.Conn),
		addr:   laddr,
		stop:   make(chan struct{}),
		closed: atomic.Bool{},
	}
	tls := &chanListener{
		master: master,
		conn:   make(chan net.Conn),
		addr:   laddr,
		stop:   make(chan struct{}),
		closed: atomic.Bool{},
	}
	subs := map[*chanListener]struct{}{tcp: {}, tls: {}}
	master.tcp = tcp
	master.tls = tls
	master.subs = subs

	go func() { _ = master.Accept() }()

	return master, nil
}

type Listener struct {
	raw     net.Listener
	tcp     *chanListener
	tls     *chanListener
	timeout time.Duration // 首字节报文等待超时时间。
	closed  atomic.Bool
	mutex   sync.Mutex
	subs    map[*chanListener]struct{}
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
	timeout := pl.timeout
	if timeout > 0 {
		_ = conn.SetReadDeadline(time.Now().Add(timeout))
	}
	buf := bufio.NewReader(conn)
	peek, err := buf.Peek(1)
	if err != nil {
		_ = conn.Close()
		return
	}
	if timeout > 0 {
		_ = conn.SetReadDeadline(time.Time{})
	}

	pc := &peekConn{conn: conn, read: io.MultiReader(buf, conn)}
	if peek[0] == 0x16 { // TLS 首字符特征
		err = pl.tls.enqueue(pc, timeout)
	} else {
		err = pl.tcp.enqueue(pc, timeout)
	}
	if err != nil {
		_ = conn.Close()
	}
}

// close 当子 listener 关闭时会触发此函数。
// 确保每个子 listener 只会触发一次
func (pl *Listener) close(sub *chanListener) {
	pl.mutex.Lock()
	delete(pl.subs, sub)
	zero := len(pl.subs) == 0
	pl.mutex.Unlock()

	if zero { // 所有的子 listener 关闭后，主 listener 就应该被关闭。
		_ = pl.raw.Close()
	}
}

type chanListener struct {
	master *Listener
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
		cl.master.close(cl)
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
