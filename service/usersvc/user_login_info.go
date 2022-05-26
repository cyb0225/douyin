package usersvc

type UserJsonInfo struct {
	ID            uint64    `json:"id"`
	Username      string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json:"is_follow"`
}
