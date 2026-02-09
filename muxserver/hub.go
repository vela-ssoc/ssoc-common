package muxserver

import (
	"sync"

	"github.com/vela-ssoc/ssoc-proto/muxconn"
	"github.com/vela-ssoc/ssoc-proto/muxproto"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Huber interface {
	Put(id bson.ObjectID, mux muxconn.Muxer, inf PeerInfo) Peer
	Get(host string) Peer
	Del(host string) Peer
	GetID(id bson.ObjectID) Peer
	DelID(id bson.ObjectID) Peer
	Peers() []Peer
	Domain() string
}

func NewBrokerHub() Huber {
	return &mapHub{
		domain: muxproto.BrokerDomain,
		peers:  make(map[string]Peer, 32),
	}
}

type mapHub struct {
	domain string
	mutex  sync.RWMutex
	peers  map[string]Peer
}

func (m *mapHub) Put(id bson.ObjectID, mux muxconn.Muxer, inf PeerInfo) Peer {
	host := m.resolveHostname(id)
	peer := &muxPeer{
		id:   id,
		mux:  mux,
		host: host,
		info: inf,
	}

	m.mutex.Lock()
	defer m.mutex.Unlock()

	if _, exists := m.peers[host]; exists {
		return nil
	}
	m.peers[host] = peer

	return peer
}

func (m *mapHub) Get(host string) Peer {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	return m.peers[host]
}

func (m *mapHub) Del(host string) Peer {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	peer := m.peers[host]
	if peer != nil {
		delete(m.peers, host)
	}

	return peer
}

func (m *mapHub) GetID(id bson.ObjectID) Peer {
	host := m.resolveHostname(id)
	return m.Get(host)
}

func (m *mapHub) DelID(id bson.ObjectID) Peer {
	host := m.resolveHostname(id)
	return m.Del(host)
}

func (m *mapHub) Peers() []Peer {
	m.mutex.RLock()
	defer m.mutex.RUnlock()

	res := make([]Peer, 0, len(m.peers))
	for _, peer := range m.peers {
		res = append(res, peer)
	}

	return res
}

func (m *mapHub) Domain() string {
	return m.domain
}

func (m *mapHub) resolveHostname(id bson.ObjectID) string {
	return muxproto.ResolveHostname(id.Hex(), m.domain)
}
