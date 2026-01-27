package model

import "time"

type RiskFile struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Checksum  string    `json:"checksum"   gorm:"column:checksum;size:100"`
	Algorithm string    `json:"algorithm"  gorm:"column:algorithm;size:10"`
	Kind      string    `json:"kind"       gorm:"column:kind;size:50"`
	Origin    string    `json:"origin"     gorm:"column:origin;size:100"`
	Desc      string    `json:"desc"       gorm:"column:desc;size:500"`
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (RiskFile) TableName() string {
	return "risk_file"
}

type RiskFiles []*RiskFile

func (rfs RiskFiles) ChecksumKinds() map[string][]string {
	size := len(rfs)
	ret := make(map[string][]string, size)
	for _, file := range rfs {
		sum := file.Checksum
		if kinds, exist := ret[sum]; exist {
			ret[sum] = append(kinds, file.Kind)
		} else {
			ret[sum] = []string{file.Kind}
		}
	}

	return ret
}
