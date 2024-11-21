package models

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	gorm.Model
	Name        string
	Upload_time time.Time
	Text        string `gorm:"type:text"`
}

func (table *Message) TableName() string {
	return "message"
}
