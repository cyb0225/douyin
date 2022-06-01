package videosvc

import (
	"github.com/2103561941/douyin/repository"
	"github.com/2103561941/douyin/service/usersvc"
	"log"
)

type CommentList struct {
	UserID  uint64
	VideoID uint64
	Videos  []*CommentResponseWrapper
}

func (list *CommentList) GetCommentList() error {

	temp := &repository.CommentTable{
		VideoId: list.VideoID,
	}
	preprocess, err := temp.GetCommentListRep()
	if err != nil {
		return err
	}

	tmpList := make([]*CommentResponseWrapper, len(preprocess))
	for i := 0; i < len(preprocess); i++ {
		videoInfo := &CommentResponseWrapper{}
		if err := videoInfo.SetVideoInfo(list.UserID, preprocess[i]); err != nil {
			return err
		}
		tmpList[i] = videoInfo
	}
	list.Videos = tmpList
	log.Printf("service   %d\n", len(list.Videos))

	return nil

}

func (warrper *CommentResponseWrapper) SetVideoInfo(userId uint64, record *repository.CommentTable) error {
	warrper.CommentID = record.Id

	tempuserinfo := &usersvc.UserInfo{
		Id: record.UserId,
	}
	if err := tempuserinfo.SetUserInfo(userId); err != nil {
		return err
	}
	warrper.Userinfo = *tempuserinfo
	warrper.CommentText = record.CommentText
	timeString := record.CreatedAt.Format("01-02") //2015-06-15 08:52:32
	warrper.CreateDate = timeString
	return nil
}
