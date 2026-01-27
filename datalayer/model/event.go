package model

import "time"

// Event 存放节点事件信息
type Event struct {
	ID        int64          `json:"id,string"        gorm:"column:id;primaryKey;autoIncrement;comment:ID"`
	MinionID  int64          `json:"minion_id,string" gorm:"column:minion_id"`                          // 节点 ID
	Inet      string         `json:"inet"             gorm:"column:inet;size:20"`                       // 节点 IPv4
	Subject   string         `json:"subject"          gorm:"column:subject;size:100"`                   // 主题
	FromCode  string         `json:"from_code"        gorm:"column:from_code;size:50"`                  // 来源模块
	Typeof    string         `json:"typeof"           gorm:"column:typeof;size:50"`                     // 模块类型
	User      string         `json:"user"             gorm:"column:user;size:50"`                       // 用户信息
	Auth      string         `json:"auth"             gorm:"column:auth;size:50"`                       // 认证信息
	Msg       string         `json:"msg"              gorm:"column:msg;type:text"`                      // 上报消息
	Error     string         `json:"error"            gorm:"column:error;type:text"`                    // 错误信息
	Level     EventLevel     `json:"level"            gorm:"column:level"`                              // 告警级别
	HaveRead  bool           `json:"have_read"        gorm:"column:have_read;notnull;default:false"`    // 是否已读确认
	SendAlert bool           `json:"send_alert"       gorm:"column:send_alert;notnull;default:false"`   // 是否需要发送告警
	Secret    string         `json:"-"                gorm:"column:secret;size:100"`                    // 如果告警，生成随机字符串防止恶意遍历
	Metadata  map[string]any `json:"metadata"         gorm:"column:metadata;type:json;serializer:json"` // 附带数据
	OccurAt   time.Time      `json:"occur_at"         gorm:"column:occur_at"`                           // 事件发生的时间
	CreatedAt time.Time      `json:"created_at"       gorm:"column:created_at"`                         // 创建时间
}

// TableName implement gorm schema.Tabler
func (Event) TableName() string {
	return "event"
}

// ShortMsg 短消息
func (evt Event) ShortMsg(size ...int) string {
	sz := 100
	if len(size) > 0 && size[0] > 0 {
		sz = size[0]
	}
	return evt.shortString(evt.Msg, sz)
}

// ShortError 短错误信息
func (evt Event) ShortError(size ...int) string {
	sz := 100
	if len(size) > 0 && size[0] > 0 {
		sz = size[0]
	}
	return evt.shortString(evt.Error, sz)
}

func (Event) shortString(str string, size int) string {
	// len([]byte("你好世界")) == 12
	// len([]rune("你好世界")) == 4
	unicodes := []rune(str)
	if size >= len(unicodes) {
		return str
	}
	return string(unicodes[:size]) + "..."
}
