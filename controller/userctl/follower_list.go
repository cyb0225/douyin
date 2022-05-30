package userctl

import (
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

type FollowerListResponse struct {
	commonctl.Response
	Followers []*usersvc.UserInfo `json:"user_list"`
}

func FollowerList(c *gin.Context) {

	// query type transform
	toUserIdInt, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user_id is undefined",
		})
		return
	}
	toUserId := uint64(toUserIdInt) //对象ID

	// Determine if the user is logged in
	token := c.Query("token")

	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	inputData := &usersvc.FollowerListResponse{
		ToUserId: toUserId,
	}

	// get follower_list
	if err := inputData.FollowerList(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	c.JSON(http.StatusOK, FollowListResponse{
		Response:  commonctl.Response{Status_code: 0},
		Followers: inputData.Followers,
	})
}
