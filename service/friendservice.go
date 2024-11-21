package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
)

// FindFriend 寻找朋友信息
// @Tags 朋友模块
// @Summary 查询朋友信息
// @Accept multipart/form-data
// @Param name formData string false "您的名字"
// @Success 200 {string} json{"code","message"}
// @Router /friend/findFriend [post]
func FindFriend(c *gin.Context) {
	friend := &models.Friend{}
	friend.Friend_name = c.PostForm("name")
	good_friend := controller.FindFriend(friend)
	if good_friend.ID != 0 {
		c.JSON(200, gin.H{
			"code":    0, //成功
			"message": "查询成功",
			"data":    good_friend,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    -1, //失败
			"message": "未查询到此人",
		})
	}

}
