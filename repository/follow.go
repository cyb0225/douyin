package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Follow struct {
	Id       uint64 `gorm:"column:id;AUTO_INCREMENT"` //自增
	UserId   uint64 `gorm:"column:user_id"`
	ToUserId uint64 `gorm:"column:to_user_id"`
	Status   int `gorm:"column:status"`
}

func (*Follow) TableName() string {
	return "follow"
}

func (user *Follow) Insert() error {
	if err := Db.Table(user.TableName()).Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase -- Follow tabel error")
	}
	return nil

}

// select user record by user_id and  
func (user *Follow) Select() error {

	result := Db.Table(user.TableName()).Where("user_id = ? AND to_user_id = ?", user.UserId, user.ToUserId).First(user)

	// not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// updata follow status
func (user *Follow) UpdateStatus(newStatus int) error {
	result := Db.Table(user.TableName()).Where("user_id = ? AND to_user_id = ?", user.UserId, user.ToUserId).First(user).UpdateColumn("status", newStatus)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}


func (user *Follow) GetFollowList() ([]*User, error) {
	var records []*User
	result := Db.Table(user.TableName()).Where("user_id = ?", user.UserId).Find(records)
	
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	return records, nil
}

func (user *Follow) GetFollowerList() ([]*User, error) {
	var records []*User
	result := Db.Table(user.TableName()).Where("to_user_id = ?", user.ToUserId).Find(records)
	
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	return records, nil
}