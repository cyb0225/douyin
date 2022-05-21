package controller

import (
	"log"

	"github.com/2103561941/douyin/user/server"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")
	c.JSON(200, gin.H{
		"username": username,
		"password": password,
	})
	server.Register(&server.UserRegInfo{
		Username: username,
		Password: password,
	})

	log.Println("into controller")

}
