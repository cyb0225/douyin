package repository

import (
	"errors"

	"gorm.io/gorm"
)

type Follow struct {
	Id       uint64 `gorm:"column:id;AUTO_INCREMENT"` 
	UserId   uint64 `gorm:"column:user_id; index:idx_UserId"` // 当前用户
	ToUserId uint64 `gorm:"column:to_user_id"` // 被关注的用户
	Status   int    `gorm:"column:status"`
}

func (*Follow) TableName() string {
	return "follow"
}

// 插入数据
func (user *Follow) Insert() error {
	user.Status = 0
	if err := Db.Table(user.TableName()).Create(&user).Error; err != nil {
		return errors.New("Insert to UserDatabase -- Follow tabel error")
	}
	err := Db.Migrator().HasIndex(&Follow{}, "idx_UserId")
	println(err)
	return nil

}

// 通过用户id查询用户记录
func (user *Follow) Select() error {

	result := Db.Table(user.TableName()).Where("user_id = ? AND to_user_id = ?", user.UserId, user.ToUserId).First(user)

	// 查找失败
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// 更新表单数据
func (user *Follow) UpdateStatus(newStatus int) error {
	result := Db.Table(user.TableName()).Where("user_id = ? AND to_user_id = ?", user.UserId, user.ToUserId).First(user).UpdateColumn("status", newStatus)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// 回调函数
func (user *Follow) Undo(follow *Follow) error {
	result := Db.Table(user.TableName()).Where("user_id = ? AND to_user_id = ?", user.UserId, user.ToUserId).First(user).UpdateColumn("status", follow.Status)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}

// 获取用户关注列表
func (user *Follow) GetFollowList() ([]*Follow, error) {
	var records []*Follow
	result := Db.Table(user.TableName()).Where("user_id = ? AND status = ?", user.UserId, 1).Find(&records)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	return records, nil
}

// 获取用户粉丝列表
func (user *Follow) GetFollowerList() ([]*Follow, error) {
	var records []*Follow
	result := Db.Table(user.TableName()).Where("to_user_id = ? AND status = ?", user.ToUserId, 1).Find(&records)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	return records, nil
}
