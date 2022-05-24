// 用户登录
package controller

import (
	"net/http"

	"github.com/2103561941/douyin/user/service"
	"github.com/gin-gonic/gin"
)

// 登录返回响应
type loginResponse struct {
	// 返回相应基本信息
	Response
	// 用户id
	User_id int `json:"user_id"`
	// 用户鉴权
	Token string `json:"token"`
}

// 用户登录的接口
/*
接收用户发送的用户名和密码，判断用户名是否存在，密码是否正确
根据是否成功登录，返回不同的状态码
创建token，用于进入用户页面后的鉴权
*/
func Login(c *gin.Context) {

	// 读取用户登录数据
	user := service.UserLoginInfo{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	token := user.Username + user.Password

	// 登录并返回登录响应
	if err := user.Login(); err != nil { // 登录失败
		c.JSON(http.StatusOK, Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else { // 登录成功
		userLoginInfos[token] = &service.UserJsonInfo{
			ID:            user.ID,
			Username:      user.Username,
			FollowCount:   0,
			FollowerCount: 100000000000000,
		}
		c.JSON(http.StatusOK, loginResponse{
			Response: Response{Status_code: 0},
			User_id:  user.ID,
			Token:    token,
		})
	}

}
