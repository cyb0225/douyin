package videoctl

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/2103561941/douyin/controller/commonctl"
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
			Status_code: 1,
			Status_msg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	user := commonctl.UserLoginMap[token]
	finalName := fmt.Sprintf("%d_%s", user.Id, filename)
	saveFile := filepath.Join("./video_content/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: 1,
			Status_msg:  err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, commonctl.Response{Status_code: 0})
}
