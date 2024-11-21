package models

import "gorm.io/gorm"

type Skill struct {
	gorm.Model
	Name        string
	Description string
	Time        string
	Image       []byte `gorm:"type:longblob"`
	Url         string
}

func (Table *Skill) TableName() string {
	return "skill"
}
