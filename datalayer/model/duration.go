package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Duration like https://tc39.es/proposal-temporal/docs/duration.html#constructor
type Duration struct {
	Days    int `json:"days"    validate:"gte=0"` // A number of days.
	Hours   int `json:"hours"   validate:"gte=0"` // A number of hours.
	Minutes int `json:"minutes" validate:"gte=0"` // A number of minutes.
	Seconds int `json:"seconds" validate:"gte=0"` // A number of seconds.
}

func (d Duration) Duration() time.Duration {
	return time.Duration(d.Days)*(24*time.Hour) +
		time.Duration(d.Hours)*time.Hour +
		time.Duration(d.Minutes)*time.Minute +
		time.Duration(d.Seconds)*time.Second
}

func (d Duration) String() string {
	return d.Duration().String()
}

func (d *Duration) Scan(src any) error {
	bs, _ := src.([]byte)
	return json.Unmarshal(bs, d)
}

func (d Duration) Value() (driver.Value, error) {
	return json.Marshal(d)
}
