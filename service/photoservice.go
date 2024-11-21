package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GetPhoto 获取照片信息集
// @Tags 照片模块
// @Summary 获取所有照片信息
// @Success 200 {string} json{"code","message"}
// @Router /photo/getPhotoList [get]
func GetPhoto(c *gin.Context) {
	data := make([]*models.Photo, 10)
	data = controller.GetPhotoList()
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "PhotoList获取成功",
		"data":    data,
	})
}

// AddPhoto 添加新的照片
// @Tags 照片模块
// @Summary 添加新的照片
// @Accept multipart/form-data
// @Param name formData string false "照片名称"
// @Param image formData file false "照片"
// @Param description formData string false "照片描述"
// @Success 200 {string} json{"code","message"}
// @Router /photo/addPhoto [post]
func AddPhoto(c *gin.Context) {
	photo := &models.Photo{} // 将 project 声明为指针类型
	photo.PhotoName = c.PostForm("name")
	photo.Upload_time = time.Now()
	file, err := c.FormFile("image")
	if err == nil {
		// 读取文件内容并赋值给 Image 字段
		imageFile, _ := file.Open()
		imageBytes := make([]byte, file.Size)
		imageFile.Read(imageBytes)
		photo.Image = imageBytes
	}

	photo.Description = c.PostForm("description")
	controller.CreatePhoto(photo)

	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "添加照片成功！！！",
	})
}

// DeletePhoto 删除选中照片
// @Tags 照片模块
// @Summary 删除照片
// @Param id query string false "照片ID"
// @Success 200 {string} json{"code","message"}
// @Router /photo/deletePhoto [post]
func DeletePhoto(c *gin.Context) {
	photo := &models.Photo{}
	id, _ := strconv.Atoi(c.Query("id"))
	photo.ID = uint(id)
	controller.DeletePhoto(photo)
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "删除成功！！！！",
	})
}

// UpdatePhoto 修改指定的照片
// @Tags 照片模块
// @Summary 修改指定的照片
// @Accept multipart/form-data
// @Param id formData string false "照片ID"
// @Param name formData string false "照片名称"
// @Param image formData file false "照片"
// @Param description formData string false "照片描述"
// @Success 200 {string} json{"code","message"}
// @Router /photo/updatePhoto [post]
func UpdatePhoto(c *gin.Context) {
	photo := &models.Photo{} // 将 project 声明为指针类型
	id, _ := strconv.Atoi(c.PostForm("id"))
	photo.ID = uint(id)
	photo.PhotoName = c.PostForm("name")
	file, err := c.FormFile("image")
	if err == nil {
		// 读取文件内容并赋值给 Image 字段
		imageFile, _ := file.Open()
		imageBytes := make([]byte, file.Size)
		imageFile.Read(imageBytes)
		photo.Image = imageBytes
	}

	photo.Description = c.PostForm("description")
	controller.UpdatePhoto(photo)

	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "照片更新成功！！！",
	})
}
