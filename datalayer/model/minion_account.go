package model

import "time"

type MinionAccount struct {
	ID          int64     `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID    int64     `json:"minion_id,string" gorm:"column:minion_id;index;comment:节点ID"`
	Inet        string    `json:"inet"             gorm:"column:inet;size:20;comment:节点IP"`
	Name        string    `json:"name"             gorm:"column:name;size:255;comment:用户名"`
	LoginName   string    `json:"login_name"       gorm:"column:login_name;size:255;comment:登录名"`
	UID         string    `json:"uid"              gorm:"column:uid;size:255;comment:UID"`
	GID         string    `json:"gid"              gorm:"column:gid;size:255;comment:GID"`
	HomeDir     string    `json:"home_dir"         gorm:"column:home_dir;size:500;comment:家目录"`
	Description string    `json:"description"      gorm:"column:description;type:text;comment:用户描述"`
	Status      string    `json:"status"           gorm:"column:status;size:100;comment:状态"`
	Raw         string    `json:"raw"              gorm:"column:raw;type:text"`
	CreatedAt   time.Time `json:"created_at"       gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt   time.Time `json:"updated_at"       gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (MinionAccount) TableName() string {
	return "minion_account"
}
