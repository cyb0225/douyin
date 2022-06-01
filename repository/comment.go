package repository

import (
	"errors"
	"fmt"
	"time"
)

type CommentTable struct {
	Id          uint64    `gorm:"column:id"` //comment ID
	UserId      uint64    `gorm:"column:user_id; index:idx_UserId"`
	ToUserID    uint64    `gorm:"column:to_user_id"` //author ID
	VideoId     uint64    `gorm:"column:video_id"`
	CommentText string    `gorm:"column:comment_text"`
	CreatedAt   time.Time `gorm:"column:create_time"`
}

func (*CommentTable) TableName() string {
	return "CommentTable"
}

func (comment *CommentTable) Create() error {
	if err := Db.Table(comment.TableName()).Create(&comment).Error; err != nil {
		return errors.New("Insert to UserDatabase -- comment tabel error")
	}

	err := Db.Migrator().HasIndex(&CommentTable{}, "idx_UserId")
	println(err)
	return nil
}

func (comment *CommentTable) Delete() error {
	result := Db.Table(comment.TableName()).Where("user_id = ?  AND id = ? AND video_id = ?", comment.UserId, comment.Id, comment.VideoId).Delete(&comment)
	//if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	//	return errors.New(result.Error.Error())
	//}
	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", comment.Id)
	}
	return nil
}

func (comment *CommentTable) AddComment() error {
	if err := comment.Create(); err != nil {
		return err
	}
	return nil
}

func (comment *CommentTable) DeleteComment() error {
	if err := comment.Delete(); err != nil {
		return err
	}
	return nil
}
