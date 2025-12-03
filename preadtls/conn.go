package preadtls

import (
	"io"
	"net"
	"time"
)

type peekConn struct {
	conn net.Conn
	read io.Reader
}

func (pc *peekConn) Read(b []byte) (int, error) {
	return pc.read.Read(b)
}

func (pc *peekConn) Write(b []byte) (int, error) {
	return pc.conn.Write(b)
}

func (pc *peekConn) Close() error {
	return pc.conn.Close()
}

func (pc *peekConn) LocalAddr() net.Addr {
	return pc.conn.LocalAddr()
}

func (pc *peekConn) RemoteAddr() net.Addr {
	return pc.conn.RemoteAddr()
}

func (pc *peekConn) SetDeadline(t time.Time) error {
	return pc.conn.SetDeadline(t)
}

func (pc *peekConn) SetReadDeadline(t time.Time) error {
	return pc.conn.SetReadDeadline(t)
}

func (pc *peekConn) SetWriteDeadline(t time.Time) error {
	return pc.conn.SetWriteDeadline(t)
}
