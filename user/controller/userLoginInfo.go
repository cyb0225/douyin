// 登录后进入的用户界面
package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户基本信息 用于json格式返回
type userJsonInfo struct {
	ID            int    `json:"id"`
	Username      string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
}

// 用户登录情况，记录
var (
	userLoginInfos = make(map[string]*userJsonInfo)
)

// 返回相应
type interfaceResponse struct {
	// 返回相应基本信息
	Response
	// 用户基本信息
	userJsonInfo
}

// 返回用户基本信息
/*
检查用户发送的id，检查其鉴权是否正确
*/
func GetUserInfo(c *gin.Context) {

	token := c.Query("token")
	
	// 判断用户是否已经登录
	if user, exist := userLoginInfos[token]; exist {
		c.JSON(http.StatusOK, interfaceResponse{
			Response: Response{
				Status_code: 0,
			},
			userJsonInfo: *user,
		})
	} else {
		c.JSON(http.StatusOK, Response{
			Status_code: -1,
			Status_msg: "Uesr don't exist",
		})
	}

}
