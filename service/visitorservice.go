package service

import (
	"black/controller"
	"black/models"
	"black/util"
	"fmt"
	"github.com/gin-gonic/gin"
	"unicode/utf8"
)

// Login 寻找访问者信息
// @Tags 访问者模块
// @Summary 查询访问者信息
// @Accept multipart/form-data
// @Param username formData string false "请输入用户名"
// @Param password formData string false "请输入密码"
// @Success 200 {string} json{"code","message"}
// @Router /visitor/findVisitor [post]
func Login(c *gin.Context) {
	visitor := &models.Visitor{}
	visitor.Username = c.PostForm("username")
	visitor.Password = c.PostForm("password")
	goodVisitor := controller.FindVisitor(visitor)

	if goodVisitor.ID != 0 {
		// 生成 Token
		token, err := util.GenerateToken(goodVisitor.ID)
		if err != nil {
			c.JSON(500, gin.H{
				"code":    -1,
				"message": "Token 生成失败",
			})
			return
		}

		// 更新 Visitor 的 Token 字段
		goodVisitor.Token = token
		util.DB.Save(&goodVisitor)

		c.JSON(200, gin.H{
			"code":    0,
			"message": "登录成功！！！",
			"data":    goodVisitor,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或密码错误",
		})
	}
}

// Register 注册
// @Tags 访问者模块
// @Summary 注册
// @Accept multipart/form-data
// @Param username formData string false "请输入用户名"
// @Param password formData string false "请输入密码"
// @Success 200 {string} json{"code","message"}
// @Router /visitor/registerVisitor [post]
func Register(c *gin.Context) {
	visitor := &models.Visitor{}
	if c.PostForm("username") == "" || c.PostForm("password") == "" {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名或密码不能为空！！！",
		})
		return
	}

	// 计算用户名字符数
	username := c.PostForm("username")
	if utf8.RuneCountInString(username) < 6 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "用户名不少于6个字符",
		})
		return
	}

	visitor.Username = username
	goodVisitor := controller.CheckUsername(visitor)
	if goodVisitor.ID != 0 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "该用户已存在！！！",
		})
		return
	}

	visitor.Password = c.PostForm("password")
	if len(visitor.Password) < 6 {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "密码不少于6个字符",
		})
		return
	}

	visitor.Role = "visitor"
	controller.CreateVisitor(visitor)

	c.JSON(200, gin.H{
		"code":    0,
		"message": "注册成功！！！",
	})
}

// Check 寻找访问者角色
// @Tags 访问者模块
// @Summary 查询访问者角色
// @Accept multipart/form-data
// @Param token formData string false "请输入Token"
// @Success 200 {string} json{"code","message"}
// @Router /visitor/check [post]
func Check(c *gin.Context) {
	visitor := &models.Visitor{}
	visitor.Token = c.PostForm("token")
	goodVisitor := controller.FindRole(visitor)
	fmt.Println(goodVisitor)
	if goodVisitor.ID != 0 {

		c.JSON(200, gin.H{
			"code":    0,
			"message": "查询成功！！！",
			"data":    goodVisitor,
		})
	} else {
		c.JSON(200, gin.H{
			"code":    -1,
			"message": "查询失败",
		})
	}
}
