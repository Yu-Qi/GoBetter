package api

import (
	"github.com/Yu-Qi/GoBetter/NFT/pkg/db/model"
	"github.com/Yu-Qi/GoBetter/NFT/pkg/service/rdb"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	user := model.User{
		UID:             c.Param("uid"),
		Name:            "test",
		EthereumAddress: "0x123",
	}
	result := rdb.RestfulDB.Create(&user)
	if result.Error != nil {
		c.JSON(500, gin.H{
			"error": result.Error,
		})
		return
	}
	c.JSON(200, gin.H{
		"data": "insert success",
	})
}

func GetOneUser(c *gin.Context) {
	uid := c.Param("uid")
	var user model.User
	rdb.RestfulDB.Take(&user, uid)
	c.JSON(200, gin.H{
		"data": user,
	})
}

func GetUsers(c *gin.Context) {
	// TODO: paging
	uid := c.Param("uid")
	var user model.User
	rdb.RestfulDB.Take(&user, uid)
	c.JSON(200, gin.H{
		"data": user,
	})
}

func UpdateUser(c *gin.Context) {
	// uid := c.Param("uid")
	c.JSON(200, gin.H{
		"data": "pong",
	})
}

func DeleteUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "pong",
	})
}
