package router

import (
	"black/docs"
	"black/service"
	"black/util"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.Use(util.CorsHandler())
	/*r.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders: []string{"Origin", "Content-Type"},
	}))*/
	docs.SwaggerInfo.BasePath = ""
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.GET("/index", service.GetIndex)
	r.GET("/", service.Yeyeye)
	//项目模块
	r.GET("/project/getProjectList", service.GetProject)
	r.POST("/project/addProject", service.AddProject)
	r.POST("/project/deleteProject", service.DeleteProject)
	r.POST("/project/updateProject", service.UpdateProject)
	r.POST("/project/getTheProject", service.GetTheProject)
	//照片模块
	r.GET("/photo/getPhotoList", service.GetPhoto)
	r.POST("/photo/addPhoto", service.AddPhoto)
	r.POST("/photo/deletePhoto", service.DeletePhoto)
	r.POST("/photo/updatePhoto", service.UpdatePhoto)
	//留言模块
	r.GET("/message/getMessageList", service.GetMessage)
	r.POST("/message/addMessage", service.AddMessage)
	//朋友模块
	r.POST("/friend/findFriend", service.FindFriend)
	//访问者模块
	r.POST("/visitor/findVisitor", service.Login)
	r.POST("/visitor/check", service.Check)
	r.POST("/visitor/registerVisitor", service.Register)
	//信息模块
	r.GET("/information/getInformation", service.FindInformation)
	r.POST("/information/addInformation", service.AddInformation)
	//技能模块
	r.GET("/skill/getSkillList", service.GetSkill)
	r.POST("/skill/addSkill", service.AddSkill)
	//经验模块
	r.GET("/experience/getExperienceList", service.GetExperience)
	r.POST("/experience/addExperience", service.AddExperience)
	//blog模块
	r.GET("/blog/getBlogList", service.GetBlog)
	r.POST("/blog/addBlog", service.AddBlog)
	r.POST("/blog/deleteBlog", service.DeleteBlog)
	r.POST("/blog/updateBlog", service.UpdateBlog)
	r.POST("/blog/getTheBlog", service.GetTheBlog)
	return r
}
