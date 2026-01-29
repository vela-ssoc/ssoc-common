package model

import "time"

type Pyroscope struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Enabled   bool      `json:"enabled"    gorm:"column:enabled;comment:是否启用"`
	URL       string    `json:"url"        gorm:"column:url;size:255;comment:推送地址"`
	Username  string    `json:"username"   gorm:"column:username;size:50;comment:认证用户"`
	Password  string    `json:"password"   gorm:"column:password;size:255;comment:认证密码"`
	Header    MapHeader `json:"header"     gorm:"column:header;type:json;serializer:json;comment:Header"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (Pyroscope) TableName() string {
	return "pyroscope"
}
