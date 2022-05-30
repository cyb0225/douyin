// configure router
package router

import (
	"github.com/2103561941/douyin/controller/userctl"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	apiRouter := engine.Group("/douyin")
	user := apiRouter.Group("/user")
	user.POST("/register/", userctl.Register)
	user.POST("/login/", userctl.Login)
	user.GET("/", userctl.GetUserInfo)

	relation := apiRouter.Group(("/relation"))
	relation.POST("/action/", userctl.Follow)
	relation.GET("/follow/list/", userctl.FollowList)
	relation.GET("/follower/list/", userctl.FollowerList)
}
