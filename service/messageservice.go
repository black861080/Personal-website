package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
	"time"
)

// GetMessage 获取留言信息集
// @Tags 留言模块
// @Summary 获取所有留言信息
// @Success 200 {string} json{"code","message"}
// @Router /message/getMessageList [get]
func GetMessage(c *gin.Context) {
	data := make([]*models.Message, 10)
	data = controller.GetMessageList()
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "MessageList获取成功",
		"data":    data,
	})
}

// AddMessage 添加新的留言
// @Tags 留言模块
// @Summary 添加新的留言
// @Accept multipart/form-data
// @Param name formData string false "名字"
// @Param text formData string false "留言内容"
// @Success 200 {string} json{"code","message"}
// @Router /message/addMessage [post]
func AddMessage(c *gin.Context) {
	message := &models.Message{}
	message.Upload_time = time.Now()
	message.Name = c.PostForm("name")
	message.Text = c.PostForm("text")
	controller.CreateMessage(message)
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "添加留言成功！！！",
	})
}
