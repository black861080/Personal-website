package service

import (
	"black/controller"
	"black/models"
	"github.com/gin-gonic/gin"
)

// GetSkill 获取技能集
// @Tags 技能模块
// @Summary 获取所有技能信息
// @Success 200 {string} json{"code","message"}
// @Router /skill/getSkillList [get]
func GetSkill(c *gin.Context) {
	data := make([]*models.Skill, 10)
	data = controller.GetSkillList()
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "SkillList获取成功",
		"data":    data,
	})
}

// AddSkill 添加新的技能
// @Tags 技能模块
// @Summary 添加新的技能
// @Accept multipart/form-data
// @Param name formData string false "技能名字"
// @Param url formData string false "技能网址"
// @Param description formData string false "技能描述"
// @Param time formData string false "技能时间"
// @Param image formData file false "技能图片"
// @Success 200 {string} json{"code","message"}
// @Router /skill/addSkill [post]
func AddSkill(c *gin.Context) {
	skill := &models.Skill{}
	skill.Name = c.PostForm("name")
	skill.Url = c.PostForm("url")
	skill.Description = c.PostForm("description")
	skill.Time = c.PostForm("time")
	file, err := c.FormFile("image")
	if err == nil {
		imageFile, _ := file.Open()
		imageBytes := make([]byte, file.Size)
		imageFile.Read(imageBytes)
		skill.Image = imageBytes
	}
	controller.CreateSkill(skill)
	c.JSON(200, gin.H{
		"code":    0, //成功
		"message": "添加技能成功！！！",
	})
}
