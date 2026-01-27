package model

import "time"

type AuthTemp struct {
	ID        int64     `json:"id"         gorm:"column:id;primaryKey;comment:用户ID"`
	UID       string    `json:"uid"        gorm:"column:uid;size:255;notnull;unique;comment:临时唯一凭证"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
}

func (AuthTemp) TableName() string {
	return "auth_temp"
}

func (t AuthTemp) Expired(current time.Time, du time.Duration) bool {
	return t.CreatedAt.Add(du).Before(current)
}
