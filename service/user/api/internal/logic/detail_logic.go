package logic

import (
	"context"
	"encoding/json"
	"errors"

	"github.com/sjxiang/mall/service/user/api/internal/svc"
	"github.com/sjxiang/mall/service/user/api/internal/types"
	"github.com/sjxiang/mall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail(req *types.DetailRequest) (resp *types.DetailResponse, err error) {
	
	// jwt鉴权后，如何获取解析出来的数据
	logx.Debugv(l.ctx.Value("user_id"))

	currentUserId, err := l.ctx.Value("user_id").(json.Number).Int64()
	if err != nil {
		return nil, errors.New("类型断言 int64 失败")
	}
	logx.Debug("当前用户 id：", currentUserId)


	// 1. 拿到请求参数

	// 2. 根据用户id查询数据库
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, req.UserID)
	if err != nil {
		if err != model.ErrNotFound {
			logx.Errorw("User_Detail_UserModel.FindOneByUserId failed", logx.Field("err", err))
			return nil, ErrInternal
		}
		
		return nil, ErrNoRecord
	}

	// 3. 格式化数据(数据库里存的字段和前端要求的字段不太一致)
	// 4. 返回响应
	return &types.DetailResponse{
		Username: user.Username,
		Gender:   int(user.Gender),
	}, nil
}
