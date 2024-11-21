package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GetProject 获取项目信息集
// @Tags 项目模块
// @Summary 获取所有项目信息
// @Success 200 {string} json{"code","message"}
// @Router /project/getProjectList [get]
func GetProject(c *gin.Context) {
	data := make([]*models.Project, 10)
	data = controller.GetProjectList()
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "ProjectList获取成功",
		"data":    data,
	})

}

// AddProject 添加新的项目
// @Tags 项目模块
// @Summary 添加新的项目
// @Accept multipart/form-data
// @Param name formData string false "项目名称"
// @Param url formData string false "项目 URL"
// @Param describe formData string false "项目描述"
// @Param describe1 formData string false "项目描述1"
// @Param image1 formData file false "项目图片1"
// @Param describe2 formData string false "项目描述2"
// @Param image2 formData file false "项目图片2"
// @Param describe3 formData string false "项目描述3"
// @Param image3 formData file false "项目图片3"
// @Param describe4 formData string false "项目描述4"
// @Param image4 formData file false "项目图片4"
// @Param describe5 formData string false "项目描述5"
// @Param image5 formData file false "项目图片5"
// @Param describe6 formData string false "项目描述6"
// @Param image6 formData file false "项目图片6"
// @Param describe7 formData string false "项目描述7"
// @Param image7 formData file false "项目图片7"
// @Param describe8 formData string false "项目描述8"
// @Param image8 formData file false "项目图片8"
// @Param describe9 formData string false "项目描述9"
// @Param image9 formData file false "项目图片9"
// @Param describe10 formData string false "项目描述10"
// @Success 200 {string} json{"code","message"}
// @Router /project/addProject [post]
func AddProject(c *gin.Context) {
	project := &models.Project{} // 将 project 声明为指针类型
	project.Name = c.PostForm("name")
	project.URL = c.PostForm("url")
	project.Upload_time = time.Now()
	file1, err1 := c.FormFile("image1")
	file2, err2 := c.FormFile("image2")
	file3, err3 := c.FormFile("image3")
	file4, err4 := c.FormFile("image4")
	file5, err5 := c.FormFile("image5")
	file6, err6 := c.FormFile("image6")
	file7, err7 := c.FormFile("image7")
	file8, err8 := c.FormFile("image8")
	file9, err9 := c.FormFile("image9")
	if err1 == nil {
		imageFile1, _ := file1.Open()
		imageBytes1 := make([]byte, file1.Size)
		imageFile1.Read(imageBytes1)
		project.Image1 = imageBytes1
	}
	if err2 == nil {
		imageFile2, _ := file2.Open()
		imageBytes2 := make([]byte, file2.Size)
		imageFile2.Read(imageBytes2)
		project.Image2 = imageBytes2
	}
	if err3 == nil {
		imageFile3, _ := file3.Open()
		imageBytes3 := make([]byte, file3.Size)
		imageFile3.Read(imageBytes3)
		project.Image3 = imageBytes3
	}
	if err4 == nil {
		imageFile4, _ := file4.Open()
		imageBytes4 := make([]byte, file4.Size)
		imageFile4.Read(imageBytes4)
		project.Image4 = imageBytes4
	}
	if err5 == nil {
		imageFile5, _ := file5.Open()
		imageBytes5 := make([]byte, file5.Size)
		imageFile5.Read(imageBytes5)
		project.Image5 = imageBytes5
	}
	if err6 == nil {
		imageFile6, _ := file6.Open()
		imageBytes6 := make([]byte, file6.Size)
		imageFile6.Read(imageBytes6)
		project.Image6 = imageBytes6
	}
	if err7 == nil {
		imageFile7, _ := file7.Open()
		imageBytes7 := make([]byte, file7.Size)
		imageFile7.Read(imageBytes7)
		project.Image7 = imageBytes7
	}
	if err8 == nil {
		imageFile8, _ := file8.Open()
		imageBytes8 := make([]byte, file8.Size)
		imageFile8.Read(imageBytes8)
		project.Image8 = imageBytes8
	}
	if err9 == nil {
		imageFile9, _ := file9.Open()
		imageBytes9 := make([]byte, file9.Size)
		imageFile9.Read(imageBytes9)
		project.Image9 = imageBytes9
	}
	project.Describe = c.PostForm("describe")
	project.Describe1 = c.PostForm("describe1")
	project.Describe2 = c.PostForm("describe2")
	project.Describe3 = c.PostForm("describe3")
	project.Describe4 = c.PostForm("describe4")
	project.Describe5 = c.PostForm("describe5")
	project.Describe6 = c.PostForm("describe6")
	project.Describe7 = c.PostForm("describe7")
	project.Describe8 = c.PostForm("describe8")
	project.Describe9 = c.PostForm("describe9")
	project.Describe10 = c.PostForm("describe10")
	controller.CreateProject(project)

	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "创建新项目成功！！！",
	})
}

// DeleteProject 删除选中项目
// @Tags 项目模块
// @Summary 删除项目
// @Param id query string false "项目ID"
// @Success 200 {string} json{"code","message"}
// @Router /project/deleteProject [post]
func DeleteProject(c *gin.Context) {
	project := &models.Project{}
	id, _ := strconv.Atoi(c.Query("id"))
	project.ID = uint(id)
	controller.DeleteProject(project)
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "删除成功！！！！",
	})
}

