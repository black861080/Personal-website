package service

import "github.com/gin-gonic/gin"

func Yeyeye(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "耶耶耶！！！",
	})
}
