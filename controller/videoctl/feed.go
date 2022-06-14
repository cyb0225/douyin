package videoctl

import (
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
	"net/http"
)

type FeedResponse struct {
	commonctl.Response
	Next_time string
	Videos    []*videosvc.VideoInfo `json:"video_list"`
}

func Feed(c *gin.Context) {
	token := c.Query("token")

	println(c.Query("latest_time"))
	userId := commonctl.UserLoginMap[token].Id // 主动去访问的用户id
	list := videosvc.Feedliststruct{
		Latest_time: c.Query("latest_time"),
		UserID:      userId,
	}
	if err := list.Feedlist(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})

	}
	c.JSON(http.StatusOK, FeedResponse{
		Response:  commonctl.Response{Status_code: 0},
		Next_time: list.Earlist_video,
		Videos:    list.Videos,
	})

}
