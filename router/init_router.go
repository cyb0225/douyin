// 配置路由
package router

import (
	"log"

	"github.com/2103561941/douyin/user/controller"

	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	g := engine.Group("/douyin")
	user := g.Group("/user")
	user.POST("/register", controller.Register)
	log.Println("into router")
}
