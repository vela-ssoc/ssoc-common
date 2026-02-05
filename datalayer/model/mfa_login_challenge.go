package model

import "time"

// MFALoginChallenge 用户名密码验证成功后，多因子认证前的临时 token。
// 每个 token 一旦验证通过后即刻失效，为了避免用户输入错误，每个临时
// token 可以被错误验证 MaxAttempt 次。
type MFALoginChallenge struct {
	ID           int64     `json:"id,string"      gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	UserID       int64     `json:"user_id,string" gorm:"column:user_id;uniqueIndex:uk_user_id_token;comment:用户ID"`
	Token        string    `json:"token"          gorm:"column:token;size:255;uniqueIndex:uk_user_id_token;comment:临时令牌"`
	FailCount    int       `json:"fail_count"     gorm:"column:fail_count;comment:已经错误次数"`
	MaxFailCount int       `json:"max_fail_count" gorm:"column:max_fail_count;comment:最大错误次数"`
	ExpiredAt    time.Time `json:"expired_at"     gorm:"column:expired_at;notnull;comment:过期时间"`
	CreatedAt    time.Time `json:"created_at"     gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt    time.Time `json:"updated_at"     gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (MFALoginChallenge) TableName() string {
	return "mfa_login_challenge"
}
