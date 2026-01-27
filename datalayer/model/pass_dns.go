package model

import "time"

// PassDNS DNS 白名单
type PassDNS struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Domain    string    `json:"domain"     gorm:"column:domain;size:500;index"` // 域名
	Kind      string    `json:"kind"       gorm:"column:kind;size:50"`          // 类型
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`             // 有效期
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (PassDNS) TableName() string {
	return "pass_dns"
}

type PassDNSs []*PassDNS

func (pds PassDNSs) DomainKinds() map[string][]string {
	size := len(pds)
	ret := make(map[string][]string, size)
	for _, pd := range pds {
		domain := pd.Domain
		if kinds, exist := ret[domain]; exist {
			ret[domain] = append(kinds, pd.Kind)
		} else {
			ret[domain] = []string{pd.Kind}
		}
	}

	return ret
}
