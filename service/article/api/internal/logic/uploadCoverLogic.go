package logic

import (
	"BuzzBox/pkg/util"
	"BuzzBox/service/article/api/internal/pkg/article_api_code"
	"BuzzBox/service/article/api/internal/svc"
	"BuzzBox/service/article/api/internal/types"
	"context"
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/zeromicro/go-zero/core/logx"
)

type UploadCoverLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadCoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UploadCoverLogic {
	return &UploadCoverLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

const maxFileSize = 4 << 20

func (l *UploadCoverLogic) UploadCover(req *http.Request) (resp *types.UploadCoverResponse, err error) {
	_ = req.ParseMultipartForm(maxFileSize)

	file, handler, err := req.FormFile("cover")
	if err != nil {
		return nil,

			article_api_code.InvalidFileTypeError
	}

	defer file.Close()

	if handler.Size > maxFileSize {
		return nil, article_api_code.InvalidFileSizeError
	}

	// 检查文件类型
	ext := filepath.Ext(handler.Filename)
	if ext != ".jpg" && ext != ".jpeg" && ext != ".png" {
		return nil, article_api_code.InvalidFileTypeError
	}

	bucket, err := l.svcCtx.OssClient.Bucket(l.svcCtx.Config.Oss.BucketName)
	if err != nil {
		logx.Errorf("get bucket failed, err: %v", err)
		return nil, article_api_code.GetBucketError
	}
	objectKey := "images/" + genFilename(handler.Filename)
	err = bucket.PutObject(objectKey, file)
	if err != nil {
		logx.Errorf("put object failed, err: %v", err)
		return nil, article_api_code.PutBucketError
	}

	return &types.UploadCoverResponse{CoverUrl: genFileURL(objectKey)}, nil
}

func genFilename(filename string) string {
	// 获取文件的后缀名
	ext := filepath.Ext(filename)
	// 生成一个新的文件名，使用 UUID
	// 例如：如果传入的文件名为 "example.png"，则返回 "example_5f9c5e8d2f.png"
	return fmt.Sprintf("%v_%s", util.GenUUID(20), ext)
}

func genFileURL(objectKey string) string {
	return fmt.Sprintf("https://buzzbox.oss-cn-shenzhen.aliyuncs.com/%s", objectKey)
}
