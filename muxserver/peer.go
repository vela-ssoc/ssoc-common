package muxserver

import "github.com/vela-ssoc/ssoc-proto/muxconn"

type Peer interface {
	ID() int64
	MUX() muxconn.Muxer
	Info() PeerInfo
	Host() string
}

type PeerInfo struct {
	Name     string `json:"name"`
	Semver   string `json:"semver"`
	Inet     string `json:"inet"`
	Goos     string `json:"goos"`
	Goarch   string `json:"goarch"`
	Hostname string `json:"hostname"`
}

type muxPeer struct {
	id   int64
	mux  muxconn.Muxer
	host string
	info PeerInfo
}

func (m *muxPeer) ID() int64          { return m.id }
func (m *muxPeer) MUX() muxconn.Muxer { return m.mux }
func (m *muxPeer) Info() PeerInfo     { return m.info }
func (m *muxPeer) Host() string       { return m.host }
