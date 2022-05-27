package userctl

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/2103561941/douyin/service/usersvc"
	"github.com/2103561941/douyin/controller/commonctl"
)

type registerResponse struct {
	commonctl.Response
	ID    uint64 `json:"user_id"`
	Token string `json:"token"`
}

//
func Register(c *gin.Context) {

	user := usersvc.UserRegisterInfo{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	token := user.Username + user.Password

	if err := user.Register(); err != nil { // register wrong
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else { // register success
		commonctl.UserLoginMap[token] = struct{}{}
		c.JSON(http.StatusOK, registerResponse{
			Response: commonctl.Response{Status_code: 0},
			ID:       user.ID,
			Token:    token,
		})
	}
}
