package models

import (
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Blog struct {
	gorm.Model
	Name  string
	Label string
	json.Decoder
	Describe    string `gorm:"type:text"`
	Describe1   string `gorm:"type:text"`
	Image1      []byte `gorm:"type:longblob"`
	Describe2   string `gorm:"type:text"`
	Image2      []byte `gorm:"type:longblob"`
	Describe3   string `gorm:"type:text"`
	Image3      []byte `gorm:"type:longblob"`
	Describe4   string `gorm:"type:text"`
	Image4      []byte `gorm:"type:longblob"`
	Describe5   string `gorm:"type:text"`
	Image5      []byte `gorm:"type:longblob"`
	Describe6   string `gorm:"type:text"`
	Image6      []byte `gorm:"type:longblob"`
	Describe7   string `gorm:"type:text"`
	Image7      []byte `gorm:"type:longblob"`
	Describe8   string `gorm:"type:text"`
	Image8      []byte `gorm:"type:longblob"`
	Describe9   string `gorm:"type:text"`
	Image9      []byte `gorm:"type:longblob"`
	Describe10  string `gorm:"type:text"`
	Upload_time time.Time
}

func (table *Blog) TableName() string {
	return "blog"
}
