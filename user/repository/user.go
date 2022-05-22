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

	//插入失败
	if err := db.Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase error")
	}

	return nil
}

// 通过用户名查询记录， 查询到后将记录保存到调用的变量里，查询不到返回err
func (user *UserInfo) QueryByUsername() error {

	// 查询不到记录
	result := db.Where("username = ?", user.Username).First(user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New("record not found")
	}

	return nil
}


