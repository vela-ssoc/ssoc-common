package model

import "time"

type VictoriaMetrics struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Enabled   bool      `json:"enabled"    gorm:"column:enabled;comment:是否启用"`
	URL       string    `json:"url"        gorm:"column:url;size:255;comment:推送地址"`
	Method    string    `json:"method"     gorm:"column:method;size:10;comment:请求方法"`
	Header    MapHeader `json:"header"     gorm:"column:header;type:json;comment:Header"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (VictoriaMetrics) TableName() string {
	return "victoria_metrics"
}
