package main

import (
	"github.com/2103561941/douyin/config" // 加载配置文件
	"github.com/2103561941/douyin/controller/videoctl"
	"github.com/2103561941/douyin/repository"
	"github.com/2103561941/douyin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 配置导入
	if err := config.InitConfig(); err != nil {
		panic(err.Error())
	}

	// 数据库初始化
	if err := repository.InitDatabase(); err != nil {
		panic(err.Error())
	}

	// 对象存储初始化
	if err := videoctl.InitOss(); err != nil {
		panic(err.Error())
	}

	// 开启服务
	r := gin.Default()
	router.InitRouter(r)

	r.Run(":9999")
}
