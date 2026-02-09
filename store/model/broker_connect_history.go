package model

import "go.mongodb.org/mongo-driver/v2/bson"

type BrokerConnectHistory struct {
	ID         bson.ObjectID     `bson:"_id,omitempty"    json:"id"`
	Broker     bson.ObjectID     `bson:"broker_id"        json:"broker_id"`
	Name       string            `bson:"name"             json:"name"`
	Semver     string            `bson:"semver,omitempty" json:"semver,omitzero"`
	Inet       string            `bson:"inet,omitempty"   json:"inet,omitzero"`
	Goos       string            `bson:"goos,omitempty"   json:"goos,omitzero"`
	Goarch     string            `bson:"goarch,omitempty" json:"goarch,omitzero"`
	TunnelStat TunnelStatHistory `bson:"tunnel_stat"      json:"tunnel_stat"`
}

func (BrokerConnectHistory) CollectionName() string { return "broker_connect_history" }
