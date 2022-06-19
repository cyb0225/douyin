// get token and user_id, return basic userinfo

package userctl

import (
	"log"
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	commonctl.Response
	usersvc.UserInfo `json:"user"`
}

func GetUserInfo(c *gin.Context) {

	// type transfor : string -> uint64
	userIntId, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user_id input ",
		})
	}

	userId := uint64(userIntId)

	//token := c.Query("token")
	testcal, boolen := c.Get("middleware_geted_user_id")
	if boolen == false {
		log.Println("user_page didn't get")
	}

	// token is login, and call service to get user infos by user_id
	user := &usersvc.UserInfo{
		Id: userId,
	}

	//callerId := commonctl.UserLoginMap[token]
	callerId := testcal.(uint64)
	if err := user.SetUserInfo(callerId); err != nil { // read record error
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
