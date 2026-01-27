package model

import "time"

type GridFile struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;size:255;notnull;index;comment:名字"`
	Size      int64     `json:"size"       gorm:"column:size;comment:文件大小"`
	Checksum  string    `json:"checksum"   gorm:"column:checksum;type:char(32);comment:MD5"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (GridFile) TableName() string {
	return "gridfs_file"
}
