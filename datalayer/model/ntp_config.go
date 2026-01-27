package model

import "time"

type NTPConfig struct {
	ID       int64     `json:"-"                  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Server   string    `json:"server,omitempty"   gorm:"column:server;size:255;comment:NTP服务器"`
	LastedAt time.Time `json:"lasted_at,omitzero" gorm:"column:lasted_at;comment:设置时间"`
}

func (NTPConfig) TableName() string {
	return "ntp_config"
}
