package model

import "time"

// BrokerBin 存放 broker 节点的可执行程序。
type BrokerBin struct {
	ID           int64     `json:"id,string"            gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Name         string    `json:"name"                 gorm:"column:name;size:255;notnull;comment:名字"`
	FileID       int64     `json:"-"                    gorm:"column:file_id;notnull;comment:文件ID"`
	Size         int64     `json:"size"                 gorm:"column:size;comment:文件大小"`
	Hash         string    `json:"hash"                 gorm:"column:hash;size:50;comment:文件哈希"`
	Goos         string    `json:"goos"                 gorm:"column:goos;size:10;notnull;comment:操作系统"`
	Arch         string    `json:"arch"                 gorm:"column:arch;size:15;notnull;comment:系统架构"`
	Semver       Semver    `json:"semver"               gorm:"column:semver;size:50;notnull;comment:版本号"`
	SemverWeight uint64    `json:"semver_weight,string" gorm:"column:semver_weight;notnull;default:0;comment:版本号数值"`
	Changelog    string    `json:"changelog"            gorm:"column:changelog;type:text;comment:更新日志"`
	CreatedAt    time.Time `json:"created_at"           gorm:"column:created_at;autoCreateTime(3);comment:创建时间"`
	UpdatedAt    time.Time `json:"updated_at"           gorm:"column:updated_at;autoUpdateTime(3);comment:更新时间"`
}

func (BrokerBin) TableName() string {
	return "broker_bin"
}
