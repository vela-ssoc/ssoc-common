package model

type FieldValue struct {
	Field string `json:"field"`
	Value string `json:"value"`
}

type FieldValues []*FieldValue
