// user login
// get username and password
// send to user service to deal with
// return stauts and some messages
package userctl

import (
	"github.com/2103561941/douyin/middleware"
	"net/http"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

// json struct to send back
type loginResponse struct {
	commonctl.Response
	Id    uint64 `json:"user_id"`
	Token string `json:"token"`
}

func Login(c *gin.Context) {

	user := usersvc.UserLogin{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	//token := commonctl.CreatToken(user.Username, user.Password)
	user.Password = commonctl.MD5(user.Password)
	if err := user.Login(); err != nil { // 登录失败
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else {
		token, err := middleware.SetUpToken(user.Id)
		if err != nil {
			c.Abort()
		}
		commonctl.UserLoginMap[token] = commonctl.UserLoginComp{Id: user.Id}
		c.JSON(http.StatusOK, loginResponse{
			Response: commonctl.Response{Status_code: 0},
			Id:       user.Id,
			Token:    token,
		})
	}
}
