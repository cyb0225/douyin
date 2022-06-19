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
	// 鉴权
	token := c.PostForm("token")
	//if _, ok := commonctl.UserLoginMap[token]; !ok {
	//	c.JSON(http.StatusOK, commonctl.Response{
	//		Status_code: -1,
	//		Status_msg:  "user is not login",
	//	})
	//	return
	//}

	// 获取视频数据
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	// 生成视频名
	videoFileName := filepath.Base(data.Filename)
	user := commonctl.UserLoginMap[token]
	title := c.PostForm("title")

	//文件名后加unix时间戳
	finalVideoName := fmt.Sprintf("%d_%d_%s", user.Id, time.Now().Unix(), videoFileName)

	// 将数据放入到OS对象存储中
	playUrl, err := OS.PutVideoObject(finalVideoName, data)
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	// 获取保存封面的url
	coverUrl, err := GetCover(finalVideoName, playUrl)
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
		return
	}

	videoinfo := &videosvc.PublishVideo{
		UserID:   user.Id,
		PlayURL:  playUrl,
		CoverURL: coverUrl,
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
