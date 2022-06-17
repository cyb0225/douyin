// store userinfos and related CRUD interface

package repository

import (
	"errors"
	"sync"

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
	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	//insert error
	if err := tx.Table(user.TableName()).Create(&user).Error; err != nil {
		tx.Rollback()
		return errors.New("Insert to UserDatabase error, roll backed")
	}
	tx.Commit()
	mutex.Unlock()
	//
	//err := Db.Transaction(func(tx *gorm.DB) error {
	//	//if err := Db.Table(user.TableName()).Create(&user).Error; err != nil {
	//	//	return errors.New("Insert to UserDatabase error")
	//	//}
	//	if err := tx.Table(user.TableName()).Create(&user).Error; err != nil {
	//		return errors.New("Insert to UserDatabase error")
	//
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return err
	//}
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
func (user *User) UpdateFollowCount(n int) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()

	if err := tx.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follow_count", int(user.FollowCount)+n).Error; err != nil {
		tx.Rollback()
		return errors.New("update follow count error, roll backed")

	}

	tx.Commit()
	mutex.Unlock()

	//err := Db.Transaction(func(tx *gorm.DB) error {
	//
	//	if err := tx.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follow_count", int(user.FollowCount)+n).Error; err != nil {
	//		return errors.New("update follow count error")
	//
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return err
	//}
	return nil
	//
	//result := Db.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follow_count", int(user.FollowCount)+n)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	return result.Error
	//}

}

// updata user follower_count by add n
func (user *User) UpdateFollowerCount(n int) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()

	if err := tx.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follower_count", int(user.FollowerCount)+n).Error; err != nil {
		tx.Rollback()
		return errors.New("update follower count error, roll backed")

	}

	tx.Commit()
	mutex.Unlock()
	return nil

	//err := Db.Transaction(func(tx *gorm.DB) error {
	//
	//	if err := tx.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follower_count", int(user.FollowerCount)+n).Error; err != nil {
	//		return errors.New("update follower count error")
	//
	//	}
	//	return nil
	//})
	//if err != nil {
	//	return err
	//}
	//return nil

	//result := Db.Table(user.TableName()).Where("id = ?", user.Id).First(user).Update("follower_count", int(user.FollowerCount)+n)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	return result.Error
	//}
	//
	//return nil
}
