package model

import "time"

type KVAudit struct {
	ID        int64     `json:"-"          gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID  int64     `json:"minion_id"  gorm:"column:minion_id;uniqueIndex:uk_minion_id_bucket_key;comment:节点ID"`
	Inet      string    `json:"inet"       gorm:"column:inet;size:20"`
	Bucket    string    `json:"bucket"     gorm:"column:bucket;size:255;uniqueIndex:uk_minion_id_bucket_key"`
	Key       string    `json:"key"        gorm:"column:key;size:255;uniqueIndex:uk_minion_id_bucket_key"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (KVAudit) TableName() string {
	return "kv_audit"
}
