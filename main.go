package main

import (
	"github.com/2103561941/douyin/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.InitRouter(r)

	r.Run(":9999")
}
