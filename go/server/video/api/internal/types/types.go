// Code generated by goctl. DO NOT EDIT.
package types

type CommonResp struct {
}

type Video struct {
	User          User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	VideoId       int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Information   string `json:"information"`    // 视频简介
}

type VideoSimple struct {
	VideoId       int64  `json:"video_id"`
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	Information   string `json:"information"`    // 视频简介
}

type User struct {
	UserId   int64  `json:"user_id"`   // 用户id
	IsFollow bool   `json:"is_follow"` // true-已关注，false-未关注
	UserName string `json:"user_name"` // 用户名称
	Avator   string `json:"avator"`    //头像地址
}

type VideoListReq struct {
	UserId int64 `json:"user_id"`
	Tag    int64 `json:"tag"`
}

type VideoListResp struct {
	CommonResp
	VideoList []Video `json:"video_list"` // 视频列表
}

type PublishVideoReq struct {
	Information string `json:"information"`
	Tag         int64  `json:"tag"`
	UserId      int64  `json:"user_id"`
	Username    string `json:"user_name"`
	Avatar      string `json:"avatar"`
	Key         string `json:"key"`
}

type PublishVideoResp struct {
	Video Video `json:"video"`
}

type PublishTokenReq struct {
}

type PublishTokenResp struct {
	CommonResp
	Token string `json:"token"`
}

type PublishListReq struct {
	UserId int64 `json:"user_id"`
}

type PublishListResp struct {
	CommonResp
	VideoList []VideoSimple `json:"video_list"` // 用户发布的视频列表
}

type DeleteVideoReq struct {
	VideoId int64 `json:"video_id"`
}

type DeleteVideoResp struct {
	CommonResp
}

type FavoriteVideoListReq struct {
	UserId int64 `json:"user_id"`
}

type FavoriteVideoListResp struct {
	CommonResp
	VideoList []VideoSimple `json:"video_list"` // 用户喜欢的视频列表
}

type FavoriteReq struct {
	UserId  int64 `json:"user_id"`
	VideoId int64 `json:"video_id"`
	Status  int64 `json:"status"`
}

type FavoriteResp struct {
	CommonResp
}

type InfoReq struct {
	VideoId int64 `json:"video_id"`
}

type InfoResp struct {
	Video Video `json:"video"`
}

type SearchReq struct {
	Word string `json:"word"`
}

type SearchResp struct {
	VideoList []VideoSimple `json:"video_list"`
}
