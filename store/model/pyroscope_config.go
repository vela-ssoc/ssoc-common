package model

import (
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

type PyroscopeConfig struct {
	ID        bson.ObjectID `bson:"_id,omitempty"        json:"id"`
	Name      string        `bson:"name"                 json:"name"`
	Enabled   bool          `bson:"enabled"              json:"enabled"`
	URL       string        `bson:"url"                  json:"url"`
	Username  string        `bson:"username"             json:"username"`
	Password  string        `bson:"password"             json:"password"`
	Headers   MapHeader     `bson:"headers,omitempty"    json:"headers"`
	CreatedAt time.Time     `bson:"created_at,omitempty" json:"created_at"`
	UpdatedAt time.Time     `bson:"updated_at,omitempty" json:"updated_at"`
}

func (PyroscopeConfig) CollectionName() string { return "pyroscope_config" }
