package model

type Purl struct {
	ID string `json:"id" gorm:"column:id;primaryKey;size:500;notnull;comment:PURL"`
}

func (p Purl) TableName() string {
	return "purl"
}
