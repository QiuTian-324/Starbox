package logic

import (
	"BuzzBox/pkg/encrypt"
	"BuzzBox/pkg/util"
	"BuzzBox/service/applet/internal/pkg/applet_code"
	"BuzzBox/service/applet/internal/pkg/applet_resp"
	"BuzzBox/service/user/rpc/user"
	"BuzzBox/service/user/rpc/userClient"
	"context"
	"strings"

	"BuzzBox/service/applet/internal/svc"
	"BuzzBox/service/applet/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)
	req.Password = strings.TrimSpace(req.Password)

	// 检查手机号是否为空
	if len(req.Mobile) == 0 {
		return nil, applet_code.ErrMobileEmpty
	}

	// 检查手机号格式
	if !CheckMobile(req.Mobile) {
		return nil, applet_code.ErrMobileFormatError
	}

	// 检查密码是否为空
	if len(req.Password) == 0 {
		return nil, applet_code.ErrPasswordEmpty
	}

	// 对手机号进行加密处理
	mobile, err := encrypt.EncMobile(req.Mobile)
	if err != nil {
		logx.Errorf("手机号加密失败: %s error: %v", req.Mobile, err)
		return nil, applet_code.ErrEncMobile
	}

	// 判断手机号是否已注册
	u, err := l.svcCtx.UserRPC.FindByMobile(l.ctx, &userClient.FindByMobileRequest{
		Mobile: mobile,
	})

	// 如果手机号未注册，则返回错误
	if err != nil || u.Id == 0 {
		// 手机号未注册
		return nil, err
	}

	// 根据用户手机号和密码进行登陆
	loginResp, err := l.svcCtx.UserRPC.Login(l.ctx, &user.LoginRequest{
		Mobile:   mobile,
		Password: encrypt.EncPassword(req.Password),
	})

	if err != nil {
		return nil, err
	}

	// 生成访问令牌,将用户ID作为JWT的payload
	token, err := util.BuildTokens(util.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": loginResp.Id,
		},
	})

	return &types.LoginResponse{
		ID: loginResp.Id,
		Token: types.Token{
			AccessToken:  token.AccessToken,
			AccessExpire: token.AccessExpire,
		},
		Message: applet_resp.LoginSuccess,
	}, nil
}
