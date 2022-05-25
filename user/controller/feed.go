package controller

import (
	"github.com/2103561941/douyin/vedio/controller"
	"github.com/gin-gonic/gin"
)

type feed struct {
	Response
	NextTime int `json:"next_time"`
	controller.VideoInfo
}

func Feed(c *gin.Context) {
	//latestTime := c.Query("latest_time")
	//token := c.Query("token")
}
