package model

import "time"

type MinionGroup struct {
	ID          int64     `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID    int64     `json:"minion_id,string" gorm:"column:minion_id;index"`
	Inet        string    `json:"inet"             gorm:"column:inet;size:20"`
	Name        string    `json:"name"             gorm:"column:name;size:255"`
	GID         string    `json:"gid"              gorm:"column:gid;size:500"`
	Description string    `json:"description"      gorm:"column:description;type:text"`
	CreatedAt   time.Time `json:"created_at"       gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt   time.Time `json:"updated_at"       gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (MinionGroup) TableName() string {
	return "minion_group"
}
