package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func GetSkillList() []*models.Skill {
	data := make([]*models.Skill, 10)
	util.DB.Find(&data)
	return data
}

func CreateSkill(skill *models.Skill) *gorm.DB {
	return util.DB.Create(skill)
}
