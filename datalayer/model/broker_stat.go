package model

import "time"

type BrokerStat struct {
	ID         int64     `json:"id,string"   gorm:"column:id;primaryKey;comment:ID"`
	Name       string    `json:"name"        gorm:"name;size:100;comment:名字"`
	MemUsed    uint64    `json:"mem_used"    gorm:"mem_used;comment:内存使用量"`
	MemTotal   uint64    `json:"mem_total"   gorm:"mem_total;comment:内存总量"`
	CPUPercent float64   `json:"cpu_percent" gorm:"cpu_percent;comment:CPU使用百分比"`
	CreatedAt  time.Time `json:"created_at"  gorm:"column:created_at;autoCreateTime(3);comment:创建时间"`
	UpdatedAt  time.Time `json:"updated_at"  gorm:"column:updated_at;autoUpdateTime(3);comment:更新时间"`
}

func (BrokerStat) TableName() string {
	return "broker_stat"
}
