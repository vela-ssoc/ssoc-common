package model

// Email 邮件发送配置
type Email struct {
	ID       int64  `json:"id,string" gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Host     string `json:"host"      gorm:"column:host;size:255;comment:邮箱服务器"`
	Username string `json:"username"  gorm:"column:username;size:255;comment:账号"`
	Password string `json:"password"  gorm:"column:password;size:255;comment:密码"`
	Enable   bool   `json:"enable"    gorm:"column:enable;notnull;default:false;comment:是否启用"`
}

// TableName implement gorm schema.Tabler
func (Email) TableName() string {
	return "email"
}
