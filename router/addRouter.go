// 配置路由
package addrouter

import (
	"github.com/2103561941/douyin/vedioctl"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	apiRouter := engine.Group("/douyin")

	user := apiRouter.Group("/user")
	user.POST("/register/", vedioctl.Register)
	user.POST("/login/", vedioctl.Login)
	user.GET("/", vedioctl.GetUserInfo)
	apiRouter.GET("/publish/list", vedioctl.Publish_list)
	apiRouter.GET("/feed/", vedioctl.Feed)
	apiRouter.POST("/publish/action", vedioctl.Publish)
}
