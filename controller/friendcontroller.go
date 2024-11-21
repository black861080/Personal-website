package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
	"time"
)

func FindFriend(friend *models.Friend) *models.Friend {
	var result *models.Friend
	util.DB.Where("friend_name = ?", friend.Friend_name).First(&result)
	if result.ID != 0 {
		Recover(result)
	}
	return result
}

func Recover(friend *models.Friend) *gorm.DB {
	return util.DB.Where(friend).Updates(models.Friend{
		Times:     friend.Times + 1,
		LoginTime: time.Now(),
	})
}
