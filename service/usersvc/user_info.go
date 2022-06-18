// 用户基本信息
package usersvc

import (
	"errors"
	"github.com/2103561941/douyin/repository"
)

type UserInfo struct {
	Id             uint64 `json:"id"`
	Name           string `json:"name"`
	FollowCount    uint64 `json:"follow_count"`
	FollowerCount  uint64 `json:"follower_count"`
	IsFollow       bool   `json:"is_follow"`
	Signature      string `json:"signature"`
	Avatar         string `json:"avatar"`
	Background_img string `json:"background_image"`
}

// 获取用户关注状态
func (user *UserInfo) getFollowStatus(id uint64) bool {
	status := &repository.Follow{
		UserId:   id,      //token对应的
		ToUserId: user.Id, //传入ID
	}
	if ret := status.Select(); ret != nil {
		return false
	} else {
		if status.Status == 1 {
			return true
		} else {
			return false
		}
	}
}

// 设置用户基本信息
func (user *UserInfo) SetUserInfo(id uint64) error {
	record := &repository.User{
		Id: user.Id,
	}

	if err := record.SelectByUserId(); err != nil {
		return err
	}

	if err := record.RewriteToRedis(); err != nil { //数据写回
		return errors.New("REDIS --- Rewrite setuserinfo error")
	}

	user.Id = record.Id
	user.Name = record.Username
	user.FollowCount = record.FollowCount
	user.FollowerCount = record.FollowerCount
	user.IsFollow = user.getFollowStatus(id)
	// 用户页面基本信息（头像、签名、背景），由于这方面客户端没给修改的接口，我们直接选择默认
	user.Signature = "抖音青训营-源石技艺队-No.1107"
	user.Avatar = "https://sterben-01.github.io/assets/blog_res/2022-06-11-imgtemp.assets/68214616_p0.png"
	user.Background_img = "https://sterben-01.github.io/assets/blog_res/2022-06-11-imgtemp.assets/61163969_p0.jpg"

	return nil
}
