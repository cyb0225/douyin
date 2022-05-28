package usersvc

import (
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/repository"
)

type FollowListResponse struct {
	commonctl.Response
	user []repository.User `json:"user_list"`
}

func 
