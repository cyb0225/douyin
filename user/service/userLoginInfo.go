// 用户鉴权逻辑判断，并返回用户信息
package service

// user 页面需要的用户基本信息
type UserJsonInfo struct {
	ID            int    `json:"id"`
	Username      string `json:"name"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
	IsFollow      bool   `json: "is_follow"`
}

