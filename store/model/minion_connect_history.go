package model

import "go.mongodb.org/mongo-driver/v2/bson"

type MinionConnectHistory struct {
	ID          bson.ObjectID     `bson:"_id,omitempty"    json:"id"`
	MinionID    bson.ObjectID     `bson:"minion_id"        json:"minion_id"`
	Inet        string            `bson:"inet,omitempty"   json:"inet,omitzero"`
	Goos        string            `bson:"goos,omitempty"   json:"goos,omitzero"`
	Goarch      string            `bson:"goarch,omitempty" json:"goarch,omitzero"`
	ExecuteStat ExecuteStat       `bson:"execute_stat"     json:"execute_stat"`
	TunnelStat  TunnelStatHistory `bson:"tunnel_stat"      json:"tunnel_stat"`
}

func (MinionConnectHistory) CollectionName() string { return "minion_connect_history" }
