package model

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserDomain uint8

const (
	UdLocal UserDomain = iota + 1
	UdOA
)

//func (ud UserDomain) String() string {
//	switch ud {
//	case UdLocal:
//		return "本地帐号"
//	case UdOA:
//		return "OA帐号"
//	default:
//		return "未知帐号"
//	}
//}

// User 用户表
type User struct {
	ID         int64          `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Username   string         `json:"username"   gorm:"column:username;size:50;index;comment:用户名"`
	Nickname   string         `json:"nickname"   gorm:"column:nickname;size:50;comment:昵称"`
	Password   string         `json:"-"          gorm:"column:password;size:255;comment:密码"`
	Dong       string         `json:"dong"       gorm:"column:dong;size:10;comment:咚咚账户"`
	Enable     bool           `json:"enable"     gorm:"column:enable;notnull;default:false;comment:是否启用"`
	Domain     UserDomain     `json:"domain"     gorm:"column:domain;comment:用户类型"`
	AccessKey  string         `json:"access_key" gorm:"column:access_key;size:100;index;comment:接口调用密钥"` // AccessKey
	Token      string         `json:"-"          gorm:"column:token;size:255;index;comment:Token"`
	TotpSecret string         `json:"-"          gorm:"column:totp_secret;size:255;comment:TOTP密钥"`
	TotpBind   bool           `json:"-"          gorm:"column:totp_bind;notnull;default:false;comment:TOTP是否已使用"`
	IssueAt    sql.NullTime   `json:"issue_at"   gorm:"column:issue_at;comment:Token签发时间"`
	SessionAt  sql.NullTime   `json:"session_at" gorm:"column:session_at;comment:Session活动时间"`
	DeletedAt  gorm.DeletedAt `json:"-"          gorm:"column:deleted_at;comment:逻辑删除时间"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (User) TableName() string {
	return "user"
}

func (u User) IsLocal() bool {
	return u.Domain == UdLocal
}

func (u User) IsOA() bool {
	return u.Domain == UdOA
}
