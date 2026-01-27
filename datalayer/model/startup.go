package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

// Startup minion 节点启动脚本参数配置
type Startup struct {
	ID        int64          `json:"id,string"  gorm:"column:id;primaryKey"`                                   // 取自 Minion.ID
	Logger    *StartupLogger `json:"logger"     gorm:"column:logger;type:json;serializer:json;comment:日志输出配置"` // 日志配置项
	Failed    bool           `json:"failed"     gorm:"column:failed;notnull;default:false"`                    // 是否失败
	Reason    string         `json:"reason"     gorm:"column:reason;type:text;comment:下发失败原因"`                 // 失败原因
	CreatedAt time.Time      `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time      `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (Startup) TableName() string {
	return "startup"
}

type StartupLogger struct {
	Level    string `json:"level"    validate:"oneof=debug info warn error"`
	Filename string `json:"filename"`
	Console  bool   `json:"console"`
	Format   string `json:"format"   validate:"oneof=text json"`
	Caller   bool   `json:"caller"`
	Skip     int    `json:"skip"     validate:"gt=-100,lt=100"`
}

func (sl *StartupLogger) Scan(v any) error {
	return json.Unmarshal(v.([]byte), sl)
}

func (sl *StartupLogger) Value() (driver.Value, error) {
	return json.Marshal(sl)
}
