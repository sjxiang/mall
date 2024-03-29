package handler

import (
	"net/http"

	"github.com/sjxiang/mall/service/user/api/internal/logic"
	"github.com/sjxiang/mall/service/user/api/internal/svc"
	"github.com/sjxiang/mall/service/user/api/internal/types"
	
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/go-playground/validator/v10"
)

func DetailHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DetailRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		if err := validator.New().StructCtx(r.Context(), &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDetailLogic(r.Context(), svcCtx)
		resp, err := l.Detail(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
