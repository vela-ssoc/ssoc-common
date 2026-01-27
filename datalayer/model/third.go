package model

import "time"

// Third 3rd 文件表
type Third struct {
	ID         int64     `json:"id,string"         gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	FileID     int64     `json:"file_id,string"    gorm:"column:file_id;comment:文件ID"`
	Name       string    `json:"name"              gorm:"column:name;size:255;comment:文件名"`
	Hash       string    `json:"hash"              gorm:"column:hash;size:100;comment:文件MD5"`
	Desc       string    `json:"desc"              gorm:"column:desc;type:text;comment:文件简介"`
	Size       int64     `json:"size"              gorm:"column:size;comment:文件大小"`
	Customized string    `json:"customized"        gorm:"column:customized;size:100;comment:分类"` // 归类
	Extension  string    `json:"extension"         gorm:"column:extension;size:20;comment:扩展名"`  // 扩展名
	CreatedID  int64     `json:"created_id,string" gorm:"column:created_id;comment:创建者"`
	UpdatedID  int64     `json:"updated_id,string" gorm:"column:updated_id;comment:修改者"`
	CreatedAt  time.Time `json:"created_at"        gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt  time.Time `json:"updated_at"        gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (Third) TableName() string {
	return "third"
}
