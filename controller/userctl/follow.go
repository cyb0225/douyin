package userctl

import (
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
	"net/http"
)

// json struct to send back

type followResponse struct {
	commonctl.Response //response struct
}

func Follow(c *gin.Context) {

	user := usersvc.UserFollow{
		User_id:     c.Query(("user_id")),
		To_user_id:  c.Query("to_user_id"),
		Action_type: c.Query("action_type"),
	}

	Token := c.Query("token")
	if _, ok := commonctl.UserLoginMap[Token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
	}

	if err := user.Follow(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})

	} else {
		//commonctl.UserLoginMap[Token] = struct{}{} 不知道加不加这一行
		c.JSON(http.StatusOK, loginResponse{
			Response: commonctl.Response{Status_code: 0},
		})
	}

}
