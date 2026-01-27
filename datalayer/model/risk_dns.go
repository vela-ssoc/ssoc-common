package model

import "time"

// RiskDNS 含有风险的 DNS 解析记录
type RiskDNS struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Domain    string    `json:"domain"     gorm:"column:domain;size:500"` // 域名
	Kind      string    `json:"kind"       gorm:"column:kind;size:50"`    // 类型
	Origin    string    `json:"origin"     gorm:"column:origin;size:100"` // 风险来源
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`       // 有效期
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implement gorm schema.Tabler
func (RiskDNS) TableName() string {
	return "risk_dns"
}

type RiskDNSs []*RiskDNS

func (rds RiskDNSs) DomainKinds() map[string][]string {
	size := len(rds)
	ret := make(map[string][]string, size)
	for _, rd := range rds {
		domain := rd.Domain
		if kinds, exist := ret[domain]; exist {
			ret[domain] = append(kinds, rd.Kind)
		} else {
			ret[domain] = []string{rd.Kind}
		}
	}

	return ret
}
