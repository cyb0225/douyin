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
	records, err := temp.SelectLikeList(video)

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
