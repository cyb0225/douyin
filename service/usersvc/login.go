// 用户登录

package usersvc

import (
	"errors"

	"github.com/2103561941/douyin/repository"
)

type UserLogin struct {
	Id       uint64
	Username string
	Password string
}


func (user *UserLogin) Login() error {

	record := &repository.User{
		Username: user.Username,
	}

	// username is not exist
	if err := record.SelectByUsername(); err != nil {
		return err
	}

	// password is wrong
	if record.Password != user.Password {
		return errors.New("wrong password")
	}

	// return user_id
	user.Id = record.Id

	return nil
}
