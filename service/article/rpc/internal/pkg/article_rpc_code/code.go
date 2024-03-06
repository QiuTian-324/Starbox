package article_rpc_code

import "BuzzBox/pkg/xcode"

var (
	CodeNotLogin = xcode.New(31000, "未登录")

	CodeArticlePublishFailed = xcode.New(31001, "文章发布失败")
)
