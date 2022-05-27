// store videoinfos and related CRUD interface

package repository

import (
	"gorm.io/gorm"

)

var (
	VideoDB *gorm.DB
)

// video mesages in database table
type Video struct {
	ID             uint64
	Title          string
	PlayUrl        string
	CoverUrl       string
	FavouriteCount uint64
	CommentCount   uint64
}
