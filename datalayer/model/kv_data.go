package model

import (
	"encoding/json"
	"time"
)

type KVData struct {
	Bucket    string          `json:"bucket"     gorm:"column:bucket;size:255;primaryKey"` // 存储桶
	Key       string          `json:"key"        gorm:"column:key;size:255;primaryKey"`    // key
	Value     json.RawMessage `json:"value"      gorm:"column:value"`                      // value
	Count     int64           `json:"count"      gorm:"column:count"`                      // INCR 计数字段
	Version   int64           `json:"-"          gorm:"column:version"`                    // 乐观锁
	Lifetime  time.Duration   `json:"lifetime"   gorm:"column:lifetime"`                   // 生命时长，大于 0 代表有过期时间。
	ExpiredAt time.Time       `json:"expired_at" gorm:"column:expired_at"`                 // 过期时间
	CreatedAt time.Time       `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time       `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (KVData) TableName() string {
	return "kv_data"
}

func (d KVData) Expired(now time.Time) bool {
	if d.Lifetime <= 0 {
		return false
	}
	return d.ExpiredAt.Before(now)
}
