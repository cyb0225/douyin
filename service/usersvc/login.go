// login
// check if username exists
// check if password is right
// return user_id
package usersvc

import (
	"errors"

	"github.com/2103561941/douyin/repository"
)

type UserLoginInfo struct {
	ID       uint64
	Username string
	Password string
}

// login
func (user *UserLoginInfo) Login() error {

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
	user.ID = record.ID

	return nil
}
