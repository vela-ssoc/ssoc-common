package model

import "time"

// Certificate 证书模块
type Certificate struct {
	ID             int64     `json:"id,string"       gorm:"column:id;primaryKey;comment:用户ID"`
	Name           string    `json:"name"            gorm:"column:name;size:255;comment:证书名"`
	Certificate    string    `json:"certificate"     gorm:"column:certificate;type:text;notnull;comment:证书"`
	PrivateKey     string    `json:"private_key"     gorm:"column:private_key;type:text;notnull;comment:私钥"`
	Version        int       `json:"version"         gorm:"column:version;comment:证书版本"`
	IssCountry     []string  `json:"iss_country"     gorm:"column:iss_country;json;comment:颁发者国家"`
	IssProvince    []string  `json:"iss_province"    gorm:"column:iss_province;json;comment:颁发者省份"`
	IssOrg         []string  `json:"iss_org"         gorm:"column:iss_org;json;comment:颁发者组织"`
	IssCN          string    `json:"iss_cn"          gorm:"column:iss_cn"`               // 颁发者 Common Name
	IssOrgUnit     []string  `json:"iss_org_unit"    gorm:"column:iss_org_unit;json"`    // 颁发者组织单位
	SubCountry     []string  `json:"sub_country"     gorm:"column:sub_country;json"`     // 主题国家
	SubOrg         []string  `json:"sub_org"         gorm:"column:sub_org;json"`         // 主题组织
	SubProvince    []string  `json:"sub_province"    gorm:"column:sub_province;json"`    // 主题省份
	SubCN          string    `json:"sub_cn"          gorm:"column:sub_cn"`               // 主题 Common Name
	DNSNames       []string  `json:"dns_names"       gorm:"column:dns_names;json"`       // 证书 DNS Name
	IPAddresses    []string  `json:"ip_addresses"    gorm:"column:ip_addresses;json"`    // IP Addresses
	EmailAddresses []string  `json:"email_addresses" gorm:"column:email_addresses;json"` // Email Addresses
	URIs           []string  `json:"uris"            gorm:"column:uris;json"`            // URIs
	NotBefore      time.Time `json:"not_before"      gorm:"column:not_before"`           // 证书生效时间
	NotAfter       time.Time `json:"not_after"       gorm:"column:not_after"`            // 证书过期时间
	CreatedAt      time.Time `json:"created_at"      gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	UpdatedAt      time.Time `json:"updated_at"      gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

// TableName implemented schema.Tabler
func (Certificate) TableName() string {
	return "certificate"
}
