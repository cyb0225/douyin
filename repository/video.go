// store videoinfos and related CRUD interface

package repository

// video mesages in database table
type Video struct {
	Id             uint64
	Title          string
	PlayUrl        string
	CoverUrl       string
	FavouriteCount uint64
	CommentCount   uint64
}
