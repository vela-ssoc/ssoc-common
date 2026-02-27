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
		return "未知：" + strconv.Itoa(int(s))
	}
}

type Minion struct {
	ID          bson.ObjectID `bson:"_id,omitempty"          json:"id"`
	MachineID   string        `bson:"machine_id"             json:"machine_id"`
	Status      MinionStatus  `bson:"status"                 json:"status"`
	Tags        []string      `bson:"tags"                   json:"tags"` // 节点标签，配置下发。
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
