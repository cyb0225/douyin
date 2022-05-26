// user login
// get username and password
// send to user service to deal with
// return stauts and some messages
package userctl

import (
	"net/http"

	"github.com/2103561941/douyin/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

// json struct to send back
type loginResponse struct {
	commonctl.Response
	ID int `json:"user_id"`
	Token string `json:"token"`
}



func Login(c *gin.Context) {

	// 读取用户登录数据
	user := usersvc.UserLoginInfo{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	token := user.Username + user.Password

	// 登录并返回登录响应
	if err := user.Login(); err != nil { // 登录失败
		c.JSON(http.StatusOK,commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else { // 登录成功
		userLoginInfos[token] = &usersvc.UserJsonInfo{
			ID:            user.ID,
			Username:      user.Username,
			FollowCount:   0,
			FollowerCount: 100000000000000,
		}
		c.JSON(http.StatusOK, loginResponse{
			Response: commonctl.Response{Status_code: 0},
			ID:  user.ID,
			Token:    token,
		})
	}

}
