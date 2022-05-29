package usersvc

import (
	"github.com/2103561941/douyin/repository"
)

type CheckFollowList struct {
	Id   uint64
	user []repository.User
}

// 拿到登陆用户id OK
// 通过id进行select
// 结果传回
func (list *FollowerListResponse) FollowList(id *CheckFollowList) error {
	//利用select
	/*
		SELECT * from follow
		WHERE
		user_id = id,
		status = 1
	*/
	status := &repository.Follow{
		UserId: id.Id,
	}
	status.GetFollowList()

}
