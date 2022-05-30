package videoctl

import (
	"net/http"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/gin-gonic/gin"
)

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	if _, ok := commonctl.UserLoginMap[token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg: "post file errro",
		})
	}

}