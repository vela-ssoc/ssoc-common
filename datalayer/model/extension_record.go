package model

import "time"

type ExtensionRecord struct {
	ID          int64     `json:"-"                   gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	ExtensionID int64     `json:"extension_id,string" gorm:"column:extension_id;notnull;uniqueIndex:idx_extension_id_version;comment:插件ID"`
	Version     int64     `json:"version"             gorm:"column:version;uniqueIndex:idx_extension_id_version;comment:版本号"`
	Content     string    `json:"content"             gorm:"column:content;type:text;notnull;comment:代码"`
	ContentSHA1 string    `json:"content_sha1"        gorm:"column:content_sha1;type:char(40);notnull;comment:SHA1"`
	CreatedBy   Operator  `json:"created_by"          gorm:"column:created_by;type:json;notnull;serializer:json;comment:创建者"`
	CreatedAt   time.Time `json:"created_at"          gorm:"column:updated_at;notnull;autoCreateTime(3);comment:更新时间"`
}

func (ExtensionRecord) TableName() string {
	return "extension_record"
}
