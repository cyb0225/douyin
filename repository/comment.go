// 评论数据表模块

package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sync"
	"time"
)

type CommentTable struct {
	Id     uint64 `gorm:"column:id"` //comment ID
	UserId uint64 `gorm:"column:user_id; index:idx_UserId"`
	VideoId     uint64    `gorm:"column:video_id"`
	CommentText string    `gorm:"column:comment_text"`
	CreatedAt   time.Time `gorm:"column:create_time"`
}

func (*CommentTable) TableName() string {
	return "CommentTable"
}

func (comment *CommentTable) Create() error {

	var mutex sync.Mutex
	mutex.Lock()
	tx := Db.Begin()
	if err := tx.Table(comment.TableName()).Create(&comment).Error; err != nil {
		tx.Rollback()
		return errors.New("Insert to UserDatabase -- comment tabel error, roll backed")
	}
	tx.Commit()
	mutex.Unlock()

	//
	//if err := Db.Table(comment.TableName()).Create(&comment).Error; err != nil {
	//	return errors.New("Insert to UserDatabase -- comment tabel error")
	//}

	err := tx.Migrator().HasIndex(&CommentTable{}, "idx_UserId")
	println(err)
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

	//result := Db.Table(comment.TableName()).Where("user_id = ?  AND id = ? AND video_id = ?", comment.UserId, comment.Id, comment.VideoId).Delete(&comment)
	////if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	////	return errors.New(result.Error.Error())
	////}
	//if result.Error != nil {
	//	return result.Error
	//} else if result.RowsAffected < 1 {
	//	return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", comment.Id)
	//}
	return nil
}

// 插入评论数据
func (comment *CommentTable) AddComment() error {
	if err := comment.insert(); err != nil {
		return err
	}
	return nil
}

// 删除评论数据
func (comment *CommentTable) DeleteComment() error {
	if err := comment.delete(); err != nil {
		return err
	}
	return nil
}

// 获取评论id
func (comment *CommentTable) GetCommentID() uint64 {
	var id []uint64
	Db.Raw("select LAST_INSERT_ID() as id from CommentTable").Pluck("id", &id)
	return id[0]
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

// -------------------------------------------------------------------------------------------------------------------
// 插入数据
func (comment *CommentTable) insert() error {
	if err := Db.Table(comment.TableName()).Create(&comment).Error; err != nil {
		return errors.New("Insert to UserDatabase -- comment tabel error")
	}

	// err := Db.Migrator().HasIndex(&CommentTable{}, "idx_UserId")
	// println(err)

	return nil
}

// 删除数据
func (comment *CommentTable) delete() error {
	result := Db.Table(comment.TableName()).
	Where("user_id = ?  AND id = ? AND video_id = ?", comment.UserId, comment.Id, comment.VideoId).Delete(&comment)

	if result.Error != nil {
		return result.Error
	} else if result.RowsAffected < 1 {
		return fmt.Errorf("row with id=%d cannot be deleted because it doesn't exist", comment.Id)
	}
	return nil
}