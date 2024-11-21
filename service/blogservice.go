package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

// GetBlog 获取blog信息集
// @Tags blog模块
// @Summary 获取所有blog信息
// @Success 200 {string} json{"code","message"}
// @Router /blog/getBlogList [get]
func GetBlog(c *gin.Context) {
	data := make([]*models.Blog, 10)
	data = controller.GetBlogList()
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "BlogList获取成功",
		"data":    data,
	})
}

// AddBlog 添加新的blog
// @Tags blog模块
// @Summary 添加新的blog
// @Accept multipart/form-data
// @Param name formData string false "blog名称"
// @Param label formData string false "blog标签"
// @Param describe formData string false "blog描述"
// @Param describe1 formData string false "blog描述1"
// @Param image1 formData file false "blog图片1"
// @Param describe2 formData string false "blog描述2"
// @Param image2 formData file false "blog图片2"
// @Param describe3 formData string false "blog描述3"
// @Param image3 formData file false "blog图片3"
// @Param describe4 formData string false "blog描述4"
// @Param image4 formData file false "blog图片4"
// @Param describe5 formData string false "blog描述5"
// @Param image5 formData file false "blog图片5"
// @Param describe6 formData string false "blog描述6"
// @Param image6 formData file false "blog图片6"
// @Param describe7 formData string false "blog描述7"
// @Param image7 formData file false "blog图片7"
// @Param describe8 formData string false "blog描述8"
// @Param image8 formData file false "blog图片8"
// @Param describe9 formData string false "blog描述9"
// @Param image9 formData file false "blog图片9"
// @Param describe10 formData string false "blog描述10"
// @Success 200 {string} json{"code","message"}
// @Router /blog/addBlog [post]
func AddBlog(c *gin.Context) {
	blog := &models.Blog{} // 将 project 声明为指针类型
	blog.Name = c.PostForm("name")
	blog.Label = c.PostForm("label")
	blog.Upload_time = time.Now()
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
		// 读取文件内容并赋值给 Image 字段
		imageFile1, _ := file1.Open()
		imageBytes1 := make([]byte, file1.Size)
		imageFile1.Read(imageBytes1)
		blog.Image1 = imageBytes1
	}
	if err2 == nil {
		// 读取文件内容并赋值给 Image 字段
		imageFile2, _ := file2.Open()
		imageBytes2 := make([]byte, file2.Size)
		imageFile2.Read(imageBytes2)
		blog.Image2 = imageBytes2
	}
	if err3 == nil {
		imageFile3, _ := file3.Open()
		imageBytes3 := make([]byte, file3.Size)
		imageFile3.Read(imageBytes3)
		blog.Image3 = imageBytes3
	}
	if err4 == nil {
		imageFile4, _ := file4.Open()
		imageBytes4 := make([]byte, file4.Size)
		imageFile4.Read(imageBytes4)
		blog.Image4 = imageBytes4
	}
	if err5 == nil {
		imageFile5, _ := file5.Open()
		imageBytes5 := make([]byte, file5.Size)
		imageFile5.Read(imageBytes5)
		blog.Image5 = imageBytes5
	}
	if err6 == nil {
		imageFile6, _ := file6.Open()
		imageBytes6 := make([]byte, file6.Size)
		imageFile6.Read(imageBytes6)
		blog.Image6 = imageBytes6
	}
	if err7 == nil {
		imageFile7, _ := file7.Open()
		imageBytes7 := make([]byte, file7.Size)
		imageFile7.Read(imageBytes7)
		blog.Image7 = imageBytes7
	}
	if err8 == nil {
		imageFile8, _ := file8.Open()
		imageBytes8 := make([]byte, file8.Size)
		imageFile8.Read(imageBytes8)
		blog.Image8 = imageBytes8
	}
	if err9 == nil {
		imageFile9, _ := file9.Open()
		imageBytes9 := make([]byte, file9.Size)
		imageFile9.Read(imageBytes9)
		blog.Image9 = imageBytes9
	}
	blog.Describe = c.PostForm("describe")
	blog.Describe1 = c.PostForm("describe1")
	blog.Describe2 = c.PostForm("describe2")
	blog.Describe3 = c.PostForm("describe3")
	blog.Describe4 = c.PostForm("describe4")
	blog.Describe5 = c.PostForm("describe5")
	blog.Describe6 = c.PostForm("describe6")
	blog.Describe7 = c.PostForm("describe7")
	blog.Describe8 = c.PostForm("describe8")
	blog.Describe9 = c.PostForm("describe9")
	blog.Describe10 = c.PostForm("describe10")

	controller.CreateBlog(blog)

	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "创建新blog成功！！！",
	})
}

// DeleteBlog 删除选中blog
// @Tags blog模块
// @Summary 删除blog
// @Param id query string false "blogID"
// @Success 200 {string} json{"code","message"}
// @Router /blog/deleteBlog [post]
func DeleteBlog(c *gin.Context) {
	blog := &models.Blog{}
	id, _ := strconv.Atoi(c.Query("id"))
	blog.ID = uint(id)
	controller.DeleteBlog(blog)
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "删除成功！！！！",
	})
}

