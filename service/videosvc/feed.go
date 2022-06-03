package videosvc

import (
	"errors"
	"github.com/2103561941/douyin/repository"
	"log"
	"time"
)

type Feedliststruct struct {
	Latest_time   string
	UserID        uint64
	Videos        []*VideoInfo
	Earlist_video string
}

func (list *Feedliststruct) Feedlist() error {
	if list.Latest_time == "0" {
		log.Printf("true")
		list.Latest_time = time.Now().Format("2006-01-02 15:04:05")
	}
	println(list.Latest_time)
	the_time, err := time.ParseInLocation("2006-01-02 15:04:05", list.Latest_time, time.Local) //输入的时间转换为时间类型
	if err != nil {
		return errors.New("time convert error")
	}
	selecttime := &repository.Video{
		CreatedAt: the_time,
	}
	log.Printf("-------------", the_time)

	videolist, earliest_video, err := selecttime.GetvideoBefore()

	if err != nil {
		return err
	}
	tmpList := make([]*VideoInfo, len(videolist))
	for i := 0; i < len(videolist); i++ {
		videoInfo := &VideoInfo{}
		if err := videoInfo.SetVideoInfo(list.UserID, videolist[i]); err != nil {
			return err
		}
		tmpList[i] = videoInfo
	}
	list.Videos = tmpList
	list.Earlist_video = earliest_video.Format("2006-01-02 15:04:05")
	println(list.Earlist_video)
	log.Printf("service   %d\n", len(list.Videos))

	return nil
}
