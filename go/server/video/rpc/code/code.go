package code

import "Thinkphoto/pkg/xcode"

var (
	UserIdEmpty   = xcode.New(20001, "用户id为空")
	VideoNil      = xcode.New(20002, "无此条视频")
	DeleteUserErr = xcode.New(20003, "删除者不是本人")
	SqlEmpty      = xcode.New(20004, "Sql返回为空")
)
