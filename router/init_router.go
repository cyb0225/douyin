// 配置路由
package router

import (
	"github.com/2103561941/douyin/user/controller"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	apiRouter := engine.Group("/douyin")

	user := apiRouter.Group("/user")
	user.POST("/register/", controller.Register)
	user.POST("/login/", controller.Login)
	user.GET("/", controller.GetUserInfo)
	apiRouter.GET("/publish/list", controller.Publish_list)
	apiRouter.GET("/feed/", controller.Feed)
	apiRouter.POST("/publish/action", controller.Publish)
}
