package videoctl

import (
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
)

type likeresponse struct {
	commonctl.Response
}

type rawlikedata struct {
	UserID     uint64
	videoID    string
	actiontype string
}

func Like(c *gin.Context) {
	Token := c.Query("token")
	//user not login
	if _, ok := commonctl.UserLoginMap[Token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}
	inputdata := rawlikedata{
		UserID:     commonctl.UserLoginMap[Token].Id,
		videoID:    c.Query("video_id"),
		actiontype: c.Query("action_type"),
	}

	user, err := inputdata.converter()
	user.UserId = commonctl.UserLoginMap[Token].Id // 主动去访问的用户id

	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "transform fail",
		})
		return
	}

	if err := user.Like(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, likeresponse{
			Response: commonctl.Response{Status_code: 0},
		})
	}
}

func (data *rawlikedata) converter() (*videosvc.Like, error) {

	videoID, err := strconv.Atoi(data.videoID)
	if err != nil {
		return nil, err
	}

	actiontype, err := strconv.Atoi(data.actiontype)
	if err != nil {
		return nil, err
	}

	user := &videosvc.Like{
		VideoId:    uint64(videoID),
		ActionType: actiontype,
	}

	return user, nil
}
