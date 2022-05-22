package controller

import (
	"log"

	"github.com/2103561941/douyin/user/service"

	"github.com/gin-gonic/gin"
)

// 注册接口的返回相应
type regResponse struct {
	// 返回状态码 0 表示成功， 其他表示失败
	Status_code int `json:"status_code"`
	// 返回状态描述
	Status_msg string `json:"status_msg"`
	// 用户id
	User_id int `json:"user_id"`
	// 用户鉴权
	Token string `json:"token"`
}

// 注册的请求接口
/*
主要流程就是通过service层的结构体(UserRegInfo)接收请求，然后传给service层进行逻辑处理
通过service层逻辑处理返回的结果，通过返回响应(regResponse)这个类进行返回
*/


func Register(c *gin.Context) {

	// 用于接收json的post请求
	newUser := service.UserRegInfo{}

	// 接收json请求
	c.BindJSON(&newUser)

	// 注册账号
	err := newUser.Register()

	// 创建返回响应
	reR := regResponse{}

	// 根据是否注册成功返回结果
	if err == nil {
		reR.Status_code = 0
		reR.Status_msg = "success"
		reR.User_id = newUser.ID
		log.Println("register successfully")
	} else {
		reR.Status_code = -1
		reR.Status_msg = err.Error()
		log.Println(err.Error())
	}

	// 鉴权
	reR.Token = newUser.Username + newUser.Password

	c.JSON(200, reR)
}
