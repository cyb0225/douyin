package videosvc

import "github.com/2103561941/douyin/repository"

type PublishVideo struct {
	UserID   uint64
	PlayURL  string
	CoverURL string
	Title    string
}

func (video *PublishVideo) PublishVideo() error {
	videoinfo := &repository.Video{
		UserId:   video.UserID,
		PlayUrl:  video.PlayURL,
		CoverUrl: video.CoverURL,
		Title:    video.Title,
	}

	if err := videoinfo.Create(); err != nil {
		return err
	}

	return nil
}
