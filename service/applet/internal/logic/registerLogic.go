package logic

import (
	"BuzzBox/pkg/encrypt"
	"BuzzBox/service/applet/internal/pkg/applet_code"
	"BuzzBox/service/applet/internal/pkg/applet_resp"
	"BuzzBox/service/user/rpc/userClient"
	"context"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"regexp"
	"strings"

	"BuzzBox/service/applet/internal/svc"
	"BuzzBox/service/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	prefixActivation = "biz#activation#%s"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterRequest) (resp *types.RegisterResponse, err error) {
	// 去除请求中的空白字符
	req.Username = strings.TrimSpace(req.Username)
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)
	req.VerificationCode = strings.TrimSpace(req.VerificationCode)

	// 检查用户名是否为空
	if len(req.Username) == 0 {
		return nil, applet_code.ErrNameEmpty
	}

	// 检查手机号是否为空
	if len(req.Mobile) == 0 {
		return nil, applet_code.ErrMobileEmpty
	}

	// 检查手机号格式
	if !CheckMobile(req.Mobile) {
		return nil, applet_code.ErrMobileFormatError
	}

	// 检查验证码是否为空
	if len(req.VerificationCode) == 0 {
		return nil, applet_code.ErrVerificationCodeEmpty
	}

	// 检查密码是否为空
	if len(req.Password) == 0 {
		return nil, applet_code.ErrPasswordEmpty
	}

	// 检查验证码是否正确
	err = checkVerificationCode(l.svcCtx.BizRedis, req.Mobile, req.VerificationCode)
	if err != nil {
		return nil, err
	}

	// 对手机号进行加密处理
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("手机号加密失败: %s error: %v", req.Mobile, err)
		return nil, applet_code.ErrEncMobile
	}

	// 判断手机号是否已注册
	u, _ := l.svcCtx.UserRPC.FindByMobile(l.ctx, &userClient.FindByMobileRequest{
		Mobile: mobile,
	})

	// 如果手机号已注册，则返回错误
	if u != nil && u.Id > 0 {
		return nil, applet_code.ErrMobileExist
	}

	// 对密码进行加密处理
	password := encrypt.EncPassword(req.Password)

	// 创建用户
	register, err := l.svcCtx.UserRPC.Register(l.ctx, &userClient.RegisterRequest{
		Username: req.Username,
		Mobile:   mobile,
		Password: password,
	})
	if err != nil {
		logx.Errorf("注册用户报错: %v", err)
		// 注册失败
		return nil, err
	}

	//// 删除验证码缓存
	//_ = delActivationCache(req.Mobile, req.VerificationCode, l.svcCtx.BizRedis)

	// 返回注册成功的响应
	logx.Infof("注册成功, userId: %d", register)
	return &types.RegisterResponse{
		ID:      register.Id,
		Message: applet_resp.RegisterSuccess,
	}, err
}

// checkVerificationCode 函数用于检查验证码是否正确
func checkVerificationCode(rds *redis.Redis, mobile, c string) error {
	// 从缓存中获取验证码
	cacheCode, err := getActivationCache(mobile, rds)
	if err != nil {
		return applet_code.ErrVerificationCode
	}
	// 如果缓存中的验证码与请求中的验证码不一致，返回验证码错误
	if cacheCode != c {
		return applet_code.ErrVerificationCode
	}
	// 如果缓存中的验证码为空，说明验证码已过期
	if cacheCode == "" {
		return applet_code.ErrVerificationCodeExpired
	}
	return nil
}

// CheckMobile 检查手机号格式是否正确
func CheckMobile(mobile string) bool {
	// 手机号正则表达式
	// 中国大陆手机号格式：11位数字，以1开头
	// 正则表达式中：^表示匹配字符串的开头，$表示匹配字符串的结尾
	reg := `^1\d{10}$`

	// 编译正则表达式
	regex := regexp.MustCompile(reg)

	// 使用正则表达式匹配手机号
	return regex.MatchString(mobile)
}
