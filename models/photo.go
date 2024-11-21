package models

import (
	"gorm.io/gorm"
	"time"
)

type Photo struct {
	gorm.Model
	PhotoName   string
	Image       []byte `gorm:"type:longblob"`
	Upload_time time.Time
	Description string
}

func (table *Photo) TableName() string {
	return "photo"
}
