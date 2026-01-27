package muxproto

import (
	"strconv"
	"sync"

	"github.com/vela-ssoc/ssoc-common/muxtunnel/muxstream"
)

type Huber interface {
	Put(id int64, mux muxstream.Muxer, inf Info) (peer Peer, putOK bool)
	Get(host string) Peer
	GetID(id int64) Peer
	Del(host string) Peer
	DelID(id int64) Peer
	Peers() []Peer
	Domain() string
}

func NewAgentHub() Huber {
	return &vDomainHub{
		domain: AgentHost,
		peers:  make(map[string]Peer, 1024),
	}
}

func NewBrokerHub() Huber {
	return &vDomainHub{
		domain: BrokerHost,
		peers:  make(map[string]Peer, 16),
	}
}

type vDomainHub struct {
	domain string          // 内部域名后缀
	mutex  sync.RWMutex    // peers 读写锁
	peers  map[string]Peer // peers 池
}

func (h *vDomainHub) Put(id int64, mux muxstream.Muxer, inf Info) (Peer, bool) {
	host := resolveHost(id, h.domain)
	peer := &suffixPeer{
		id:   id,
		mux:  mux,
		host: host,
		info: inf,
	}

	h.mutex.Lock()
	defer h.mutex.Unlock()

	if _, exists := h.peers[host]; exists {
		return nil, false
	}
	h.peers[host] = peer

	return peer, true
}

func (h *vDomainHub) Get(host string) Peer {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	return h.peers[host]
}

func (h *vDomainHub) GetID(id int64) Peer {
	host := resolveHost(id, h.domain)
	return h.Get(host)
}

func (h *vDomainHub) Del(host string) Peer {
	h.mutex.Lock()
	defer h.mutex.Unlock()

	peer, exists := h.peers[host]
	if exists {
		delete(h.peers, host)
	}

	return peer
}

func (h *vDomainHub) DelID(id int64) Peer {
	host := resolveHost(id, h.domain)
	return h.Del(host)
}

func (h *vDomainHub) Peers() []Peer {
	h.mutex.RLock()
	defer h.mutex.RUnlock()

	peers := make([]Peer, 0, len(h.peers))
	for _, peer := range h.peers {
		peers = append(peers, peer)
	}

	return peers
}

func (h *vDomainHub) Domain() string { return h.domain }

func resolveHost(id int64, suffix string) string {
	sid := strconv.FormatInt(id, 10)
	return sid + "." + suffix
}
