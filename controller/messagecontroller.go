package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func GetMessageList() []*models.Message {
	data := make([]*models.Message, 10)
	util.DB.Find(&data)
	return data
}

func CreateMessage(message *models.Message) *gorm.DB {
	return util.DB.Create(message)
}
