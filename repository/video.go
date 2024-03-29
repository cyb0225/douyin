// store videoinfos and related CRUD interface

package repository

import (
	"errors"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

// video mesages in database table
type Video struct {
	Id             uint64    `gorm:"column:id"` //video ID
	UserId         uint64    `gorm:"column:user_id; index:idx_UserId"`
	Title          string    `gorm:"column:title"`
	PlayUrl        string    `gorm:"column:play_url"`
	CoverUrl       string    `gorm:"column:cover_url"`
	FavouriteCount uint64    `gorm:"column:favourite_count"`
	CommentCount   uint64    `gorm:"column:comment_count"`
	CreatedAt      time.Time `gorm:"column:create_time"`
}

func (*Video) TableName() string {
	return "video"
}

func (video *Video) Create() error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()

	if err := tx.Table(video.TableName()).Create(&video).Error; err != nil {
		tx.Rollback()
		return errors.New("Insert to UserDatabase -- video tabel error, roll backed")

	}

	tx.Commit()
	mutex.Unlock()

	return nil

}

func (video *Video) SelectPublishList() ([]*Video, error) {
	var records []*Video

	result := Db.Table(video.TableName()).Where("user_id = ?", video.UserId).Find(&records)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}

	log.Printf("repository   %d\n", len(records))

	return records, nil
}

func (video *Video) SelectVideoList(inputlist []*Video) ([]*Video, error) {
	records := make([]*Video, len(inputlist))
	if len(inputlist) == 0 {
		return records, nil
	}
	println(len(inputlist))
	for i := 0; i < len(inputlist); i++ {
		info := &Video{}
		result := Db.Table(video.TableName()).Where("id = ?", inputlist[i].Id).Find(&info)
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New(result.Error.Error())
		}
		records[i] = info
	}

	log.Printf("repository   %d\n", len(records))

	return records, nil
}

func (video *Video) GetLikeInfo() error {
	result := Db.Table(video.TableName()).Where("id = ?", video.Id).First(video)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return errors.New(result.Error.Error())
	}
	return nil
}

func (video *Video) Like(input *Video) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()

	if err := tx.Table(video.TableName()).Where("user_id = ? AND id = ?", video.UserId, video.Id).First(video).UpdateColumn("favourite_count", input.FavouriteCount+1).Error; err != nil {
		tx.Rollback()
		return errors.New("update like info error, roll backed")

	}

	tx.Commit()
	mutex.Unlock()

	return nil

}

func (video *Video) UnLike(input *Video) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(video.TableName()).Where("user_id = ? AND id = ?", video.UserId, video.Id).First(video).UpdateColumn("favourite_count", input.FavouriteCount-1).Error; err != nil {
		tx.Rollback()
		return errors.New("update unlike info error, roll backed")

	}

	tx.Commit()
	mutex.Unlock()

	return nil
}

func (video *Video) AddComment(input *Video) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(video.TableName()).Where("user_id = ? AND id = ?", video.UserId, video.Id).First(video).UpdateColumn("comment_count", input.CommentCount+1).Error; err != nil {
		tx.Rollback()
		return errors.New("add comment count error, roll backed")

	}
	tx.Commit()
	mutex.Unlock()

	return nil
}

func (video *Video) DelComment(input *Video) error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(video.TableName()).Where("user_id = ? AND id = ?", video.UserId, video.Id).First(video).UpdateColumn("comment_count", input.CommentCount-1).Error; err != nil {
		tx.Rollback()
		return errors.New("delete comment count error, roll backed")

	}
	tx.Commit()
	mutex.Unlock()

	return nil
}

func (video *Video) GetvideoBefore() ([]*Video, time.Time, error) {
	var records []*Video
	result := Db.Table(video.TableName()).Where("create_time < ?", video.CreatedAt).Find(&records)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, time.Time{}, errors.New(result.Error.Error())
	}
	if len(records) == 0 {
		return nil, time.Time{}, errors.New("NO VIDEO")

	}
	returned_video_earlist := records[0].CreatedAt //最早的视频的时间
	return records, returned_video_earlist, nil
}
