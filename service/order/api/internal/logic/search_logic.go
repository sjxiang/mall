package logic

import (
	"context"
	"errors"

	"github.com/sjxiang/mall/service/order/api/internal/svc"
	"github.com/sjxiang/mall/service/order/api/internal/types"
	"github.com/sjxiang/mall/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchLogic {
	return &SearchLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchLogic) Search(req *types.SearchRequest) (resp *types.SearchResponse, err error) {

	// 1. 根据请求参数中的订单号查询数据库找到订单记录
	// fake

	// 2. 根据订单记录中的 user_id 去查询用户数据（通过RPC调用user服务）
	// fake 
	user, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &pb.UserInfoRequest{UserId: 1706537330})
	if err != nil {
		logx.Errorw(
			"UserRPC.GetUser failed", 
			logx.Field("err", err),
		)
		return nil, errors.New("系统内部错误")
	}

	// 3. 拼接返回结果（因为我们这个接口的数据不是由我一个服务组成的）
	return &types.SearchResponse{
		OrderID:  "fake",
		Status:   100,
		Username: user.GetUsername(),
	}, nil 
}

