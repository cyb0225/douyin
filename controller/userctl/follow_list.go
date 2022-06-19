package userctl

import (
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

type FollowListResponse struct {
	commonctl.Response
	Followers []*usersvc.UserInfo `json:"user_list"`
}

func FollowList(c *gin.Context) {

	// query type transform
	userIdInt, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user_id is undefined",
		})
		return
	}
	userId := uint64(userIdInt) //对象ID

	inputData := &usersvc.FollowListResponse{
		UserId: userId,
	}

	if err := inputData.FollowList(); err != nil {
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
