package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func GetBlogList() []*models.Blog {
	data := make([]*models.Blog, 10)
	util.DB.Find(&data)
	return data
}

func CreateBlog(blog *models.Blog) *gorm.DB {
	return util.DB.Create(blog)
}

func UpdateBlog(blog *models.Blog) *gorm.DB {
	return util.DB.Model(blog).Updates(models.Blog{
		Name:        blog.Name,
		Label:       blog.Label,
		Describe:    blog.Describe,
		Describe1:   blog.Describe1,
		Image1:      blog.Image1,
		Describe2:   blog.Describe2,
		Image2:      blog.Image2,
		Describe3:   blog.Describe3,
		Image3:      blog.Image3,
		Describe4:   blog.Describe4,
		Image4:      blog.Image4,
		Describe5:   blog.Describe5,
		Image5:      blog.Image5,
		Describe6:   blog.Describe6,
		Image6:      blog.Image6,
		Describe7:   blog.Describe7,
		Image7:      blog.Image7,
		Describe8:   blog.Describe8,
		Image8:      blog.Image8,
		Describe9:   blog.Describe9,
		Image9:      blog.Image9,
		Describe10:  blog.Describe10,
		Upload_time: blog.Upload_time,
	})
}

func DeleteBlog(blog *models.Blog) *gorm.DB {
	return util.DB.Where("id = ?", blog.ID).Delete(blog)
}

func FindBlog(blog *models.Blog) *models.Blog {
	var result *models.Blog
	util.DB.Where("id = ?", blog.ID).First(&result)
	return result
}
