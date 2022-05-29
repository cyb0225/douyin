package userctl

//import (
//	"net/http"
//	"strconv"
//
//	"github.com/2103561941/douyin/controller/commonctl"
//	"github.com/2103561941/douyin/service/usersvc"
//	"github.com/gin-gonic/gin"
//)
//
//func FollowList(c *gin.Context) {
//
//	userIdInt, err := strconv.Atoi(c.Query("user_id"))
//	if err != nil {
//		c.JSON(http.StatusOK, commonctl.Response{
//			Status_code: -1,
//			Status_msg:  "user_id is undefined",
//		})
//		return
//	}
//
//	userId := uint64(userIdInt)
//
//	token := c.Query("token")
//
//	if _, ok := commonctl.UserLoginMap[token]; !ok {
//		c.JSON(http.StatusOK, commonctl.Response{
//			Status_code: -1,
//			Status_msg:  "user is not login",
//		})
//		return
//	}
//
//	follow_list := usersvc.FollowListResponse{
//
//	}
//}
