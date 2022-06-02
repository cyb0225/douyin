package videosvc

import (
	"github.com/2103561941/douyin/repository"
	"log"
)

type LikeList struct {
	Author uint64
	UserId uint64
	Videos []*VideoInfo
}

func (list *PublishList) GetLikeList() error {
	temp := &repository.Video{
		UserId: list.Author,
	}

	video := &repository.LikeTable{
		UserId: list.Author,
	}
	preprocess, err := temp.SelectLikeList(video)
	if err != nil {
		return err
	}

	records, err := temp.SelectVideoList(preprocess)
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

func (video *VideoInfo) presetinfo(userId uint64, record *repository.Video) error {
	video.Id = record.Id

	// get userInfo

	video.UserInfo.Id = record.UserId
	if err := video.UserInfo.SetUserInfo(userId); err != nil {
		return err
	}

	video.PlayUrl = SocksUrl + record.PlayUrl
	video.CoverUrl = SocksUrl + record.CoverUrl
	video.FavouriteCount = record.FavouriteCount
	video.CommentCount = record.CommentCount
	video.IsFavorite = true
	video.Title = record.Title
	return nil
}
