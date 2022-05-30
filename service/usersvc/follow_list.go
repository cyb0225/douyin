package usersvc

import (
	"github.com/2103561941/douyin/repository"
)

type FollowListResponse struct {
	UserId    uint64
	Followers []*UserInfo
}

// 拿到登陆用户id OK
// 通过id进行select
// 结果传回
func (list *FollowListResponse) FollowList() error {

	status := &repository.Follow{
		UserId: list.UserId,
	}

	records, err := status.GetFollowList()
	if err != nil {
		return err
	}

	// 创建一个临时存储list的变量，防止一半报错了
	tmpList :=  make([]*UserInfo, len(records))

	for i := 0; i < len(records); i++ {
		userInfo := &UserInfo{
			Id: records[i].ToUserId,
		}
		if err := userInfo.SetUserInfo(list.UserId); err != nil {
			return err
		}
		tmpList[i] = userInfo
	}

	list.Followers = tmpList

	return nil
}
