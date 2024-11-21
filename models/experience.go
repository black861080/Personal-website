package models

import "gorm.io/gorm"

type Experience struct {
	gorm.Model
	Name        string
	Description string
	Time        string
	Image       []byte `gorm:"type:longblob"`
	Url         string
}

func (table *Experience) TableName() string {
	return "experience"
}
