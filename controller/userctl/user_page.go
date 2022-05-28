// get token and user_id, return basic userinfo

package userctl

import (
	"net/http"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)



func UserInfo(c *gin.Context) {

	userId := c.Query("user_id")
	
	token := c.Query("token")

	// this token is not login before
	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg: "user is not login",
		})
	}

	// token is login, and call service to get user infos by user_id
	user := &usersvc.UserResponse{
		
	}
	if err :=  user.SetUserInfo(); err != nil{
		c.JSON(http.StatusOK, usersvc.UserResponse{
			
		})

	} else {

	}
}
