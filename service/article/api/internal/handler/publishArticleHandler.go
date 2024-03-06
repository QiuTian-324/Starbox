package handler

import (
	"BuzzBox/service/article/api/internal/logic"
	"net/http"

	"BuzzBox/service/article/api/internal/svc"
	"BuzzBox/service/article/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func PublishArticleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.PublishRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewPublishArticleLogic(r.Context(), svcCtx)
		resp, err := l.PublishArticle(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
