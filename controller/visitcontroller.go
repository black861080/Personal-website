package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func FindVisitor(visitor *models.Visitor) *models.Visitor {
	var result *models.Visitor
	util.DB.Where("username=? and password=?", visitor.Username, visitor.Password).Find(&result)
	if result.ID != 0 {
		AddOne(result)
	}
	result.Times = result.Times + 1
	return result
}

func AddOne(visitor *models.Visitor) *gorm.DB {
	return util.DB.Where(visitor).Updates(models.Visitor{
		Times: visitor.Times + 1,
	})
}

func FindRole(visitor *models.Visitor) *models.Visitor {
	var result models.Visitor
	util.DB.Where("token = ?", visitor.Token).First(&result)
	return &result
}

func CreateVisitor(visitor *models.Visitor) *gorm.DB {
	return util.DB.Create(visitor)
}

func CheckUsername(visitor *models.Visitor) *models.Visitor {
	var result *models.Visitor
	util.DB.Where("username=? ", visitor.Username).Find(&result)
	if result.ID != 0 {
		AddOne(result)
	}
	result.Times = result.Times + 1
	return result
}
