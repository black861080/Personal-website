package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func FirstInformation() *models.Information {
	data := make([]*models.Information, 10)
	util.DB.Find(&data)
	return data[0]
}

func CreateInformation(data *models.Information) *gorm.DB {
	return util.DB.Create(data)
}
