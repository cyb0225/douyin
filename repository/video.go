// store videoinfos and related CRUD interface

package repository

// video mesages in database table
type Video struct {
	Id             uint64 `gorm:"column:id"`
	UserId		   uint64 `gorm:"column:user_id"`
	Title          string `gorm:"column:title"`
	PlayUrl        string `gorm:"column:play_url"`
	CoverUrl       string `gorm:"column:cover_url"`
	FavouriteCount uint64 `gorm:"column:favourite_count"`
	CommentCount   uint64 `gorm:"column:comment_count"`
}

