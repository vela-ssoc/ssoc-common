package model

import "time"

type MinionLogon struct {
	ID        int64     `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID  int64     `json:"minion_id,string" gorm:"column:minion_id;index"`
	Inet      string    `json:"inet"             gorm:"column:inet;size:20"`
	User      string    `json:"user"             gorm:"column:user;size:255"`
	Addr      string    `json:"addr"             gorm:"column:addr;size:100"`
	Msg       string    `json:"msg"              gorm:"column:msg;size:10;index:idx_msg"`
	Type      string    `json:"type"             gorm:"column:type;size:50"`
	PID       int       `json:"pid"              gorm:"column:pid"`
	Device    string    `json:"device"           gorm:"column:device;size:255"`
	Process   string    `json:"process"          gorm:"column:process;size:255"`
	LogonAt   time.Time `json:"logon_at"         gorm:"column:logon_at;notnull;autoCreateTime(3);index:idx_logon_at"`
	Ignore    bool      `json:"ignore"           gorm:"column:ignore;notnull;default:false"`
	CreatedAt time.Time `json:"created_at"       gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
}

func (MinionLogon) TableName() string {
	return "minion_logon"
}
