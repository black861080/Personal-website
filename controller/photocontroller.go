package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func GetPhotoList() []*models.Photo {
	data := make([]*models.Photo, 10)
	util.DB.Find(&data)
	return data
}

func CreatePhoto(photo *models.Photo) *gorm.DB {
	return util.DB.Create(photo)
}

func UpdatePhoto(photo *models.Photo) *gorm.DB {
	return util.DB.Model(photo).Updates(models.Photo{
		PhotoName:   photo.PhotoName,
		Image:       photo.Image,
		Upload_time: photo.Upload_time,
		Description: photo.Description,
	})
}

func DeletePhoto(photo *models.Photo) *gorm.DB {
	return util.DB.Where("id=?", photo.ID).Delete(photo)
}
