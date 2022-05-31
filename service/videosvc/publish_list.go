package videosvc

import (
	"log"

	"github.com/2103561941/douyin/repository"
	"github.com/2103561941/douyin/service/usersvc"
)

// userInfo 是 author 的info
type VideoInfo struct {
	Id               uint64 `json:"id"`
	usersvc.UserInfo `json:"author"`
	PlayUrl          string `json:"play_url"`
	CoverUrl         string `json:"cover_url"`
	FavouriteCount   uint64 `json:"favorite_count"`
	CommentCount     uint64 `json:"comment_count"`
	IsFavorite       bool   `json:"is_favorite"`
	Title            string `json:"title"`
}

type PublishList struct {
	Author uint64
	UserId uint64
	Videos []*VideoInfo
}

func (list *PublishList) GetPublishList() error {
	video := &repository.Video{
		UserId: list.Author,
	}

	records, err := video.SelectPublishList()
	if err != nil {
		return err
	}

	tmpList := make([]*VideoInfo, len(records))

	for i := 0; i < len(records); i++ {
		videoInfo := &VideoInfo{}
		if err := videoInfo.SetVideoInfo(list.UserId, records[i]); err != nil {
			return err
		}
		tmpList[i] = videoInfo
	}

	list.Videos = tmpList

	log.Printf("service   %d\n", len(list.Videos))

	return nil
}

// UserId 是视频的发布者，
func (video *VideoInfo) SetVideoInfo(userId uint64, record *repository.Video) error {
	video.Id = record.Id

	// get userInfo
	
	video.UserInfo.Id = record.UserId
	if err := video.UserInfo.SetUserInfo(userId); err != nil {
		return err
	}

	video.PlayUrl = record.PlayUrl
	video.CoverUrl = record.CoverUrl
	video.FavouriteCount = record.FavouriteCount
	video.CommentCount = record.CommentCount
	video.IsFavorite = true
	video.Title = record.Title
	return nil
}
