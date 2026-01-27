package model

type SyslogConfig struct {
	ID       int64  `json:"-"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	Enabled  bool   `json:"enabled"  gorm:"column:enabled;comment:是否启用"`
	AppName  string `json:"app_name" gorm:"column:app_name;size:50;comment:应用名"`
	Protocol string `json:"protocol" gorm:"column:protocol;size:10;comment:协议"`
	Address  string `json:"address"  gorm:"column:address;size:255;comment:服务器地址"`
}

func (SyslogConfig) TableName() string {
	return "syslog_config"
}
