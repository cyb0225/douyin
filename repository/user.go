// store userinfos and related CRUD interface

package repository

import (
	"errors"

	"gorm.io/gorm"
)

// user mesages in database table
type User struct {
	Id            uint64 `gorm:"column:id"` //5.28 update typo
	Username      string `gorm:"column:username"`
	Password      string `gorm:"column:password"`
	FollowCount   uint64 `gorm:"follow_count"`
	FollowerCount uint64 `gorm:"follower_count"`
}

func (*User) TableName() string {
	return "user"
}

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
