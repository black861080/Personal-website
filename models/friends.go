package models

import (
	"gorm.io/gorm"
	"time"
)

type Friend struct {
	gorm.Model
	Friend_name  string
	Route        string
	LoginTime    time.Time `gorm:"column:login_time" json:"login_time"`
	LoginOutTime time.Time `gorm:"column:login_out_time" json:"login_out_time"`
	LoginIP      string
	Times        uint
	Text         string `gorm:"type:text"`
	Upload_time  time.Time
}

func (table *Friend) TableName() string {
	return "friends"
}
