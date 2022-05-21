package controller

import (
	"github.com/2103561941/douyin/user/service"

	"github.com/gin-gonic/gin"
)

// 注册接口的返回相应
type regResponse struct {
	// 返回状态码 0 表示成功， 其他表示失败
	status_code int `json:"status_code"`
	// 返回状态描述
	status_msg string `json:"status_msg"`
	// 用户id
	user_id int
	// 用户鉴权
	// token string
}

func Register(c *gin.Context) {
	username := c.Query("username")
	password := c.Query("password")

	newUser := service.UserRegInfo{
		Username: username,
		Password: password,
	}

	err := service.Register(&newUser)

	reR := regResponse{}

	// 根据是否注册成功返回结果
	if err == nil {
		reR.status_code = 0
		reR.status_msg = "success"
	} else {
		reR.status_code = -1
		reR.status_msg = err.Error()
	}

	c.JSON(200, reR)

}
