// Code generated by goctl. DO NOT EDIT.
package types

type InfoReq struct {
	UserId int64 `json:"user_id"`
}

type InfoResp struct {
	Avatar         string `json:"avatar"`
	FollowerCount  int64  `json:"follower_count"`
	FollowingCount int64  `json:"following_count"`
	IsFollow       bool   `json:"Is_follow"`
	UserID         int64  `json:"user_id"`
	UserName       string `json:"user_name"`
}

type FollowinglistReq struct {
	UserID int64 `json:"user_id"`
}

type FollowinglistResp struct {
	UserList []UserSimple `json:"user_list"`
}

type FollowerlistReq struct {
	UserID int64 `json:"user_id"`
}

type FollowerlistResp struct {
	UserList []UserSimple `json:"user_list"`
}

type PublishAvatarReq struct {
	Key    string `json:"key"`
	UserID int64  `json:"user_id"`
}

type PublishAvatarResp struct {
	Avatar string `json:"avatar"`
}

type FollowReq struct {
	FollowerID  int64 `json:"follower_id"`
	FollowingID int64 `json:"following_id"`
	Status      int64 `json:"status"`
}

type FollowResp struct {
}

type NameReq struct {
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}

type NameResp struct {
}

type UserSimple struct {
	Avatar   string `json:"avatar"`
	IsFollow bool   `json:"Is_follow"`
	UserID   int64  `json:"user_id"`
	UserName string `json:"user_name"`
}