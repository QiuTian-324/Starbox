package applet_code

import "BuzzBox/pkg/xcode"

var (
	ErrMobileEmpty             = xcode.New(10001, "手机号不能为空")
	ErrNameEmpty               = xcode.New(10002, "用户名不能为空")
	ErrVerificationCodeEmpty   = xcode.New(10003, "验证码不能为空")
	ErrPasswordEmpty           = xcode.New(10004, "密码不能为空")
	ErrMobileFormatError       = xcode.New(10005, "手机号格式错误")
	ErrVerificationCodeExpired = xcode.New(10006, "验证码已过期")
	ErrVerificationCode        = xcode.New(10007, "验证码错误")
	ErrMobileExist             = xcode.New(10008, "手机号已经注册")
	ErrEncMobile               = xcode.New(10009, "加密手机号失败")
)
