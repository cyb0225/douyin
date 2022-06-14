package main

import (
	"github.com/2103561941/douyin/conf"
	"github.com/2103561941/douyin/controller/videoctl"
	"github.com/2103561941/douyin/repository"
	"github.com/2103561941/douyin/router"
	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/pprof"
)

func main() {
	// 导入配置
	if err := config.InitConfig(); err != nil {
		panic(err.Error())
	}

	// 连接数据库，创建表
	if err := repository.InitDatabase(); err != nil {
		panic(err.Error())
	}

	// 连接oss对象存储
	if err := videoctl.InitOss(); err != nil {
		panic(err.Error())
	}

	// 开启路由服务
	engine := gin.Default()
	router.InitRouter(engine)

	// 注册pprof的路由
	pprof.Register(engine) 
	
	engine.Run(":9999")

}
