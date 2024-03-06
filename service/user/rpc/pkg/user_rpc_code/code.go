package user_rpc_code

import "BuzzBox/pkg/xcode"

var (
	ErrRegisterFail     = xcode.New(21001, "注册失败")
	ErrPassword         = xcode.New(21002, "密码错误")
	ErrNotRegister      = xcode.New(21003, "用户未注册")
	ErrUserIDInsertFail = xcode.New(21004, "用户ID插入失败")
	ErrMobileNotExist   = xcode.New(21012, "手机号未注册")
	ErrMobileExist      = xcode.New(21013, "手机号已注册")
)
