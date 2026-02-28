package model

import "time"

type TunnelStat struct {
	Inet           string        `bson:"inet,omitempty"            json:"inet,omitzero"`
	ConnectedAt    time.Time     `bson:"connected_at,omitempty"    json:"connected_at,omitzero"`
	DisconnectedAt time.Time     `bson:"disconnected_at,omitempty" json:"disconnected_at,omitzero"`
	KeepaliveAt    time.Time     `bson:"keepalive_at,omitempty"    json:"keepalive_at,omitzero"`
	Library        TunnelLibrary `bson:"library,omitempty"         json:"library,omitzero"`
	LocalAddr      string        `bson:"local_addr,omitempty"      json:"local_addr,omitzero"`
	RemoteAddr     string        `bson:"remote_addr,omitempty"     json:"remote_addr,omitzero"`
	ReceiveBytes   uint64        `bson:"receive_bytes"             json:"receive_bytes"`
	TransmitBytes  uint64        `bson:"transmit_bytes"            json:"transmit_bytes"`
}

type TunnelLibrary struct {
	Name   string `bson:"name,omitempty"   json:"name,omitzero"`
	Module string `bson:"module,omitempty" json:"module,omitzero"`
}

type ExecuteStat struct {
	Goos       string   `bson:"goos,omitempty"       json:"goos,omitzero"`
	Goarch     string   `bson:"goarch,omitempty"     json:"goarch,omitzero"`
	Semver     string   `bson:"semver"               json:"semver"`
	Version    uint64   `bson:"version"              json:"version"`
	Unstable   bool     `bson:"unstable"             json:"unstable"` // 内测版本，主要用于 agent 节点。
	PID        int      `bson:"pid,omitempty"        json:"pid,omitzero"`
	Args       []string `bson:"args,omitempty"       json:"args,omitzero"`
	Hostname   string   `bson:"hostname,omitempty"   json:"hostname,omitzero"`
	Workdir    string   `bson:"workdir,omitempty"    json:"workdir,omitzero"`
	Executable string   `bson:"executable,omitempty" json:"executable,omitzero"`
}

type TunnelStatHistory struct {
	Inet           string        `bson:"inet"                      json:"inet,omitzero"`
	ConnectedAt    time.Time     `bson:"connected_at,omitempty"    json:"connected_at,omitzero"`
	DisconnectedAt time.Time     `bson:"disconnected_at,omitempty" json:"disconnected_at,omitzero"`
	ConnectSeconds uint64        `bson:"connect_seconds"           json:"connect_seconds,omitzero"` // 连接持续时长秒数
	Library        TunnelLibrary `bson:"library,omitempty"         json:"library,omitzero"`
	LocalAddr      string        `bson:"local_addr,omitempty"      json:"local_addr,omitzero"`
	RemoteAddr     string        `bson:"remote_addr,omitempty"     json:"remote_addr,omitzero"`
	ReceiveBytes   uint64        `bson:"receive_bytes"             json:"receive_bytes,omitzero"`
	TransmitBytes  uint64        `bson:"transmit_bytes"            json:"transmit_bytes,omitzero"`
}
