package model

import "time"

// MinionBin minion 节点二进制发行版本记录表
type MinionBin struct {
	ID         int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	FileID     int64     `json:"-"          gorm:"column:file_id;notnull;comment:文件名"`
	Goos       string    `json:"goos"       gorm:"column:goos;size:10;notnull;comment:操作系统"`
	Arch       string    `json:"arch"       gorm:"column:arch;size:10;notnull;comment:系统架构"`
	Name       string    `json:"name"       gorm:"column:name;size:255;notnull;comment:文件名"`
	Customized string    `json:"customized" gorm:"column:customized;size:50;notnull;default:'';comment:定制版本"`
	Unstable   bool      `json:"unstable"   gorm:"column:unstable;notnull;default:false;comment:内测版本"`
	Caution    string    `json:"caution"    gorm:"column:caution;type:text;comment:注意事项"`
	Ability    string    `json:"ability"    gorm:"column:ability;type:text;comment:功能说明"`
	Size       int64     `json:"size"       gorm:"column:size;comment:文件大小"`
	Hash       string    `json:"hash"       gorm:"column:hash;size:50;comment:文件哈希"`
	Changelog  string    `json:"changelog"  gorm:"column:changelog;type:text;comment:变更日志"`
	Semver     Semver    `json:"semver"     gorm:"column:semver;size:50;comment:版本号"`
	Weight     uint64    `json:"-"          gorm:"column:weight;comment:数值版本号"`
	Deprecated bool      `json:"deprecated" gorm:"column:deprecated;notnull;default:false;comment:是否弃用"`
	CreatedAt  time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt  time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (MinionBin) TableName() string {
	return "minion_bin"
}
