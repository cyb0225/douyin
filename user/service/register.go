// 注册
package service

import (
	"errors"

	"github.com/2103561941/douyin/user/repository"
)

// 保存账号密码信息，用于和上层进行数据交换
type UserRegInfo struct {
	Username string
	Password string
	ID       int
}

// 注册, 判断账号密码有效性，加入仓库
func Register(user *UserRegInfo) error {

	// 使用错误链返回
	// 检测用户名有效性
	if err := user.checkUsername(); err != nil {
		return err
	}

	// 检测密码有效性
	if err := user.checkPassword(); err != nil {
		return err
	}

	// 创建待插入的记录
	inputUser := &repository.UserInfo{
		Username: user.Username,
		Password: user.Password,
	}

	// 用户名查重判断
	if err := inputUser.QueryByUsername(); err == nil {
		return errors.New("username already exists")
	}

	// 将新注册的用户名和密码插入数据库的user表
	if err := inputUser.Insert(); err != nil {
		return nil
	}

	// 配置用户ID（数据库给的主键）
	user.ID = inputUser.ID

	return nil
}

//-----------------------------------------------------------------

// 用户名有效性判断, 如果用户名字符长度大于32返回错误
func (msg *UserRegInfo) checkUsername() error {
	if len(msg.Username) > 32 {
		return errors.New("username is greater than 32")
	}

	return nil
}

// 密码有效性判断，如果密码长度小于等于5或者大于32返回错误
func (msg *UserRegInfo) checkPassword() error {
	if len(msg.Password) <= 5 {
		return errors.New("password length is less than or equal to 5")
	}

	if len(msg.Password) > 32 {
		return errors.New("password length is greater than 32")
	}

	return nil
}
