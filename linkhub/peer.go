package linkhub

import "github.com/xtaci/smux"

type Handler interface {
	Handle(*smux.Session)
}

type Peer interface {
	ID() int64
	Host() string
	Muxer() *smux.Session
}

func NewPeer(id int64, sess *smux.Session) Peer {
	return &tunnelPeer{
		id:   id,
		host: formatID(id),
		sess: sess,
	}
}

type tunnelPeer struct {
	id   int64
	host string
	sess *smux.Session
}

func (tp *tunnelPeer) ID() int64 {
	return tp.id
}

func (tp *tunnelPeer) Host() string {
	return tp.host
}

func (tp *tunnelPeer) Muxer() *smux.Session {
	return tp.sess
}
