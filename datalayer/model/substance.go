package model

import (
	"slices"
	"time"
)

// Substance 配置表
type Substance struct {
	ID           int64           `json:"id,string"               gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name         string          `json:"name"                    gorm:"column:name;size:100"` // 名字
	Icon         []byte          `json:"icon"                    gorm:"column:icon"`          // 图标
	Hash         string          `json:"hash"                    gorm:"column:hash"`          // 校验码
	Desc         string          `json:"desc"                    gorm:"column:desc"`          // 描述
	Chunk        []byte          `json:"chunk"                   gorm:"column:chunk"`         // 配置内容
	Links        []string        `json:"links"                   gorm:"column:links;json"`    // 外链，现已弱化，后期可能会删除
	MinionID     int64           `json:"minion_id,string"        gorm:"column:minion_id"`     // 私有配置发布的节点，NULL 或 空 代表是公有配置
	Version      int64           `json:"version"                 gorm:"column:version"`       // 乐观锁
	Priority     int64           `json:"priority"                gorm:"column:priority;comment:越大优先级越高"`
	CreatedID    int64           `json:"created_id,string"       gorm:"column:created_id"` // 创建者 ID
	UpdatedID    int64           `json:"updated_id,string"       gorm:"column:updated_id"` // 最后一个修改者 ID
	ContentQuote *ExtensionQuote `json:"content_quote,omitempty" gorm:"column:content_quote;type:json;serializer:json;comment:插件引用"`
	CreatedAt    time.Time       `json:"created_at"              gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt    time.Time       `json:"updated_at"              gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (Substance) TableName() string {
	return "substance"
}

type Substances []*Substance

// Sort 按照优先级排序，优先级越大越靠前。
// 如果优先级一样则按照 ID 升序。
func (ss Substances) Sort() {
	slices.SortFunc(ss, func(a, b *Substance) int {
		if a.Priority != b.Priority {
			return int(b.Priority - a.Priority)
		}
		return int(a.ID - b.ID)
	})
}
