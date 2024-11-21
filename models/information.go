package models

import (
	"gorm.io/gorm"
)

type Information struct {
	gorm.Model
	Image                 []byte `gorm:"type:longblob"`
	Name                  string
	NikeName              string
	Location              string
	Street                string
	Description           string
	WebDescription        string
	ProjectDescription    string
	PhotoDescription      string
	BlogDescription       string
	SuggestionDescription string
}

func (table *Information) TableName() string {
	return "information"
}
