package model

import (
	"time"
)

type Certificate2 struct {
	ID                int64               `json:"id,string"             gorm:"column:id;primaryKey;comment:ID"`
	Enabled           bool                `json:"enabled"               gorm:"column:enabled;comment:是否启用"`
	CommonName        string              `json:"common_name"           gorm:"column:common_name;size:50;comment:共用名"`
	PublicKey         []byte              `json:"public_key,omitempty"  gorm:"column:public_key;comment:证书"`
	PrivateKey        []byte              `json:"private_key,omitempty" gorm:"column:private_key;comment:私钥"`
	CertificateSHA256 string              `json:"certificate_sha256"    gorm:"column:certificate_sha256;size:100;uniqueIndex;comment:证书指纹"`
	PublicKeySHA256   string              `json:"public_key_sha256"     gorm:"column:public_key_sha256;size:100;comment:公钥指纹"`
	PrivateKeySHA256  string              `json:"private_key_sha256"    gorm:"column:private_key_sha256;size:100;comment:私钥指纹"`
	DNSNames          []string            `json:"dns_names"             gorm:"column:dns_names;type:json;serializer:json;comment:域名"`
	IPAddresses       []string            `json:"ip_addresses"          gorm:"column:ip_addresses;type:json;serializer:json;comment:IP"`
	EmailAddresses    []string            `json:"email_addresses"       gorm:"column:email_addresses;type:json;serializer:json;comment:邮箱"`
	URIs              []string            `json:"uris"                  gorm:"column:uris;type:json;serializer:json;comment:URI"`
	Version           int                 `json:"version"               gorm:"column:version;comment:版本"`
	NotBefore         time.Time           `json:"not_before"            gorm:"column:not_before;comment:生效时间"`
	NotAfter          time.Time           `json:"not_after"             gorm:"column:not_after;comment:过期时间"`
	Issuer            CertificatePKIXName `json:"issuer"                gorm:"column:issuer;type:json;serializer:json;comment:颁发者"`
	Subject           CertificatePKIXName `json:"subject"               gorm:"column:subject;type:json;serializer:json;comment:主题"`
	UpdatedAt         time.Time           `json:"updated_at"            gorm:"column:created_at;notnull;autoCreateTime(3);comment:创建时间"`
	CreatedAt         time.Time           `json:"created_at"            gorm:"column:updated_at;notnull;autoUpdateTime(3);comment:更新时间"`
}

type CertificatePKIXName struct {
	Country            []string `json:"country"`
	Organization       []string `json:"organization"`
	OrganizationalUnit []string `json:"organizational_unit"`
	Locality           []string `json:"locality"`
	Province           []string `json:"province"`
	StreetAddress      []string `json:"street_address"`
	PostalCode         []string `json:"postal_code"`
	SerialNumber       string   `json:"serial_number"`
	CommonName         string   `json:"common_name"`
}

func (Certificate2) TableName() string {
	return "certificate2"
}
