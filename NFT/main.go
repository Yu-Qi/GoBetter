package main

import (
	"github.com/Yu-Qi/GoBetter/NFT/api"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	registerUser(r)
	r.Run(":8123")
}

func registerUser(r *gin.Engine) {
	v1 := r.Group("/v1/users")
	v1.POST("/:uid", api.CreateUser)
	v1.GET("", api.GetUsers)
	v1.PUT("/:uid", api.UpdateUser)
	v1.DELETE("/:uid", api.DeleteUser)
}
