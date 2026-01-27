package model

import "time"

// Store 一些 KV 类型的数据存储
type Store struct {
	ID        string    `json:"id"         gorm:"column:id;primaryKey;size:100;notnull;comment:ID"`
	Value     []byte    `json:"value"      gorm:"column:value"`                                     // 数据的值
	Escape    bool      `json:"escape"     gorm:"column:escape;notnull;default:false;comment:是否转义"` // 是否对模板开启转义
	Desc      string    `json:"desc"       gorm:"column:desc;type:text;comment:说明"`                 // 说明
	Version   int64     `json:"version"    gorm:"column:version;notnull;default:0;comment:乐观锁"`     // 乐观锁
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (Store) TableName() string {
	return "store"
}
