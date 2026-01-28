package linkhub

import "github.com/xtaci/smux"

type Info struct {
	ID   int64  `json:"id"`
	Inet string `json:"inet"`
	Host string `json:"host"`
}

type Peer interface {
	Info() Info
	Muxer() *smux.Session
}

func NewPeer(id int64, inet string, sess *smux.Session) Peer {
	return &tunnelPeer{
		id:   id,
		host: formatID(id),
		inet: inet,
		sess: sess,
	}
}

type tunnelPeer struct {
	id   int64
	host string
	inet string
	sess *smux.Session
}

func (tp *tunnelPeer) Info() Info {
	return Info{
		ID:   tp.id,
		Inet: tp.inet,
		Host: tp.host,
	}
}

func (tp *tunnelPeer) Muxer() *smux.Session {
	return tp.sess
}
