package usersvc

import (
	"errors"

	"github.com/2103561941/douyin/repository"
)

type UserFollow struct {
	Id          string
	User_id     string
	To_user_id  string
	Action_type string
}

const (
	Not_Followed       string = "0"
	Followed           string = "1"
	Followed_by_Others string = "2"
	Mutual_Followed    string = "3"
)

const (
	User_Want_to_Follow   string = "1"
	User_Want_to_Unfollow string = "2"
)

func (user *UserFollow) Follow() error {
	status := &repository.Follow{
		UserId:   user.User_id,
		ToUserId: user.To_user_id,
		Status:   user.Action_type,
	}

	if err := status.CheckStatus(); err != nil {
		//报错应该是因为没有找到。所以建立对应行 不知道是这个意思不
		/*
			INSERT INTO follow(user_id, follower_id, status)
			VALUES
			(user.User_id, user.To_user_id, user.Action_type);
		*/

		if err := status.Insert(); err != nil {
			return err
		}

	}

	switch status.Status {
	case Not_Followed:
		// not followed
		/*
			UPDATE follow
			SET status = "1"
			WHERE User_id = user.User_id;
		*/
		if user.Action_type == User_Want_to_Follow {
			if err := status.UpdateStatus(Followed); err != nil {
				return err
			}
		}
		if user.Action_type == User_Want_to_Unfollow {
			return errors.New("you didn't follow this user")
		}

	case Followed:
		// already followed this user. 已经关注该用户
		if user.Action_type == User_Want_to_Follow {
			return errors.New("you already followed this user")
		}
		if user.Action_type == User_Want_to_Unfollow {
			if err := status.UpdateStatus(Not_Followed); err != nil {
				return err
			}
		}
	case Followed_by_Others:
		// already followed by this user 已经被该用户关注
		/*
			UPDATE follow
			SET status = "3"
			WHERE User_id = user.User_id;
		*/
		if user.Action_type == User_Want_to_Follow {
			if err := status.UpdateStatus(Mutual_Followed); err != nil {
				return err
			}
		}

		if user.Action_type == User_Want_to_Unfollow {
			return errors.New("you didn't follow this user")
		}

	case Mutual_Followed:
		// already mutual following 已经互粉
		if user.Action_type == User_Want_to_Follow {
			return errors.New("you already mutual following with this user")
		}
		if user.Action_type == User_Want_to_Unfollow {
			if err := status.UpdateStatus(Followed_by_Others); err != nil {
				return err
			}
		}
	default:
	}

	return nil
}
