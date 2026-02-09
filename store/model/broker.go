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
	Config      BrokerConfig    `bson:"config"                 json:"config"`
	Status      bool            `bson:"status"                 json:"status"`
	TunnelStat  *TunnelStat     `bson:"tunnel_stat,omitempty"  json:"tunnel_stat,omitzero"`
	ExecuteStat *ExecuteStat    `bson:"execute_stat,omitempty" json:"execute_stat,omitzero"`
	CreatedAt   time.Time       `bson:"created_at,omitempty"   json:"created_at"`
	UpdatedAt   time.Time       `bson:"updated_at,omitempty"   json:"updated_at"`
}

func (Broker) CollectionName() string { return "broker" }

type BrokerConfig struct {
	Server BrokerServerConfig `bson:"server" json:"server"`
	Logger BrokerLoggerConfig `bson:"logger" json:"logger"`
}

type BrokerLoggerConfig struct {
	Level      string `bson:"level"      json:"level"      validate:"omitempty,oneof=DEBUG INFO WARN ERROR"`
	Console    bool   `bson:"console"    json:"console"`
	Filename   string `bson:"filename"   json:"filename"   validate:"lte=255"`
	MaxSize    int    `bson:"maxsize"    json:"maxsize"    validate:"gte=0"`
	MaxAge     int    `bson:"maxage"     json:"maxage"     validate:"gte=0"`
	MaxBackups int    `bson:"maxbackups" json:"maxbackups" validate:"gte=0"`
	LocalTime  bool   `bson:"localtime"  json:"localtime"`
	Compress   bool   `bson:"compress"   json:"compress"`
}

type BrokerServerConfig struct {
	Addr   string            `bson:"addr"   json:"addr"   validate:"lte=100"`
	Static map[string]string `bson:"static" json:"static" validate:"lte=10"`
}
