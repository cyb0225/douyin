package usersvc

import (
	"github.com/2103561941/douyin/repository"
)

type UserFollow struct {
	Id          uint64
	User_id     uint64
	To_user_id  uint64
	Action_type uint64
}

const (
	Not_Followed uint64 = 0
	Followed     uint64 = 1
)

const (
	User_Want_to_Follow   uint64 = 1
	User_Want_to_Unfollow uint64 = 2
)

func (follow *UserFollow) Follow() error {

	// user or to_user is not exist
	if err := follow.checkUserExist(); err != nil {
		return err
	}

	status := &repository.Follow{
		UserId:   follow.User_id,
		ToUserId: follow.To_user_id,
		Status:   follow.Action_type,
	}

	// record not found
	if err := status.Select(); err != nil {

		if err := status.Insert(); err != nil {
			return err
		}

	} else {

		if err := status.UpdateStatus(); err != nil {
			
		}

	}

	return nil
}




func (follow *UserFollow) checkUserExist() error {
	user := &repository.User{
		Id: follow.User_id,
	}

	if err := user.SelectByUserId(); err != nil {
		return err
	}

	toUser := &repository.User{
		Id: follow.To_user_id,
	}

	if err := toUser.SelectByUserId(); err != nil {
		return err
	}

	return nil
}
