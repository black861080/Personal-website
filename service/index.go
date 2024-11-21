package service

import "github.com/gin-gonic/gin"

// GetIndex swag测试
// @Tags swag测试
// @Success 200 {string} json{"code","message"}
// @Router /index [get]
func GetIndex(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "hello world",
	})
}
