package model

import "time"

type MinionListen struct {
	ID        int64     `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID  int64     `json:"minion_id,string" gorm:"column:minion_id"`
	Inet      string    `json:"inet"             gorm:"column:inet;size:20"`
	RecordID  string    `json:"record_id"        gorm:"column:record_id;size:255"`
	PID       uint32    `json:"pid"              gorm:"column:pid"`
	FD        int       `json:"fd"               gorm:"column:fd"`
	Family    uint8     `json:"family"           gorm:"column:family"`
	Protocol  uint8     `json:"protocol"         gorm:"column:protocol"`
	LocalIP   string    `json:"local_ip"         gorm:"column:local_ip;size:50"`
	LocalPort int       `json:"local_port"       gorm:"column:local_port"`
	Path      string    `json:"path"             gorm:"column:path;size:500"`
	Process   string    `json:"process"          gorm:"column:process;size:500"`
	Username  string    `json:"username"         gorm:"column:username;size:255"`
	CreatedAt time.Time `json:"created_at"       gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at"       gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (MinionListen) TableName() string {
	return "minion_listen"
}
