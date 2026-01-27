package model

import "time"

type SIEMServer struct {
	ID        int64     `json:"-"          gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name      string    `json:"name"       gorm:"column:name;type:varchar(20);comment:名字"`
	URL       string    `json:"url"        gorm:"column:url;type:varchar(255);comment:服务器地址"`
	Token     string    `json:"token"      gorm:"column:token;type:varchar(255);comment:认证令牌"`
	CreatedAt time.Time `json:"created_at" gorm:"column:updated_at;not null;autoCreateTime(3);comment:更新时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:created_at;not null;autoUpdateTime(3);comment:创建时间"`
}

func (SIEMServer) TableName() string {
	return "siem_server"
}
