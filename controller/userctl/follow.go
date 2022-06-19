package userctl

import (
	"log"
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

func Follow(c *gin.Context) {

	inputData := inputStringData{
		To_user_id:  c.Query("to_user_id"),
		Action_type: c.Query("action_type"),
	}

	// transform string to int
	user, err := inputData.transfromToFollow()
	if err != nil {
		c.JSON(http.StatusOK, commonctl.Response{
			Status_code: -1,
			Status_msg:  "transform fail",
		})
		return
	}
	testcal, boolen := c.Get("middleware_geted_user_id")
	if boolen == false {
		log.Println("user_page didn't get")
	}

	user.User_id = testcal.(uint64)
	log.Println(user.User_id)

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

	to_user_id, err := strconv.Atoi(data.To_user_id)
	if err != nil {
		return nil, err
	}

	action_type, err := strconv.Atoi(data.Action_type)
	if err != nil {
		return nil, err
	}

	user := &usersvc.UserFollow{
		To_user_id:  uint64(to_user_id),
		Action_type: action_type,
	}

	return user, nil
}
