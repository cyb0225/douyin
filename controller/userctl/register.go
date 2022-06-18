package userctl

import (
	"github.com/2103561941/douyin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
)

type registerResponse struct {
	commonctl.Response
	Id    uint64 `json:"user_id"`
	Token string `json:"token"`
}

//
func Register(c *gin.Context) {

	user := usersvc.UserRegister{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}
	
	token, err := middleware.SetUpToken(user.Username)
	if err != nil {
		c.Abort()
	}
	if err := user.Register(); err != nil { // register wrong
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else { // register success
		commonctl.UserLoginMap[token] = commonctl.UserLoginComp{Id: user.Id}
		c.JSON(http.StatusOK, registerResponse{
			Response: commonctl.Response{Status_code: 0},
			Id:       user.Id,
			Token:    token,
		})
	}
}