// UpdateBlog 修改指定的blog
// @Tags blog模块
// @Summary 修改指定的blog
// @Accept multipart/form-data
// @Param id formData string false "blogID"
// @Param name formData string false "blog名称"
// @Param label formData string false "blog标签"
// @Param describe formData string false "blog描述"
// @Param describe1 formData string false "blog描述1"
// @Param image1 formData file false "blog图片1"
// @Param describe2 formData string false "blog描述2"
// @Param image2 formData file false "blog图片2"
// @Param describe3 formData string false "blog描述3"
// @Param image3 formData file false "blog图片3"
// @Param describe4 formData string false "blog描述4"
// @Param image4 formData file false "blog图片4"
// @Param describe5 formData string false "blog描述5"
// @Param image5 formData file false "blog图片5"
// @Param describe6 formData string false "blog描述6"
// @Param image6 formData file false "blog图片6"
// @Param describe7 formData string false "blog描述7"
// @Param image7 formData file false "blog图片7"
// @Param describe8 formData string false "blog描述8"
// @Param image8 formData file false "blog图片8"
// @Param describe9 formData string false "blog描述9"
// @Param image9 formData file false "blog图片9"
// @Param describe10 formData string false "blog描述10"
// @Success 200 {string} json{"code","message"}
// @Router /blog/updateBlog [post]
func UpdateBlog(c *gin.Context) {
	blog := &models.Blog{} // 将 project 声明为指针类型
	id, _ := strconv.Atoi(c.PostForm("id"))
	blog.ID = uint(id)
	blog.Name = c.PostForm("name")
	blog.Label = c.PostForm("label")
	blog.Upload_time = time.Now()
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
		// 读取文件内容并赋值给 Image 字段
		imageFile1, _ := file1.Open()
		imageBytes1 := make([]byte, file1.Size)
		imageFile1.Read(imageBytes1)
		blog.Image1 = imageBytes1
	}
	if err2 == nil {
		// 读取文件内容并赋值给 Image 字段
		imageFile2, _ := file2.Open()
		imageBytes2 := make([]byte, file2.Size)
		imageFile2.Read(imageBytes2)
		blog.Image2 = imageBytes2
	}
	if err3 == nil {
		imageFile3, _ := file3.Open()
		imageBytes3 := make([]byte, file3.Size)
		imageFile3.Read(imageBytes3)
		blog.Image3 = imageBytes3
	}
	if err4 == nil {
		imageFile4, _ := file4.Open()
		imageBytes4 := make([]byte, file4.Size)
		imageFile4.Read(imageBytes4)
		blog.Image4 = imageBytes4
	}
	if err5 == nil {
		imageFile5, _ := file5.Open()
		imageBytes5 := make([]byte, file5.Size)
		imageFile5.Read(imageBytes5)
		blog.Image5 = imageBytes5
	}
	if err6 == nil {
		imageFile6, _ := file6.Open()
		imageBytes6 := make([]byte, file6.Size)
		imageFile6.Read(imageBytes6)
		blog.Image6 = imageBytes6
	}
	if err7 == nil {
		imageFile7, _ := file7.Open()
		imageBytes7 := make([]byte, file7.Size)
		imageFile7.Read(imageBytes7)
		blog.Image7 = imageBytes7
	}
	if err8 == nil {
		imageFile8, _ := file8.Open()
		imageBytes8 := make([]byte, file8.Size)
		imageFile8.Read(imageBytes8)
		blog.Image8 = imageBytes8
	}
	if err9 == nil {
		imageFile9, _ := file9.Open()
		imageBytes9 := make([]byte, file9.Size)
		imageFile9.Read(imageBytes9)
		blog.Image9 = imageBytes9
	}
	blog.Describe = c.PostForm("describe")
	blog.Describe1 = c.PostForm("describe1")
	blog.Describe2 = c.PostForm("describe2")
	blog.Describe3 = c.PostForm("describe3")
	blog.Describe4 = c.PostForm("describe4")
	blog.Describe5 = c.PostForm("describe5")
	blog.Describe6 = c.PostForm("describe6")
	blog.Describe7 = c.PostForm("describe7")
	blog.Describe8 = c.PostForm("describe8")
	blog.Describe9 = c.PostForm("describe9")
	blog.Describe10 = c.PostForm("describe10")
	controller.UpdateBlog(blog)

	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "blog更新成功！！！",
	})
}

// GetTheBlog 获取指定的blog
// @Tags blog模块
// @Summary 获取指定的blog
// @Accept multipart/form-data
// @Param id formData string false "blogID"
// @Success 200 {string} json{"code","message"}
// @Router /blog/getTheBlog [post]
func GetTheBlog(c *gin.Context) {
	blog := &models.Blog{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	blog.ID = uint(id)
	good_blog := controller.FindBlog(blog)
	if good_blog == nil {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "查询失败！！！",
		})
	}
	if good_blog.ID == 0 {
		c.JSON(200, gin.H{
			"code":    0,
			"message": "查询失败！！！",
		})
	}
	c.JSON(200, gin.H{
		"code":    1,
		"data":    good_blog,
		"message": "查询成功！！！",
	})
}
