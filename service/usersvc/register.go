// 用户注册

package usersvc

import (
	"errors"
	"github.com/2103561941/douyin/controller/commonctl"
	"github.com/2103561941/douyin/repository"
)

type UserRegister struct {
	Id       uint64
	Username string
	Password string
}

// register, store into repository when success
func (user *UserRegister) Register() error {

	// Invalid username
	if err := user.checkUsername(); err != nil {
		return err
	}

	// Invalid password (too long or too short)
	if err := user.checkPassword(); err != nil {
		return err
	}

	// creat object to be  insert
	record := &repository.User{
		Username: user.Username,
		Password: user.encodePassword(),
	}

	// duplicate username
	if err := record.SelectByUsername(); err == nil {
		return errors.New("username already exists")
	}

	// insert record into repositroy
	if err := record.Insert(); err != nil {
		return err
	}

	// set user_id to return by object
	user.Id = record.Id

	return nil
}

//-----------------------------------------------------------------

// Determine the validity of the username, return error when username is longger than 32
func (user *UserRegister) checkUsername() error {
	if len(user.Username) > 32 {
		return errors.New("username is greater than 32")
	}

	return nil
}

// Determine the validity of the password,
// return error when password is longger than 32 or when is short than 5
func (user *UserRegister) checkPassword() error {
	if len(user.Password) <= 5 {
		return errors.New("password length is less than or equal to 5")
	}

	if len(user.Password) > 32 {
		return errors.New("password length is greater than 32")
	}

	return nil
}

// password encoding
func (user *UserRegister) encodePassword() string {
	// encrypted password
	enPassword := commonctl.MD5(user.Password)
	return enPassword
}
