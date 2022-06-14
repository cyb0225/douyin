// configure router
package router

import (
	"net/http"

	"github.com/2103561941/douyin/controller/userctl"
	"github.com/2103561941/douyin/controller/videoctl"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {
	
	// 主页面
	apiRouter := engine.Group("/douyin")

	apiRouter.GET("/feed", videoctl.Feed)

	// 用户
	user := apiRouter.Group("/user")
	user.POST("/register/", userctl.Register)
	user.POST("/login/", userctl.Login)
	user.GET("/", userctl.GetUserInfo)

	// 关注
	relation := apiRouter.Group(("/relation"))
	relation.POST("/action/", userctl.Follow)
	relation.GET("/follow/list/", userctl.FollowList)
	relation.GET("/follower/list/", userctl.FollowerList)

	// 投稿
	publish := apiRouter.Group("/publish")
	publish.POST("/action/", videoctl.Publish)
	publish.GET("/list/", videoctl.GetPublishList)

	apiRouter.StaticFS("/index/video", http.Dir("./video_content"))
	apiRouter.StaticFS("/index/cover", http.Dir("./cover_content"))


	// 点赞
	favorite := apiRouter.Group("/favorite")
	favorite.POST("/action/", videoctl.Like)
	favorite.GET("/list/", videoctl.GetLikeList)


	// 评论
	comment := apiRouter.Group("/comment")
	comment.POST("/action/", videoctl.Comment)
	comment.GET("/list/", videoctl.GetCommentList)

}
