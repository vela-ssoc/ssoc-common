package model

import "time"

type TunnelStat struct {
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
	Inet       string   `json:"inet,omitzero"       bson:"inet,omitempty"`
	Goos       string   `json:"goos,omitzero"       bson:"goos,omitempty"`
	Goarch     string   `json:"goarch,omitzero"     bson:"goarch,omitempty"`
	Semver     string   `json:"semver,omitzero"     bson:"semver,omitempty"`
	PID        int      `json:"pid,omitzero"        bson:"pid,omitempty"`
	Args       []string `json:"args,omitzero"       bson:"args,omitempty"`
	Hostname   string   `json:"hostname,omitzero"   bson:"hostname,omitempty"`
	Workdir    string   `json:"workdir,omitzero"    bson:"workdir,omitempty"`
	Executable string   `json:"executable,omitzero" bson:"executable,omitempty"`
}

type TunnelStatHistory struct {
	ConnectedAt    time.Time     `bson:"connected_at,omitempty"    json:"connected_at,omitzero"`
	DisconnectedAt time.Time     `bson:"disconnected_at,omitempty" json:"disconnected_at,omitzero"`
	Library        TunnelLibrary `bson:"library,omitempty"         json:"library,omitzero"`
	LocalAddr      string        `bson:"local_addr,omitempty"      json:"local_addr,omitzero"`
	RemoteAddr     string        `bson:"remote_addr,omitempty"     json:"remote_addr,omitzero"`
	ReceiveBytes   uint64        `bson:"receive_bytes"             json:"receive_bytes,omitzero"`
	TransmitBytes  uint64        `bson:"transmit_bytes"            json:"transmit_bytes,omitzero"`
}
