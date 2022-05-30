package videoctl

import (
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
)

type PublishListResponse struct {
	commonctl.Response
	Videos []*videosvc.VideoInfo `json:"voide_list"`
}

func GetPublishList(c *gin.Context) {
	token := c.Query("token")
	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	// query type transform
	authorInt, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user_id is undefined",
		})
		return
	}
	author := uint64(authorInt) //被访问的用户id

	userId := commonctl.UserLoginMap[token].Id // 主动去访问的用户id

	list := videosvc.PublishList{
		Author: userId,
		UserId: author,
	}

	if err := list.GetPublishList(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, PublishListResponse{
		Response: commonctl.Response{Status_code: 0},
		Videos:   list.Videos,
	})

}
