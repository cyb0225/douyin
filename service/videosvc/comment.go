package videosvc

import (
	"github.com/2103561941/douyin/repository"
	"github.com/2103561941/douyin/service/usersvc"
	"time"
)

type Comment struct {
	UserId      uint64
	ToUserID    uint64
	VideoId     uint64
	ActionType  int
	CommentText string
	CommentID   uint64 // comment数据库的primary key
}

type CommentResponseWrapper struct {
	CommentID   uint64
	Userinfo    usersvc.UserInfo
	CommentText string
	CreateDate  string
}

func (comment *Comment) Comment() error {

	vidinfo := &repository.Video{
		UserId: comment.ToUserID, //视频创作者ID
		Id:     comment.VideoId,  //视频ID
	}
	if err := vidinfo.GetLikeInfo(); err != nil {
		return err
	}

	addCommentinfo := &repository.CommentTable{
		UserId:      comment.UserId,
		ToUserID:    vidinfo.UserId,
		VideoId:     vidinfo.Id,
		CommentText: comment.CommentText,
	}

	deleteCommentinfo := &repository.CommentTable{
		UserId:  comment.UserId,
		VideoId: vidinfo.Id,
		Id:      comment.CommentID,
	}

	if comment.ActionType == 1 {
		if err := addCommentinfo.AddComment(); err != nil {
			return err
		}
		if err := vidinfo.AddComment(vidinfo); err != nil {
			return err
		}
	}

	if comment.ActionType == 2 {
		if err := deleteCommentinfo.DeleteComment(); err != nil {
			return err
		}
		if err := vidinfo.DelComment(vidinfo); err != nil {
			return err
		}
	}
	return nil
}

func (comment *CommentResponseWrapper) GetCommentResponse(input *Comment) error {
	GetID := &repository.CommentTable{}

	comment.CommentID = GetID.GetCommentID()

	user := &usersvc.UserInfo{
		Id: input.UserId, //评论者ID
	}
	if err := user.SetUserInfo(input.UserId); err != nil {
		return err
	}
	comment.Userinfo = *user

	comment.CommentText = input.CommentText
	timeStr := time.Now().Format("01-02")
	comment.CreateDate = timeStr
	return nil
}
