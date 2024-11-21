package controller

import (
	"black/models"
	"black/util"
	"gorm.io/gorm"
)

func GetProjectList() []*models.Project {
	data := make([]*models.Project, 10)
	util.DB.Find(&data)
	return data
}

func CreateProject(project *models.Project) *gorm.DB {
	return util.DB.Create(project)
}

func UpdateProject(project *models.Project) *gorm.DB {
	return util.DB.Model(project).Updates(models.Project{
		Name:        project.Name,
		URL:         project.URL,
		Describe:    project.Describe,
		Describe1:   project.Describe1,
		Image1:      project.Image1,
		Describe2:   project.Describe2,
		Image2:      project.Image2,
		Describe3:   project.Describe3,
		Image3:      project.Image3,
		Describe4:   project.Describe4,
		Image4:      project.Image4,
		Describe5:   project.Describe5,
		Image5:      project.Image5,
		Describe6:   project.Describe6,
		Image6:      project.Image6,
		Describe7:   project.Describe7,
		Image7:      project.Image7,
		Describe8:   project.Describe8,
		Image8:      project.Image8,
		Describe9:   project.Describe9,
		Image9:      project.Image9,
		Describe10:  project.Describe10,
		Upload_time: project.Upload_time,
	})
}

func DeleteProject(project *models.Project) *gorm.DB {
	return util.DB.Where("id = ?", project.ID).Delete(project)
}

func GetById(project *models.Project) *models.Project {
	var result *models.Project
	util.DB.Where("id = ?", project.ID).First(&result)
	return result
}
