package main

import (
	"fmt"

	"github.com/2103561941/douyin/router"
	"github.com/2103561941/douyin/user/repository"
	"github.com/gin-gonic/gin"
	// "github.com/pilu/fresh" // 热加载调试，更新服务器内容就可以不用杀死服务器进程，直接刷新就可以
)




func main() {
	// 数据库初始化
	if err := repository.InitUserRep(); err != nil {
		fmt.Println(err)
		return
	}

	r := gin.Default()
	router.InitRouter(r)

	r.Run(":9999")
}
