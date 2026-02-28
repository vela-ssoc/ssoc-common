package model

import "go.mongodb.org/mongo-driver/v2/bson"

type BrokerConnectHistory struct {
	ID          bson.ObjectID     `bson:"_id,omitempty"    json:"id"`
	BrokerID    bson.ObjectID     `bson:"broker_id"        json:"broker_id"`
	Name        string            `bson:"name"             json:"name"`
	ExecuteStat ExecuteStat       `bson:"execute_stat"     json:"execute_stat,omitzero"`
	TunnelStat  TunnelStatHistory `bson:"tunnel_stat"      json:"tunnel_stat"`
}

func (BrokerConnectHistory) CollectionName() string { return "broker_connect_history" }
