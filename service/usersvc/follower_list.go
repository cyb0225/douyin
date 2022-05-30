package usersvc

import (
	"github.com/2103561941/douyin/repository"
)

type FollowerListResponse struct {
	ToUserId  uint64
	Followers []*UserInfo
}

func (list *FollowerListResponse) FollowerList() error {

	status := &repository.Follow{
		ToUserId: list.ToUserId,
	}

	records, err := status.GetFollowerList()
	if err != nil {
		return err
	}

	// 创建一个临时存储list的变量，防止一半报错了
	tmpList := make([]*UserInfo, len(records))

	for i := 0; i < len(records); i++ {
		userInfo := &UserInfo{
			Id: records[i].UserId,
		}
		if err := userInfo.SetUserInfo(list.ToUserId); err != nil {
			return err
		}
		tmpList[i] = userInfo
	}

	list.Followers = tmpList

	return nil
}
