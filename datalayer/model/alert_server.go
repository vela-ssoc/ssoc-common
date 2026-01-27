package model

import "time"

// AlertServer 告警服务器配置
type AlertServer struct {
	ID        int64     `json:"-"          gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Mode      string    `json:"mode"       gorm:"column:mode;size:10;notnull;comment:发送模式"` // 枚举：siem dong
	Name      string    `json:"name"       gorm:"column:name;size:20;notnull;uniqueIndex;comment:名字"`
	URL       string    `json:"url"        gorm:"column:url;size:255;notnull;comment:服务器地址"`
	Token     string    `json:"token"      gorm:"column:token;size:255;notnull;comment:认证令牌"`
	Account   string    `json:"account"    gorm:"column:account;size:20;comment:咚咚账号"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (AlertServer) TableName() string {
	return "alert_server"
}
