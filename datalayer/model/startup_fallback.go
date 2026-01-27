package model

import "time"

type StartupFallback struct {
	ID        int64         `json:"-"          gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Logger    StartupLogger `json:"logger"     gorm:"column:logger;type:json;serializer:json;comment:日志输出配置"`
	CreatedAt time.Time     `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time     `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (StartupFallback) TableName() string {
	return "startup_fallback"
}
