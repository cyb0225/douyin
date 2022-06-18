// configure router
package router

import (
	"github.com/2103561941/douyin/controller/userctl"
	"github.com/2103561941/douyin/controller/videoctl"
	"github.com/2103561941/douyin/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouter(engine *gin.Engine) {

	// 主页面
	apiRouter := engine.Group("/douyin")
	apiRouter.POST("/user/register/", userctl.Register)
	apiRouter.POST("/user/login/", userctl.Login)
	{
		JWT := engine.Group("/douyin", middleware.JWTToken())
		JWT.GET("/feed", videoctl.Feed)

		// 用户
		//apiRouter.POST("/user/register/", userctl.Register)
		//apiRouter.POST("/user/login/", userctl.Login)
		JWT.GET("/user/", userctl.GetUserInfo)

		// 关注
		JWT.POST("/relation/action/", userctl.Follow)
		JWT.GET("/relation/follow/list/", userctl.FollowList)
		JWT.GET("/relation/follower/list/", userctl.FollowerList)

		// 投稿
		JWT.POST("/publish/action/", videoctl.Publish)
		JWT.GET("/publish/list/", videoctl.GetPublishList)

		//apiRouter.StaticFS("/index/video", http.Dir("./video_content"))
		//apiRouter.StaticFS("/index/cover", http.Dir("./cover_content"))

		// 点赞
		//favorite := apiRouter.Group("/favorite")
		JWT.POST("/favorite/action/", videoctl.Like)
		JWT.GET("/favorite/list/", videoctl.GetLikeList)

		// 评论
		//comment := apiRouter.Group("/comment")
		JWT.POST("/comment/action/", videoctl.Comment)
		JWT.GET("/comment/list/", videoctl.GetCommentList)
	}

}
