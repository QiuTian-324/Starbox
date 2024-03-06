package article_api_code

import "BuzzBox/pkg/xcode"

var (
	GetBucketError            = xcode.New(30001, "获取存储桶失败")
	PutBucketError            = xcode.New(30002, "创建存储桶失败")
	InvalidFileTypeError      = xcode.New(30003, "文件类型无效")
	InvalidFileSizeError      = xcode.New(30004, "文件大小超过限制,最大4M")
	ArticleTitleEmpty         = xcode.New(30005, "文章标题不能为空")
	ArticleContentTooFewWords = xcode.New(30006, "文章内容太少")
	ArticleCoverEmpty         = xcode.New(30007, "文章封面不能为空")
	ArticlePublishFailed      = xcode.New(30008, "发布文章失败")
	NotLoginError             = xcode.New(30009, "未登录")
)
