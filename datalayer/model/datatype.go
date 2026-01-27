package model

import (
	"bytes"
	"errors"
	"strconv"
)

const (
	extNumberType         extFieldType = "number"          // 数字类型
	extBoolType           extFieldType = "bool"            // 布尔类型
	extStringType         extFieldType = "string"          // 字符串类型
	extReferType          extFieldType = "ref"             // 引用类型
	extReadonlyStringType extFieldType = "readonly_string" // 只读字符串类型
)

// extFieldType startup 配置项扩展字段类型
type extFieldType string

// startupExtend 扩展字段
type startupExtend struct {
	Name  string `json:"name"  validate:"required"`                                     // 名字
	Type  string `json:"type"  validate:"oneof=number bool string ref string_readonly"` // 类型
	Value string `json:"value" validate:"required"`                                     // 值
}

// parseField 校验单个一项数据
func (s startupExtend) parse() (*startupExtend, error) {
	ret := &startupExtend{Name: s.Name, Type: s.Type, Value: s.Value}
	switch extFieldType(s.Type) {
	case extNumberType:
		if _, err := strconv.ParseFloat(s.Value, 64); err != nil {
			return nil, err
		}
	case extBoolType:
		tf, err := strconv.ParseBool(ret.Value)
		if err != nil {
			return nil, err
		}
		ret.Value = strconv.FormatBool(tf)
	case extStringType:
		ret.Value = s.escape()
	case extReferType:
	case extReadonlyStringType:
		ret.Value = "r" + s.escape()
	default:
		return nil, errors.New("不支持的数据类型: " + s.Type)
	}

	return ret, nil
}

// escape 对字符串中的字符转译：
// 例如：
//
//	a"bc"d --> a\"bc\"d
//	\r\n   --> \\r\\n
func (s startupExtend) escape() string {
	var buf bytes.Buffer
	buf.WriteByte('"')
	for _, char := range s.Value {
		switch char {
		case '"', '\\':
			buf.WriteByte('\\')
		}
		buf.WriteRune(char)
	}
	buf.WriteByte('"')
	return buf.String()
}

// TaskRunner 内部运行状态
type TaskRunner struct {
	Name     string         `json:"name"`   // 内部服务名字
	Type     string         `json:"type"`   // 类型
	Status   string         `json:"status"` // 状态
	CodeVM   string         `json:"code_vm"`
	Private  bool           `json:"private"`
	Cause    string         `json:"cause"`
	Metadata map[string]any `json:"metadata"`
	UseBy    []string       `json:"use_by"`
}
