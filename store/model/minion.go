package model

import (
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/v2/bson"
)

const (
	MinionStatusOffline MinionStatus = iota // 离线，默认状态。
	MinionStatusOnline                      // 在线。
	MinionStatusDelete                      // 删除：标记后此节点不允许上线。
)

const (
	MinionTagSystem MinionTagType = iota
	MinionTagManual
	MinionTagReport
)

type MinionStatus int

func (s MinionStatus) String() string {
	switch s {
	case MinionStatusOffline:
		return "离线"
	case MinionStatusOnline:
		return "在线"
	case MinionStatusDelete:
		return "删除"
	default:
		return "未知状态：" + strconv.Itoa(int(s))
	}
}

// MinionTagType 节点标签类型。
// 系统标签不可删除。
// 人工标签只能人工新增删除。
// 上报标签由 agent 程序扫描自身环境补充，可以人工删除可以 agent 程序自动删除。
type MinionTagType int

func (m MinionTagType) String() string {
	switch m {
	case MinionTagSystem:
		return "系统标签"
	case MinionTagManual:
		return "人工标签"
	case MinionTagReport:
		return "上报标签"
	default:
		return "未知标签类型：" + strconv.Itoa(int(m))
	}
}

type Minion struct {
	ID          bson.ObjectID `bson:"_id,omitempty"          json:"id"`
	MachineID   string        `bson:"machine_id"             json:"machine_id"`
	Status      MinionStatus  `bson:"status"                 json:"status"`
	Tags        MinionTags    `bson:"tags"                   json:"tags"`   // 节点标签，配置下发。
	Unload      bool          `bson:"unload"                 json:"unload"` // 此模式开启，此节点不会加载任何配置。
	TunnelStat  *TunnelStat   `bson:"tunnel_stat,omitempty"  json:"tunnel_stat,omitzero"`
	ExecuteStat *ExecuteStat  `bson:"execute_stat,omitempty" json:"execute_stat,omitzero"`
	CMDB        *MinionCMDB   `bson:"cmdb,omitempty"         json:"cmdb,omitempty"`
	CreatedAt   time.Time     `bson:"created_at,omitempty"   json:"created_at"`
	UpdatedAt   time.Time     `bson:"updated_at,omitempty"   json:"updated_at"`
}

func (Minion) CollectionName() string { return "minion" }

// MinionCMDB 节点 CMDB 简要信息，辅助搜索与。
type MinionCMDB struct {
	Comment string `bson:"comment" json:"comment"`
}
