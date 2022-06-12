// return uesr infos
package usersvc

import (
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

// set the userInfo response
func (user *UserInfo) SetUserInfo(id uint64) error {
	record := &repository.User{
		Id: user.Id,
	}

	if err := record.SelectByUserId(); err != nil {
		return err
	}

	user.Id = record.Id
	user.Name = record.Username
	user.FollowCount = record.FollowCount
	user.FollowerCount = record.FollowerCount
	user.IsFollow = user.getFollowStatus(id)
	user.Signature = "抖音青训营-源石技艺队-No.1107"
	user.Avatar = "https://github.com/sterben-01/sterben-01.github.io/blob/main/assets/blog_res/2022-06-10-tempimage.assets/68458315_p0.png"
	user.Background_img = "https://github.com/sterben-01/sterben-01.github.io/blob/main/assets/blog_res/2022-06-10-tempimage.assets/61163969_p0.jpg"

	return nil
}
