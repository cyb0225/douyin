package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type publishContent struct {
	Response
}

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if _, exist := userLoginInfos[token]; !exist {
		c.JSON(http.StatusOK, Response{Status_code: -1, Status_msg: "User didn't login"})
		return
	}
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			Status_code: -1,
			Status_msg:  "Invalid video",
		})
		return
	}
	title := c.PostForm("title")
	user := userLoginInfos[token]
	finalName := fmt.Sprintf("%d_%s", user.ID, title) // concat user id with file name prevent same name file
	saveFile := filepath.Join("./videoContent/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			Status_code: -1,
			Status_msg:  "Fail at saving file",
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		Status_code: 0,
		Status_msg:  finalName + " uploaded successfully",
	})
}
