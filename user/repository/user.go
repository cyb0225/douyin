// 存储用户和密码的数据，以及相关的增删改查效果

package repository

import (
	"errors"

	"gorm.io/gorm"
)

var (
	// 与数据库进行连接的变量
	db *gorm.DB
)

// 用户信息数据
type UserInfo struct {
	ID       int    `gorm:"primaryKey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
}

// 用户数据插入, 只做插入操作（查重由调用的函数进行选择）
func (user *UserInfo) Insert() error {

	//插入失败就报错返回
	if err := db.Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase error")
	}

	return nil
}

// 用户数据查询(通过用户名), 返回查询到的记录数据, 若查询不到则返回err
func (user *UserInfo) QueryByUsername() error {

	return nil
}
