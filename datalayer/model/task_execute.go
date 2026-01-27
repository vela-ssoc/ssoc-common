package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

type TaskExecute struct {
	ID            int64             `json:"id,string"                gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	TaskID        int64             `json:"task_id,string"           gorm:"column:task_id;index;comment:任务ID"`
	Name          string            `json:"name"                     gorm:"column:name;type:varchar(100);comment:名字"`
	Intro         string            `json:"intro"                    gorm:"column:intro;type:varchar(1000);comment:简介"`
	Status        TaskExecuteStatus `json:"status"                   gorm:"column:status;type:json;comment:执行状态"`
	Finished      bool              `json:"finished"                 gorm:"column:finished;notnull;default:false;comment:是否结束"`
	Code          string            `json:"code"                     gorm:"column:code;type:text;notnull;comment:执行代码"`
	CodeSHA1      string            `json:"code_sha1"                gorm:"column:code_sha1;type:char(40);notnull;comment:执行代码SHA1"`
	ContentQuote  *ExtensionQuote   `json:"content_quote"            gorm:"column:content_quote;type:json;serializer:json;comment:插件引用"`
	Cron          string            `json:"cron,omitempty"           gorm:"column:cron;size:50;comment:定时任务表达式"`
	SpecificTimes []time.Time       `json:"specific_times,omitempty" gorm:"column:specific_times;type:json;serializer:json;comment:定点任务时间"`
	Timeout       Duration          `json:"timeout"                  gorm:"column:timeout;serializer:json;comment:超时时间"`
	PushSize      int               `json:"push_size"                gorm:"column:push_size;comment:推送并发数"`
	Filters       TaskExecuteFilter `json:"filters"                  gorm:"column:filters;type:json;serializer:json;comment:过滤节点"`
	Excludes      []string          `json:"excludes"                 gorm:"column:excludes;serializer:json;comment:排除节点"`
	CreatedBy     Operator          `json:"created_by"               gorm:"column:created_by;type:json;notnull;serializer:json;comment:任务创建者"`
	UpdatedBy     Operator          `json:"updated_by"               gorm:"column:updated_by;type:json;notnull;serializer:json;comment:任务创建者"`
	CreatedAt     time.Time         `json:"created_at"               gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt     time.Time         `json:"updated_at"               gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (TaskExecute) TableName() string {
	return "task_execute"
}

type TaskExecuteStatus struct {
	Total     int       `json:"total"`      // 节点总数
	Succeed   int       `json:"succeed"`    // 成功节点数
	Failed    int       `json:"failed"`     // 失败节点数
	CreatedAt time.Time `json:"created_at"` // 任务开始时间
	UpdatedAt time.Time `json:"updated_at"` // 该进度更新时间
}

func (t *TaskExecuteStatus) Scan(src any) error {
	bs, _ := src.([]byte)
	return json.Unmarshal(bs, t)
}

func (t TaskExecuteStatus) Value() (driver.Value, error) {
	return json.Marshal(t)
}
