package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
)

// GetExperience 获取经验集
// @Tags 经验模块
// @Summary 获取所有经验信息
// @Success 200 {string} json{"code","message"}
// @Router /experience/getExperienceList [get]
func GetExperience(c *gin.Context) {
	data := make([]*models.Experience, 10)
	data = controller.GetExperienceList()
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "ExperienceList获取成功",
		"data":    data,
	})
}

// AddExperience 添加新的经验
// @Tags 经验模块
// @Summary 添加新的经验
// @Accept multipart/form-data
// @Param name formData string false "经验名字"
// @Param url formData string false "经验网址"
// @Param description formData string false "经验描述"
// @Param time formData string false "经验时间"
// @Param image formData file false "经验图片"
// @Success 200 {string} json{"code","message"}
// @Router /experience/addExperience [post]
func AddExperience(c *gin.Context) {
	experience := &models.Experience{}
	experience.Name = c.PostForm("name")
	experience.Url = c.PostForm("url")
	experience.Description = c.PostForm("description")
	experience.Time = c.PostForm("time")
	file, err := c.FormFile("image")
	if err == nil {
		// 读取文件内容并赋值给 Image 字段
		imageFile, _ := file.Open()
		imageBytes := make([]byte, file.Size)
		imageFile.Read(imageBytes)
		experience.Image = imageBytes
	}
	controller.CreateExperience(experience)
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "添加经验成功！！！",
	})
}
