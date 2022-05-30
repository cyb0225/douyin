// store videoinfos and related CRUD interface

package repository

import (
	"errors"
	"time"
)

// video mesages in database table
type Video struct {
	Id             uint64    `gorm:"column:id"` //video ID
	UserId         uint64    `gorm:"column:user_id, index:idx_UserId"`
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
	if err := Db.Table(video.TableName()).Create(&video).Error; err != nil {
		return errors.New("Insert to UserDatabase -- video tabel error")
	}

	err := Db.Migrator().HasIndex(&Video{}, "idx_UserId")
	println(err)
	return nil
}
