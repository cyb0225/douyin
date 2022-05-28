package repository

import (
	"errors"

	"gorm.io/gorm"
)


type Follow struct {
	Id        uint64
	UserId     uint64 `gorm:"column:user_id"`
	FollowerId uint64 `gorm:"column:follower_id"`
	Status     int16  `gorm:"column:status"`
}

func (*Follow) TableName() string {
	return "follow"
}


// select user record by user_id
func (user *Follow) SelectByUserId() error {

	result := Db.Table(user.TableName()).Where("user_id = ?", user.UserId).First(user)

	// not found
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

func (user *Follow) CheckStatus() error {
	switch (user.Status) {
	case 0:
	case 1:
	case 2:
	case 3:
	default:
	}


	return nil
}