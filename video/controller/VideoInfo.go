package controller

import "github.com/2103561941/douyin/user/controller"

type VideoInfo struct {
	Id            int             `json:"id"`
	Author        controller.User `json:"author"`
	PlayURL       string          `json:"play_url"`
	CoverURL      string          `json:"cover_url"`
	FavoriteCount int             `json:"favorite_count"`
	CommentCount  int             `json:"comment_count"`
	IsFavorite    bool            `json:"is_favorite"`
	Title         string          `json:"title"`
}
