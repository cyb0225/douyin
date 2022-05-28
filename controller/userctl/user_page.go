// get token and user_id, return basic userinfo

package userctl

import (
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	commonctl.Response
	usersvc.UserInfo
}

func UserInfo(c *gin.Context) {

	// type transfor : string -> uint64
	userIntId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user_id input ",
		})
	}

	userId := uint64(userIntId)

	token := c.Query("token")

	// this token is not login before
	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	// token is login, and call service to get user infos by user_id
	user := &usersvc.UserInfo{
		Id: userId,
	}
	
	if err := user.SetUserInfo(); err != nil { // read record error
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})

	} else {
		c.JSON(http.StatusOK, UserResponse{
			Response: commonctl.Response{Status_code: 0},
			UserInfo: *user,
		})
	}
}
