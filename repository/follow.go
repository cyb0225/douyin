package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Follow struct {
	Id       uint64 `gorm:"column:id;AUTO_INCREMENT"` //自增
	UserId   uint64 `gorm:"column:user_id; index:idx_UserId"`
	ToUserId uint64 `gorm:"column:to_user_id"`
	Status   int    `gorm:"column:status"`
}

func (*Follow) TableName() string {
	return "follow"
}

func (user *Follow) Insert() error {
	user.Status = 0
	if err := Db.Table(user.TableName()).Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase -- Follow tabel error")
	}
	err := Db.Migrator().HasIndex(&Follow{}, "idx_UserId")
	println(err)
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

// use to go back 
func (user *Follow) Undo(follow *Follow) error {
	result := Db.Table(user.TableName()).Where("user_id = ? AND to_user_id = ?", user.UserId, user.ToUserId).First(user).UpdateColumn("status", follow.Status)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}


/*
	SELECT * from follow
	WHERE
	user_id = id,
	status = 1
*/
func (user *Follow) GetFollowList() ([]*Follow, error) {
	var records []*Follow
	result := Db.Table(user.TableName()).Where("user_id = ? AND status = ?", user.UserId, 1).Find(&records)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	return records, nil
}

func (user *Follow) GetFollowerList() ([]*Follow, error) {
	var records []*Follow
	result := Db.Table(user.TableName()).Where("to_user_id = ? AND status = ?", user.ToUserId, 1).Find(&records)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	return records, nil
}
