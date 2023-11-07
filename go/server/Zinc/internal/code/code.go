package code

import "Thinkphoto/pkg/xcode"

var (
	SearchDataEmpty        = xcode.New(60001, "搜索内容不能为空")
	SearchTypeEmpty        = xcode.New(60002, "搜索类型不能为空")
	DeleteInformationEmpty = xcode.New(60006, "删除信息为空")
)
