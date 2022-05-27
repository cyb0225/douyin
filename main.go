package main

import (
	"fmt"

	"github.com/2103561941/douyin/router"
	"github.com/2103561941/douyin/repository"
	"github.com/gin-gonic/gin"
)




func main() {
	// 数据库初始化
	if err := repository.Init(); err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	router.InitRouter(r)

	r.Run(":9999")
}
