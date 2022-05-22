package controller

import (
	"net/http"

	"github.com/2103561941/douyin/user/service"

	"github.com/gin-gonic/gin"
)

// 注册接口的返回相应
type regResponse struct {
	// 返回相应基本信息
	Response
	// 用户id
	User_id int `json:"user_id"`
	// 用户鉴权
	Token string `json:"token"`
}

// 注册的请求接口
/*
获取用户发送的用户名和密码，使用service层进行逻辑处理，
针对注册是否有效返回不同的json报文
*/

func Register(c *gin.Context) {

	newUser := service.UserRegInfo{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	// 注册账号并返回json响应
	if err := newUser.Register(); err != nil { // 注册失败
		c.JSON(http.StatusOK, Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else { // 注册成功
		c.JSON(http.StatusOK, regResponse{
			Response: Response{Status_code: 0},
			User_id:  newUser.ID,
			Token:    newUser.Username + newUser.Password,
		})
	}

}
