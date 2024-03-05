package user_code

import "BuzzBox/pkg/xcode"

var (
	ErrRegisterFail     = xcode.New(20001, "注册失败")
	ErrPassword         = xcode.New(20002, "密码错误")
	ErrNotRegister      = xcode.New(20003, "用户未注册")
	ErrUserIDInsertFail = xcode.New(20004, "用户ID插入失败")
	ErrMobileNotExist   = xcode.New(10012, "手机号未注册")
	ErrMobileExist      = xcode.New(10013, "手机号已注册")
)
