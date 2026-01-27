package model

import "time"

type PassIP struct {
	ID        int64     `json:"id,string"  gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	IP        string    `json:"ip"         gorm:"column:ip;size:20"`
	Kind      string    `json:"kind"       gorm:"column:kind;size:50"`
	BeforeAt  time.Time `json:"before_at"  gorm:"column:before_at"`
	CreatedAt time.Time `json:"created_at" gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt time.Time `json:"updated_at" gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

func (PassIP) TableName() string {
	return "pass_ip"
}

type PassIPs []*PassIP

func (pis PassIPs) IPKinds() map[string][]string {
	size := len(pis)
	ret := make(map[string][]string, size)
	for _, ip := range pis {
		if kinds, exist := ret[ip.IP]; exist {
			ret[ip.IP] = append(kinds, ip.Kind)
		} else {
			ret[ip.IP] = []string{ip.Kind}
		}
	}

	return ret
}
