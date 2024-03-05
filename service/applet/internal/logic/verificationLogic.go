package logic

import (
	"BuzzBox/pkg/util"
	"BuzzBox/service/applet/internal/pkg/applet_code"
	"BuzzBox/service/applet/internal/pkg/applet_resp"
	"BuzzBox/service/applet/internal/svc"
	"BuzzBox/service/applet/internal/types"
	"BuzzBox/service/user/rpc/userClient"
	"context"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"strconv"
	"strings"
	"time"
)

var (
	prefixVerificationCount = "biz#verification#count#%s" // 验证码计数的 Redis 键前缀
	verificationLimitPerDay = 100                         // 每日发送验证码的上限
	expireActivation        = 60 * 100000                 // 验证码缓存失效时间，单位为秒
)

type VerificationLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVerificationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VerificationLogic {
	return &VerificationLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VerificationLogic) Verification(req *types.VerificationRequest) (resp *types.VerificationResponse, err error) {
	req.Mobile = strings.TrimSpace(req.Mobile)

	// 检查手机号是否为空
	if len(req.Mobile) == 0 {
		return nil, applet_code.ErrMobileEmpty
	}

	// 检查手机号格式
	if !CheckMobile(req.Mobile) {
		return nil, applet_code.ErrMobileFormatError
	}

	// 获取今日发送验证码次数
	count, err := l.getVerificationCount(req.Mobile)
	if err != nil {
		logx.Errorf("getVerificationCount mobile: %s error: %v", req.Mobile, err)
	}
	if count > verificationLimitPerDay {
		return nil, err
	}

	// 生成或获取已存在的验证码
	code, err := getActivationCache(req.Mobile, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("getActivationCache mobile: %s error: %v", req.Mobile, err)
	}
	if len(code) == 0 {
		code = util.RandomNumeric(6)
	}

	// 发送验证码
	_, err = l.svcCtx.UserRPC.SendSms(l.ctx, &userClient.SendSmsRequest{
		Mobile: req.Mobile,
	})

	if err != nil {
		logx.Errorf("sendSms mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}

	// 缓存验证码
	err = saveActivationCache(req.Mobile, code, l.svcCtx.BizRedis)
	if err != nil {
		logx.Errorf("saveActivationCache mobile: %s error: %v", req.Mobile, err)
		return nil, err
	}

	// 增加验证码发送次数计数
	err = l.incrVerificationCount(req.Mobile)
	if err != nil {
		logx.Errorf("incrVerificationCount mobile: %s error: %v", req.Mobile, err)
	}

	return &types.VerificationResponse{
		Message: applet_resp.SendSuccess,
	}, nil

}

// getVerificationCount 获取今日验证码发送次数
func (l *VerificationLogic) getVerificationCount(mobile string) (int, error) {
	key := fmt.Sprintf(prefixVerificationCount, mobile)
	val, err := l.svcCtx.BizRedis.Get(key)
	if err != nil {
		return 0, err
	}
	if len(val) == 0 {
		return 0, nil
	}

	return strconv.Atoi(val)
}

// incrVerificationCount 增加验证码发送次数计数
func (l *VerificationLogic) incrVerificationCount(mobile string) error {
	key := fmt.Sprintf(prefixVerificationCount, mobile)
	_, err := l.svcCtx.BizRedis.Incr(key)
	if err != nil {
		return err
	}

	return l.svcCtx.BizRedis.Expireat(key, util.EndOfDay(time.Now()).Unix())
}

// getActivationCache 获取验证码缓存
func getActivationCache(mobile string, rds *redis.Redis) (string, error) {
	key := fmt.Sprintf(prefixActivation, mobile)
	return rds.Get(key)
}

// saveActivationCache 缓存验证码
func saveActivationCache(mobile, code string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefixActivation, mobile)
	return rds.Setex(key, code, expireActivation)
}

// delActivationCache 删除验证码缓存
func delActivationCache(mobile, code string, rds *redis.Redis) error {
	key := fmt.Sprintf(prefixActivation, mobile)
	_, err := rds.Del(key)
	return err
}
