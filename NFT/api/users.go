package api

import (
	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "pong",
	})
}

func GetUsers(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "pong",
	})
}

func UpdateUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "pong",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "pong",
	})
}
