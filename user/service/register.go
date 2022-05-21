// 注册
package server

import "log"

// 保存账号密码信息，用于和上层进行数据交换
type UserRegInfo struct {
	Username string
	Password string
}

// 注册, 判断账号密码有效性，加入仓库
func Register(user *UserRegInfo) error {

	// 使用错误链返回
	if err := user.checkUsername(); err != nil {
		return err
	}

	if err := user.checkPassword(); err != nil {
		return err
	}

	// 加入数据库仓库
	log.Println("into server")
	return nil
}

//-----------------------------------------------------------------

// 账号有效性判断
func (msg *UserRegInfo) checkUsername() error {

	return nil
}

// 密码有效性判断
func (msg *UserRegInfo) checkPassword() error {

	return nil
}
