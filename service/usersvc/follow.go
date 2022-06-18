// 关注操作

package usersvc

import (
	"errors"

	"github.com/2103561941/douyin/repository"
)

type UserFollow struct {
	Id          uint64
	User_id     uint64
	To_user_id  uint64
	Action_type int
}


// 关注action
const (
	Not_Followed int = 0  
	Followed     int = 1  
	Undefined    int = -1
)

const (
	User_Want_to_Follow   int = 1
	User_Want_to_Unfollow int = 2
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
	}

	backstatus := status // 储存数据库操作前的state状态

	newStatus := follow.transformStatus(status)

	if newStatus == Undefined {
		return errors.New("action_type is undefined")
	}

	if err := status.UpdateStatus(newStatus); err != nil {
		return err
	}

	// updata user follow_count and user follower_count
	if err := follow.changeUsrFollowCount(); err != nil {
		if errr := status.Undo(backstatus); errr != nil { //如果changecount出现问题，进行数据回滚操作
			return errr
		}
		return err
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

// according to action_type, return status
func (follow *UserFollow) transformStatus(status *repository.Follow) int {
	switch follow.Action_type {

	case User_Want_to_Follow:
		if status.Status == Followed {
			return Undefined
		}
		if status.Status == Not_Followed {
			return Followed
		}
		return Undefined

	case User_Want_to_Unfollow:
		if status.Status == Followed {
			return Not_Followed
		}
		if status.Status == Not_Followed {
			return Undefined
		}
		return Undefined
	default:
		return Undefined
	}
}

// change user followcount and followercount
func (follow *UserFollow) changeUsrFollowCount() error {
	var n int
	if follow.Action_type == Followed {
		n = 1
	} else {
		n = -1
	}

	user := repository.User{
		Id: follow.User_id,
	}
	if err := user.UpdateFollowCount(n); err != nil {
		return err
	}

	to_user := repository.User{
		Id: follow.To_user_id,
	}
	if err := to_user.UpdateFollowerCount(n); err != nil {
		return err
	}

	return nil
}
