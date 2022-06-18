package repository

import (
	"errors"
	"sync"

	"gorm.io/gorm"
)

type LikeTable struct {
	Id         uint64 `gorm:"column:id"`
	UserId     uint64 `gorm:"column:user_id; index:idx_UserId"`
	VideoId    uint64 `gorm:"column:video_id"`
	ActionType int    `gorm:"column:action_type"`
}

func (*LikeTable) TableName() string {
	return "LikeTable"
}


// 插入数据
func (like *LikeTable) Insert() error {
	like.ActionType = 0

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(like.TableName()).Create(&like).Error; err != nil {
		tx.Rollback()
		return errors.New("update follow count error, roll backed")

	}
	tx.Commit()
	mutex.Unlock()

	return nil
}

// 更新用户喜欢数据
func (like *LikeTable) UpdateLike(act int) error {
	//这块不能用事务。因为要通过报错新建。而且新建已经用事务了
	if act == 1 { //如果喜欢
		result := Db.Table(like.TableName()).Where("user_id = ? AND video_id = ?", like.UserId, like.VideoId).First(like).UpdateColumn("action_type", 1)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) { //没找到就新建
			like.ActionType = 1
			if err := like.Insert(); err != nil {
				return err
			}
		}
	}
	
	if act == 2 { //如果不喜欢
		result := Db.Table(like.TableName()).Where("user_id = ? AND video_id = ?", like.UserId, like.VideoId).First(like).UpdateColumn("action_type", 0)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New(result.Error.Error())
		}
	}

	return nil
}

// 获取
func (like *LikeTable) GetLikeInfoinLike() error {
	result := Db.Table(like.TableName()).Where("user_id = ? AND video_id = ?", like.UserId, like.VideoId).First(like)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //未找到就新建
		like.ActionType = 1
		if err := like.Insert(); err != nil {
			return err
		}
	}
	return nil
}

// 获取用户喜欢列表
func (video *Video) SelectLikeList(like *LikeTable) ([]*Video, error) {
	var temp []*LikeTable
	result := Db.Table(like.TableName()).Where("user_id = ? AND action_type = ?", like.UserId, 1).Find(&temp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	records := make([]*Video, len(temp))

	for i := 0; i < len(temp); i++ {
		info := &Video{}
		info.swapinfo(temp[i])
		records[i] = info
		println(records[i].Id)
	}

	return records, nil

}


func (video *Video) swapinfo(like *LikeTable) {
	video.Id = like.VideoId
	video.UserId = like.UserId
}


// 判断用户是否喜欢该视频
func (like *LikeTable) IsFavorite() error {

	result := Db.Table(like.TableName()).Where("user_id = ? AND video_id = ? AND action_type = ?", like.UserId, like.VideoId, 1).First(like)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}
