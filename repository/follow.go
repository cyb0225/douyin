package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Follow struct {
	Id       uint64 `gorm:"column:id;AUTO_INCREMENT"` //自增
	UserId   uint64 `gorm:"column:user_id"`
	ToUserId uint64 `gorm:"column:to_user_id"`
	Status   uint64 `gorm:"column:status"`
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

func (user *Follow) UpdateStatus() error {
	result := Db.Table(user.TableName()).Where("user_id = ? OR to_user_id = ?", user.UserId, user.ToUserId).First(user).UpdateColumn("status", status)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}
