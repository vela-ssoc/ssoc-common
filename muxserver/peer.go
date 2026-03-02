package muxserver

import (
	"time"

	"github.com/vela-ssoc/ssoc-proto/muxconn"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Peer interface {
	ID() bson.ObjectID
	MUX() muxconn.Muxer
	Info() PeerInfo
	Host() string
}

type PeerInfo struct {
	Instance    string    `json:"instance"`
	Semver      string    `json:"semver"`
	Inet        string    `json:"inet"`
	Goos        string    `json:"goos"`
	Goarch      string    `json:"goarch"`
	Hostname    string    `json:"hostname"`
	ConnectedAt time.Time `json:"connected_at"`
}

type muxPeer struct {
	id   bson.ObjectID
	mux  muxconn.Muxer
	host string
	info PeerInfo
}

func (m *muxPeer) ID() bson.ObjectID  { return m.id }
func (m *muxPeer) MUX() muxconn.Muxer { return m.mux }
func (m *muxPeer) Info() PeerInfo     { return m.info }
func (m *muxPeer) Host() string       { return m.host }
