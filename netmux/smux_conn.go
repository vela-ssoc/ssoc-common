package netmux

import (
	"net"
	"time"

	"github.com/xtaci/smux"
)

type smuxConn struct {
	stm *smux.Stream
	mst *smuxMUX
}

func (s *smuxConn) Read(b []byte) (int, error) {
	n, err := s.stm.Read(b)
	s.mst.traffic.incrRX(n)

	return n, err
}

func (s *smuxConn) Write(b []byte) (int, error) {
	n, err := s.stm.Write(b)
	s.mst.traffic.incrTX(n)

	return n, err
}

func (s *smuxConn) Close() error                       { return s.stm.Close() }
func (s *smuxConn) LocalAddr() net.Addr                { return s.stm.LocalAddr() }
func (s *smuxConn) RemoteAddr() net.Addr               { return s.stm.RemoteAddr() }
func (s *smuxConn) SetDeadline(t time.Time) error      { return s.stm.SetDeadline(t) }
func (s *smuxConn) SetReadDeadline(t time.Time) error  { return s.stm.SetReadDeadline(t) }
func (s *smuxConn) SetWriteDeadline(t time.Time) error { return s.stm.SetWriteDeadline(t) }
