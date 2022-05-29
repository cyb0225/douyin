package userctl

import (
	"net/http"
	"strconv"

	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/service/usersvc"
	"github.com/gin-gonic/gin"
)

// json struct to send back
type inputStringData struct {
	User_id     string
	To_user_id  string
	Action_type string
}

type followResponse struct {
	commonctl.Response //response struct
}

func Follow(c *gin.Context) {

	inputData := inputStringData{
		User_id:     c.Query(("user_id")),
		To_user_id:  c.Query("to_user_id"),
		Action_type: c.Query("action_type"),
	}

	// transform string to int
	user, err := inputData.transfromToFollow()
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	// user is not login or register
	Token := c.Query("token")
	if _, ok := commonctl.UserLoginMap[Token]; !ok {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "user is not login",
		})
		return
	}

	if err := user.Follow(); err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, loginResponse{
			Response: commonctl.Response{Status_code: 0},
		})
	}

}

func (data *inputStringData) transfromToFollow() (*usersvc.UserFollow, error) {
	user_id, err := strconv.Atoi(data.User_id)
	if err != nil {
		return nil, err
	}

	to_user_id, err := strconv.Atoi(data.To_user_id)
	if err != nil {
		return nil, err
	}

	action_type, err := strconv.Atoi(data.Action_type)
	if err != nil {
		return nil, err
	}

	user := &usersvc.UserFollow{
		User_id:     uint64(user_id),
		To_user_id:  uint64(to_user_id),
		Action_type: action_type,
	}

	return user, nil
}