// UpdateProject 修改指定的项目
// @Tags 项目模块
// @Summary 修改指定的项目
// @Accept multipart/form-data
// @Param id formData string false "项目ID"
// @Param name formData string false "项目名称"
// @Param url formData string false "项目 URL"
// @Param describe formData string false "项目描述"
// @Param describe1 formData string false "项目描述1"
// @Param image1 formData file false "项目图片1"
// @Param describe2 formData string false "项目描述2"
// @Param image2 formData file false "项目图片2"
// @Param describe3 formData string false "项目描述3"
// @Param image3 formData file false "项目图片3"
// @Param describe4 formData string false "项目描述4"
// @Param image4 formData file false "项目图片4"
// @Param describe5 formData string false "项目描述5"
// @Param image5 formData file false "项目图片5"
// @Param describe6 formData string false "项目描述6"
// @Param image6 formData file false "项目图片6"
// @Param describe7 formData string false "项目描述7"
// @Param image7 formData file false "项目图片7"
// @Param describe8 formData string false "项目描述8"
// @Param image8 formData file false "项目图片8"
// @Param describe9 formData string false "项目描述9"
// @Param image9 formData file false "项目图片9"
// @Param describe10 formData string false "项目描述10"
// @Success 200 {string} json{"code","message"}
// @Router /project/updateProject [post]
func UpdateProject(c *gin.Context) {
	project := &models.Project{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	project.ID = uint(id)
	project.Name = c.PostForm("name")
	project.URL = c.PostForm("url")
	project.Upload_time = time.Now()
	// 获取上传的文件
	file1, err1 := c.FormFile("image1")
	file2, err2 := c.FormFile("image2")
	file3, err3 := c.FormFile("image3")
	file4, err4 := c.FormFile("image4")
	file5, err5 := c.FormFile("image5")
	file6, err6 := c.FormFile("image6")
	file7, err7 := c.FormFile("image7")
	file8, err8 := c.FormFile("image8")
	file9, err9 := c.FormFile("image9")
	if err1 == nil {
		imageFile1, _ := file1.Open()
		imageBytes1 := make([]byte, file1.Size)
		imageFile1.Read(imageBytes1)
		project.Image1 = imageBytes1
	}
	if err2 == nil {
		imageFile2, _ := file2.Open()
		imageBytes2 := make([]byte, file2.Size)
		imageFile2.Read(imageBytes2)
		project.Image2 = imageBytes2
	}
	if err3 == nil {
		imageFile3, _ := file3.Open()
		imageBytes3 := make([]byte, file3.Size)
		imageFile3.Read(imageBytes3)
		project.Image3 = imageBytes3
	}
	if err4 == nil {
		imageFile4, _ := file4.Open()
		imageBytes4 := make([]byte, file4.Size)
		imageFile4.Read(imageBytes4)
		project.Image4 = imageBytes4
	}
	if err5 == nil {
		imageFile5, _ := file5.Open()
		imageBytes5 := make([]byte, file5.Size)
		imageFile5.Read(imageBytes5)
		project.Image5 = imageBytes5
	}
	if err6 == nil {
		imageFile6, _ := file6.Open()
		imageBytes6 := make([]byte, file6.Size)
		imageFile6.Read(imageBytes6)
		project.Image6 = imageBytes6
	}
	if err7 == nil {
		imageFile7, _ := file7.Open()
		imageBytes7 := make([]byte, file7.Size)
		imageFile7.Read(imageBytes7)
		project.Image7 = imageBytes7
	}
	if err8 == nil {
		imageFile8, _ := file8.Open()
		imageBytes8 := make([]byte, file8.Size)
		imageFile8.Read(imageBytes8)
		project.Image8 = imageBytes8
	}
	if err9 == nil {
		imageFile9, _ := file9.Open()
		imageBytes9 := make([]byte, file9.Size)
		imageFile9.Read(imageBytes9)
		project.Image9 = imageBytes9
	}
	project.Describe = c.PostForm("describe")
	project.Describe1 = c.PostForm("describe1")
	project.Describe2 = c.PostForm("describe2")
	project.Describe3 = c.PostForm("describe3")
	project.Describe4 = c.PostForm("describe4")
	project.Describe5 = c.PostForm("describe5")
	project.Describe6 = c.PostForm("describe6")
	project.Describe7 = c.PostForm("describe7")
	project.Describe8 = c.PostForm("describe8")
	project.Describe9 = c.PostForm("describe9")
	project.Describe10 = c.PostForm("describe10")
	controller.UpdateProject(project)

	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "项目更新成功！！！",
	})
}

// GetTheProject 获取指定的项目
// @Tags 项目模块
// @Summary 获取指定的项目
// @Accept multipart/form-data
// @Param id formData string false "项目ID"
// @Success 200 {string} json{"code","message"}
// @Router /project/getTheProject [post]
func GetTheProject(c *gin.Context) {
	project := &models.Project{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	project.ID = uint(id)
	good_project := controller.GetById(project)
	if good_project == nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "未查询到project",
		})
		return
	}
	if good_project.ID == 0 {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "未查询到project",
		})
		return
	}
	c.JSON(200, gin.H{
		"code":    1,
		"data":    good_project,
		"message": "查询成功！！！",
	})
}
