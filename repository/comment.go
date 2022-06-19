// 评论数据表模块

package repository

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"time"

	"gorm.io/gorm"
)

type CommentTable struct {
	Id          uint64    `gorm:"column:id"` //comment ID
	UserId      uint64    `gorm:"column:user_id; index:idx_UserId"`
	VideoId     uint64    `gorm:"column:video_id"`
	CommentText string    `gorm:"column:comment_text"`
	CreatedAt   time.Time `gorm:"column:create_time"`
}

func (*CommentTable) TableName() string {
	return "CommentTable"
}

func (comment *CommentTable) Insert() error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(comment.TableName()).Create(&comment).Error; err != nil {
		tx.Rollback()
		return errors.New("Insert to UserDatabase -- comment tabel error, roll backed")
	}
	tx.Commit()
	mutex.Unlock()

	err := tx.Migrator().HasIndex(&CommentTable{}, "idx_UserId")
	fmt.Println(err)

	// 插入到缓存中

	return nil
}

func (comment *CommentTable) Delete() error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(comment.TableName()).Where("user_id = ?  AND id = ? AND video_id = ?", comment.UserId, comment.Id, comment.VideoId).Delete(&comment).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", comment.Id)
	}
	tx.Commit()
	mutex.Unlock()

	return nil
}

// 插入评论数据
func (comment *CommentTable) AddComment() error {
	if err := comment.Insert(); err != nil {
		return err
	}
	return nil
}

// 删除评论数据
func (comment *CommentTable) DeleteComment() error {
	if err := comment.Delete(); err != nil {
		return err
	}
	return nil
}

// 获取评论id
func (comment *CommentTable) GetCommentID() (uint64, error) {
	var id []uint64
	err := Db.Raw("select LAST_INSERT_ID() as id from CommentTable").Pluck("id", &id)
	if err != nil {
		return 0, errors.New("getcommentIDerror")
	}
	return id[0], nil
}

// 获取评论列表
func (comment *CommentTable) GetCommentListRep() ([]*CommentTable, error) {
	var temp []*CommentTable

	result := Db.Table(comment.TableName()).Where("video_id = ?", comment.VideoId).Find(&temp)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, errors.New(result.Error.Error())
	}
	return temp, nil
}
