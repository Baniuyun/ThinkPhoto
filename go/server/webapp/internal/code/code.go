package code

import "Thinkphoto/pkg/xcode"

var (
	UserNameEmpty         = xcode.New(10001, "用户名为空")
	PasswordEmpty         = xcode.New(10002, "密码为空")
	PasswordErr           = xcode.New(10003, "密码或账户错误")
	UserNameHasRegistered = xcode.New(10004, "用户名已经被注册")
)
