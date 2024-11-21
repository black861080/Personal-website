package models

import "gorm.io/gorm"

type Visitor struct {
	gorm.Model
	Username string
	Password string
	Times    uint
	Role     string
	Token    string
}

func (table *Visitor) TableName() string {
	return "visitor"
}
