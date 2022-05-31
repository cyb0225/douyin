package videosvc

import (
	"errors"

	"github.com/2103561941/douyin/repository"
)

type Like struct {
	UserId     uint64
	VideoId    uint64
	ActionType int
}

func (action *Like) Like() error {

	vidinfo := &repository.Video{
		UserId: action.UserId,
		Id:     action.VideoId,
	}

	if err := vidinfo.GetLikeInfo(); err != nil {
		return err
	}

	addlikeinfo := &repository.LikeTable{
		UserId:  vidinfo.UserId,
		VideoId: vidinfo.Id,
	}

	if err := addlikeinfo.GetLikeInfoinLike(); err != nil {
		return err
	}
	//上面这行是获取like表的信息。

	if action.ActionType == 1 { //点赞
		if addlikeinfo.ActionType == 1 {
			return errors.New("you can not like it again")
		}
		if err := vidinfo.Like(vidinfo); err != nil {
			return err
		}
	}
	if action.ActionType == 2 { //取消点赞
		if vidinfo.FavouriteCount == 0 {
			return errors.New("you can not unlike it again")
		}
		if err := vidinfo.UnLike(vidinfo); err != nil {
			return err
		}
	}

	if err := addlikeinfo.UpdateLike(action.ActionType); err != nil {
		return err
	}
	return nil
}
