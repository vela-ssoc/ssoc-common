package muxproto

import "github.com/vela-ssoc/ssoc-common/muxtunnel/muxstream"

type Peer interface {
	ID() int64
	Muxer() muxstream.Muxer
	Host() string
	Info() Info
}

type Info struct {
	Name      string `json:"name"`       // 名字（broker 节点）
	Semver    string `json:"semver"`     // 语义化版本号
	Inet      string `json:"inet"`       // 出口 IP
	Goos      string `json:"goos"`       // 操作系统
	Goarch    string `json:"goarch"`     // 位数
	Hostname  string `json:"hostname"`   // 主机名
	MachineID string `json:"machine_id"` // agent 节点机器码
}

type suffixPeer struct {
	id   int64
	mux  muxstream.Muxer
	host string
	info Info
}

func (p *suffixPeer) ID() int64              { return p.id }
func (p *suffixPeer) Muxer() muxstream.Muxer { return p.mux }
func (p *suffixPeer) Host() string           { return p.host }
func (p *suffixPeer) Info() Info             { return p.info }
