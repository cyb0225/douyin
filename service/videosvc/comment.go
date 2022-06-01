package videosvc

import "github.com/2103561941/douyin/repository"

type Comment struct {
	UserId      uint64
	ToUserID    uint64
	VideoId     uint64
	ActionType  int
	CommentText string
	CommentID   uint64 // comment数据库的primary key
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
