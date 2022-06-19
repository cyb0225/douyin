package videoctl

import (
	"log"
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
)

type PublishListResponse struct {
	commonctl.Response
	Videos []*videosvc.VideoInfo `json:"video_list"`
}

func GetPublishList(c *gin.Context) {

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

	testcal, boolen := c.Get("middleware_geted_user_id")
	if boolen == false {
		log.Println("user_page didn't get")
	}

	userId := testcal.(uint64)

	list := videosvc.PublishList{
		Author: author,
		UserId: userId,
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
