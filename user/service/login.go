// 账号登录
package service

import (
	"errors"

	"github.com/2103561941/douyin/user/repository"
)

// 登录用于与上层进行传输的
type UserLoginInfo struct {
	Username string
	Password string
	ID       int
}

// 判断能否登录， 用户名， 密码校验
func (user *UserLoginInfo) Login() error {

	// 构建数据库查询的记录
	record := &repository.UserInfo{
		Username: user.Username,
	}

	// 用户名不存在
	if err := record.QueryByUsername(); err != nil {
		return err
	}

	// 密码错误
	if record.Password != user.Password {
		return errors.New("wrong password")
	}

	// 装置ID
	user.ID = record.ID

	return nil
}
