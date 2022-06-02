package videoctl

import (
	"fmt"
	"net/http"
	"path/filepath"
	"time"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/videosvc"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	// 我们这块就不用前面的ip和端口了，到时候客户端申请url的时候我们把代理的url加上返回给他就可以
	localhost := "/douyin/index"

	videoFileName := filepath.Base(data.Filename)
	user := commonctl.UserLoginMap[token]
	title := c.PostForm("title")

	finalVideoName := fmt.Sprintf("%d_%d_%s", user.Id, time.Now().Unix(), videoFileName)
	//需要判断同一用户上传同一个文件两次的情况。已修改。文件名后加unix时间戳
	saveFile := filepath.Join("./video_content/", finalVideoName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	finalCoverName, err := GetCover(finalVideoName)
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	videoinfo := &videosvc.PublishVideo{
		UserID:   user.Id,
		PlayURL:  localhost + "/video/" + finalVideoName,
		CoverURL: localhost + "/cover/" + finalCoverName,
		Title:    title,
	}

	if err := videoinfo.PublishVideo(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonctl.Response{Status_code: 0})
}
