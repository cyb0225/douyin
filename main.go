package main

import (
	"github.com/2103561941/douyin/controller/videoctl"
	"github.com/2103561941/douyin/repository"
	"github.com/2103561941/douyin/router"
	"github.com/gin-gonic/gin"
)

func main() {
	// 下面这两个初始化需要去对应的init文件中更改用户密码等信息才可以使用。

	// 数据库初始化
	if err := repository.Init(); err != nil {
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
