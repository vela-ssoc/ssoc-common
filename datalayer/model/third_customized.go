package model

import "time"

type ThirdCustomized struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;size:100;uniqueIndex;comment:标签名"`
	Icon      string    `json:"icon"       gorm:"column:icon;type:text;comment:图标"`
	Remark    string    `json:"remark"     gorm:"column:remark;size:255;comment:备注"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (ThirdCustomized) TableName() string {
	return "third_customized"
}
