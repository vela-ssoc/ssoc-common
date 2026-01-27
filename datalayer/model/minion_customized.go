package model

import "time"

type MinionCustomized struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;size:10;notnull;unique"`
	Icon      string    `json:"icon"       gorm:"column:icon;type:text"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3)(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3)(3);comment:更新时间"`
}

func (MinionCustomized) TableName() string {
	return "minion_customized"
}
