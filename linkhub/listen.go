package linkhub

import (
	"net"

	"github.com/xtaci/smux"
)

func NewSMUXListener(sess *smux.Session) net.Listener {
	return &smuxListener{sess: sess}
}

type smuxListener struct {
	sess *smux.Session
}

func (sl *smuxListener) Accept() (net.Conn, error) {
	stm, err := sl.sess.AcceptStream()
	if err != nil {
		return nil, err
	}

	return stm, nil
}

func (sl *smuxListener) Close() error {
	return sl.sess.Close()
}

func (sl *smuxListener) Addr() net.Addr {
	return sl.sess.LocalAddr()
}
