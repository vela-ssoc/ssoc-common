package model

import (
	"encoding/json"
	"time"
)

type AgentEmergencySnapshot struct {
	ID         int64           `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID   int64           `json:"minion_id,string" gorm:"column:minion_id;index;uniqueIndex:uk_minion_id_type;comment:节点ID"`
	Inet       string          `json:"inet"             gorm:"column:inet;size:20"`
	Type       string          `json:"type"             gorm:"column:type;size:50;uniqueIndex:uk_minion_id_type;comment:快照类型"`
	Value      json.RawMessage `json:"value"            gorm:"column:value;type:json;serializer:json;comment:快照数据"`
	CreatedAt  time.Time       `json:"created_at"       gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	ReportedAt time.Time       `json:"reported_at"      gorm:"column:reported_at;notnull;autoUpdateTime(3);comment:上报时间"`
	UpdatedAt  time.Time       `json:"updated_at"       gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (AgentEmergencySnapshot) TableName() string {
	return "agent_emergency_snapshot"
}
