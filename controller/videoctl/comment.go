package videoctl

import (
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type commentresponse struct {
	commonctl.Response
	//下面的结构体不会写
}

type rawcommentdata struct {
	ToUserID     string // 视频作者ID
	videoID      string
	actiontype   string
	comment_text string
	comment_id   string
}

func Comment(c *gin.Context) {
	Token := c.Query("token")
	if _, ok := commonctl.UserLoginMap[Token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}
	inputdata := rawcommentdata{
		ToUserID:     c.Query("user_id"),
		videoID:      c.Query("video_id"),
		actiontype:   c.Query("action_type"),
		comment_text: c.Query("comment_text"),
		comment_id:   c.Query("comment_id"),
	}
	//------------需要调试的部分-----------------

	user, err := inputdata.converter()
	user.UserId = commonctl.UserLoginMap[Token].Id // 主动去访问的用户id
	println(commonctl.UserLoginMap[Token].Id)
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "transform fail",
		})
		return
	}
	//------------end-----------------
	if err := user.Comment(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, commentresponse{
			Response: commonctl.Response{Status_code: 0},
		})
	}

}

func (data *rawcommentdata) converter() (*videosvc.Comment, error) {
	to_user_id, err := strconv.Atoi(data.ToUserID)
	if err != nil {
		return nil, err
	}
	videoID, err := strconv.Atoi(data.videoID)
	if err != nil {
		return nil, err
	}
	actiontype, err := strconv.Atoi(data.actiontype)
	if err != nil {
		return nil, err
	}
	CommentID, err := strconv.Atoi(data.comment_id)
	if err != nil {
		println("correct")
	}

	user := &videosvc.Comment{
		ToUserID:    uint64(to_user_id),
		VideoId:     uint64(videoID),
		ActionType:  actiontype,
		CommentID:   uint64(CommentID),
		CommentText: data.comment_text,
	}

	return user, nil
}
