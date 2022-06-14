// 用户信息

package repository

import (
	"errors"

	"gorm.io/gorm"
)


type User struct {
	Id            uint64 `gorm:"column:id"`
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   uint64 `gorm:"follow_count"`
	FollowerCount uint64 `gorm:"follower_count"`
}

func (*User) TableName() string {
	return "user"
}

// 插入数据
func (user *User) Insert() error {

	//insert error
	if err := Db.Table(user.TableName()).Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase error")
	}

	return nil
}

// select user record by username
func (user *User) SelectByUsername() error {

	result := Db.Table(user.TableName()).Where("username = ?", user.Username).First(user)

	// not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// select user record by user_id
func (user *User) SelectByUserId() error {

	result := Db.Table(user.TableName()).Where("id = ?", user.Id).First(user)

	// not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// updata user follow_count by add n
func (user *User) UpdataFollowCount(n int) error {
	result := Db.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follow_count", int(user.FollowCount)+n)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}

// updata user follower_count by add n
func (user *User) UpdataFollowerCount(n int) error {
	result := Db.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follower_count", int(user.FollowerCount)+n)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return result.Error
	}

	return nil
}
