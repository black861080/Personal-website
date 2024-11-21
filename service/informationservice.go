package service

import (
	"black/controller"
	"black/models"
	"fmt"
	"github.com/gin-gonic/gin"
)

// FindInformation 获取信息
// @Tags 信息模块
// @Summary 获取信息
// @Accept multipart/form-data
// @Success 200 {string} json{"code","message"}
// @Router /information/getInformation [get]
func FindInformation(c *gin.Context) {
	data := controller.FirstInformation()
	c.JSON(200, gin.H{
		"code": 0,
		"data": data,
	})
}

// AddInformation 添加新的信息
// @Tags 信息模块
// @Summary 添加新的信息
// @Accept multipart/form-data
// @Param name formData string false "信息名字"
// @Param nickname formData string false "信息昵称"
// @Param location formData string false "信息位置"
// @Param street formData string false "信息街道"
// @Param description formData string false "信息描述"
// @Param webDescription formData string false "网站描述"
// @Param projectDescription formData string false "项目描述"
// @Param photoDescription formData string false "照片描述"
// @Param blogDescription formData string false "博客描述"
// @Param suggestionDescription formData string false "建议描述"
// @Param image formData file false "信息图片"
// @Success 200 {string} json{"code","message"}
// @Router /information/addInformation [post]
func AddInformation(c *gin.Context) {
	information := &models.Information{}
	information.Name = c.PostForm("name")
	information.NikeName = c.PostForm("nickname")
	fmt.Println(1)
	fmt.Println(c.PostForm("nickname"))
	fmt.Println(1)
	information.Location = c.PostForm("location")
	information.Street = c.PostForm("street")
	information.Description = c.PostForm("description")
	information.WebDescription = c.PostForm("webDescription")
	information.ProjectDescription = c.PostForm("projectDescription")
	information.PhotoDescription = c.PostForm("photoDescription")
	information.BlogDescription = c.PostForm("blogDescription")
	information.SuggestionDescription = c.PostForm("suggestionDescription")
	file, err := c.FormFile("image")
	if err == nil {
		imageFile, _ := file.Open()
		imageBytes := make([]byte, file.Size)
		imageFile.Read(imageBytes)
		information.Image = imageBytes
	}
	controller.CreateInformation(information)
	c.JSON(200, gin.H{
		"code":    0,
		"message": "创建成功！！！",
	})
}
