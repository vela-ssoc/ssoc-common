package model

import "time"

// SubstanceTask mapped from table <substance_task>
type SubstanceTask struct {
	ID         int64     `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement:true;comment:ID"`
	TaskID     int64     `json:"task_id,string"   gorm:"column:task_id;notnull;index;uniqueIndex:uk_broker_id_task_id_minion_id"`
	MinionID   int64     `json:"minion_id,string" gorm:"column:minion_id;notnull;uniqueIndex:uk_broker_id_task_id_minion_id;comment:节点 ID"`
	Inet       string    `json:"inet"             gorm:"column:inet;notnull;size:20;comment:节点 IP"`
	BrokerID   int64     `json:"broker_id,string" gorm:"column:broker_id;notnull;index;uniqueIndex:uk_broker_id_task_id_minion_id;comment:节点所在的 broker_id"`
	BrokerName string    `json:"broker_name"      gorm:"column:broker_name;notnull;size:100;comment:节点所在的 broker 名字"`
	Failed     bool      `json:"failed"           gorm:"column:failed;notnull;default:false;comment:是否下发失败"`
	Reason     string    `json:"reason"           gorm:"column:reason;size:500;comment:失败原因"`
	Executed   bool      `json:"executed"         gorm:"column:executed;notnull;default:false;comment:是否下发完毕"`
	CreatedAt  time.Time `json:"created_at"       gorm:"column:created_at;notnull;autoCreateTime(3);comment:任务创建时间"`
	UpdatedAt  time.Time `json:"updated_at"       gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:任务更新时间"`
}

// TableName SubstanceTask's table name
func (SubstanceTask) TableName() string {
	return "substance_task"
}
