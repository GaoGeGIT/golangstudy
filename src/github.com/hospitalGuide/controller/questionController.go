package controller

import "github.com/gin-gonic/gin"

func Question(c *gin.Context){

	c.JSON(200, gin.H{
		"a":123,
	})
}
