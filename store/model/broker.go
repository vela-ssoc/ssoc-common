package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type Broker struct {
	ID          bson.ObjectID   `bson:"_id,omitempty"          json:"id"`
	Name        string          `bson:"name"                   json:"name"`
	Secret      string          `bson:"secret"                 json:"-"`
	Exposes     ExposeAddresses `bson:"exposes"                json:"exposes"`
	Status      bool            `bson:"status"                 json:"status"`
	TunnelStat  *TunnelStat     `bson:"tunnel_stat,omitempty"  json:"tunnel_stat,omitzero"`
	ExecuteStat *ExecuteStat    `bson:"execute_stat,omitempty" json:"execute_stat,omitzero"`
	CreatedAt   time.Time       `bson:"created_at,omitempty"   json:"created_at"`
	UpdatedAt   time.Time       `bson:"updated_at,omitempty"   json:"updated_at"`
}

func (Broker) CollectionName() string { return "broker" }
