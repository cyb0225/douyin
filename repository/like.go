package repository

import (
	"errors"

	"gorm.io/gorm"
)

type LikeTable struct {
	Id         uint64 `gorm:"column:id"`
	UserId     uint64 `gorm:"column:user_id; index:idx_UserId"`
	ToUserID   uint64 `gorm:"column:to_user_id"`
	VideoId    uint64 `gorm:"column:video_id"`
	ActionType int    `gorm:"column:action_type"`
}

func (*LikeTable) TableName() string {
	return "LikeTable"
}

func (like *LikeTable) Create() error {
	like.ActionType = 0
	if err := Db.Table(like.TableName()).Create(&like).Error; err != nil {
		return errors.New("Insert to UserDatabase -- like tabel error")
	}

	err := Db.Migrator().HasIndex(&LikeTable{}, "idx_UserId")
	println(err)
	return nil
}

func (like *LikeTable) UpdateLike(act int) error {

	if act == 1 { //如果喜欢
		result := Db.Table(like.TableName()).Where("user_id = ? AND to_user_id = ? AND video_id = ?", like.UserId, like.ToUserID, like.VideoId).First(like).UpdateColumn("action_type", 1)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) { //没找到就新建
			like.ActionType = 1
			if err := like.Create(); err != nil {
				return err
			}
		}
		/*
			这个位置的没找到就新建存在逻辑问题。如果liketable没有存储视频作者ID并配套查询的话，会出现其他用户登录后可以直接取消赞的问题。
			如果想要在liketable省略视频作者ID的话就需要多查询一次映射或者是修改此处逻辑。
			尽管我们的视频ID不唯一，可以通过视频ID查找到作者ID。但是我觉得效果并不好。
		*/

	}
	if act == 2 { //如果不喜欢
		result := Db.Table(like.TableName()).Where("user_id = ? AND to_user_id = ? AND video_id = ?", like.UserId, like.ToUserID, like.VideoId).First(like).UpdateColumn("action_type", 0)

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return errors.New(result.Error.Error())
		}
	}

	return nil
}

func (like *LikeTable) GetLikeInfoinLike() error {
	result := Db.Table(like.TableName()).Where("user_id = ? AND to_user_id = ? AND video_id = ?", like.UserId, like.ToUserID, like.VideoId).First(like)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) { //没找到就新建
		like.ActionType = 1
		if err := like.Create(); err != nil {
			return err
		}
	}
	return nil
}

func (video *Video) SelectLikeList(like *LikeTable) ([]*Video, error) {
	var temp []*LikeTable
	result := Db.Table(like.TableName()).Where("user_id = ? AND action_type = ?", like.UserId, 1).Find(&temp)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	records := make([]*Video, len(temp))

	for i := 0; i < len(temp); i++ {
		info := &Video{}
		info.Swapinfo(temp[i])
		records[i] = info
		println(records[i].Id)
	}

	return records, nil

}

func (video *Video) Swapinfo(like *LikeTable) {
	video.Id = like.VideoId
	video.UserId = like.UserId
}

func (like *LikeTable) IsFavorite() error {

	result := Db.Table(like.TableName()).Where("user_id = ? AND video_id = ? AND action_type = ?", like.UserId, like.VideoId, 1).First(like)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}

	return nil
}
