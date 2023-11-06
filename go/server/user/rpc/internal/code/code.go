package code

import "Thinkphoto/pkg/xcode"

var (
	RegisterNameEmpty     = xcode.New(30001, "注册账户不能为空")
	RegisterPasswordEmpty = xcode.New(30002, "注册密码不能为空")
	RegisterNameRepeat    = xcode.New(30003, "注册账户重复")
)
