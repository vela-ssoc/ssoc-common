package model

import (
	"database/sql/driver"
	"encoding/json"
)

type Operator struct {
	ID       int64  `json:"id,string" gorm:"column:id"`
	Username string `json:"username"  gorm:"column:username"`
	Nickname string `json:"nickname"  gorm:"column:nickname"`
}

func (o Operator) Value() (driver.Value, error) {
	return json.Marshal(o)
}
