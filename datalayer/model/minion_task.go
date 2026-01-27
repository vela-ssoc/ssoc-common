package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// MinionTask 配置运行
type MinionTask struct {
	ID          int64       `json:"id,string"           gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	SubstanceID int64       `json:"substance_id,string" gorm:"column:substance_id"`
	MinionID    int64       `json:"minion_id,string"    gorm:"column:minion_id"`
	Inet        string      `json:"inet"                gorm:"column:inet;size:20"`
	Dialect     bool        `json:"dialect"             gorm:"column:dialect;notnull;default:false"`
	Name        string      `json:"name"                gorm:"column:name;size:100"`                // 配置名称
	Link        string      `json:"link"                gorm:"column:link;size:255"`                // 外链
	Status      string      `json:"status"              gorm:"column:status;size:10"`               // 运行状态
	Hash        string      `json:"hash"                gorm:"column:hash;size:50"`                 // hash
	From        string      `json:"from"                gorm:"column:from;size:50"`                 // 来源
	Uptime      time.Time   `json:"uptime"              gorm:"column:uptime"`                       // 启动时间
	Failed      bool        `json:"failed"              gorm:"column:failed;notnull;default:false"` // 是否失败
	Cause       string      `json:"cause"               gorm:"column:cause;type:text"`              // 如果发生失败，失败的原因
	Runners     TaskRunners `json:"runners"             gorm:"column:runners"`                      // task 内部服务
	CreatedAt   time.Time   `json:"created_at"          gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
}

// TableName implement gorm schema.Tabler
func (MinionTask) TableName() string {
	return "minion_task"
}

func (mt MinionTask) String() string {
	return mt.Name + "[" + mt.Status + "]"
}

type TaskRunners []*TaskRunner

// Scan implement sql.Scanner
func (trs *TaskRunners) Scan(v any) error {
	return json.Unmarshal(v.([]byte), trs)
}

// Value implement driver.Valuer
func (trs TaskRunners) Value() (driver.Value, error) {
	return json.Marshal(trs)
}
