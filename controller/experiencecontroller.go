package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func GetExperienceList() []*models.Experience {
	data := make([]*models.Experience, 10)
	util.DB.Find(&data)
	return data
}

func CreateExperience(experience *models.Experience) *gorm.DB {
	return util.DB.Create(experience)
}
