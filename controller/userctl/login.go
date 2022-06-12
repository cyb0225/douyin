// user login
// get username and password
// send to user service to deal with
// return stauts and some messages
package userctl

import (
	"net/http"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

// json struct to send back
type loginResponse struct {
	commonctl.Response
	Id    uint64 `json:"user_id"`
	Token string `json:"token"`
}

func Login(c *gin.Context) {

	user := usersvc.UserLogin{
		Username: c.Query("username"),
		Password: c.Query("password"),
	}

	token := commonctl.CreatToken(user.Username, user.Password)
	user.Password = commonctl.MD5(user.Password)
	if err := user.Login(); err != nil { // 登录失败
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else {
		commonctl.UserLoginMap[token] = commonctl.UserLoginComp{Id: user.Id}
		c.JSON(http.StatusOK, loginResponse{
			Response: commonctl.Response{Status_code: 0},
			Id:       user.Id,
			Token:    token,
		})
	}

}

//func MD5(str string) string {
//	data := []byte(str) //切片
//	has := md5.Sum(data)
//	md5str := fmt.Sprintf("%x", has) //将[]byte转成16进制
//	return md5str
//}
