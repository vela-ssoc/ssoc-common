package model

// Elastic es 代理配置
type Elastic struct {
	ID       int64    `json:"id,string" gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Host     string   `json:"host"      gorm:"column:host;size:255;comment:地址"`
	Username string   `json:"username"  gorm:"column:username;size:50;comment:账户"`
	Password string   `json:"password"  gorm:"column:password;size:50;comment:密码"`
	Hosts    []string `json:"hosts"     gorm:"column:hosts;json;comment:服务器地址"`
	Desc     string   `json:"desc"      gorm:"column:desc;size:255;comment:简介"`
	Enable   bool     `json:"enable"    gorm:"column:enable;notnull;default:false;comment:是否启用"`
}

// TableName implement gorm schema.Tabler
func (Elastic) TableName() string {
	return "elastic"
}
