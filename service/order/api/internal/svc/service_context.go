package svc

import (
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/sjxiang/mall/service/order/api/internal/config"
	"github.com/sjxiang/mall/service/order/api/internal/interceptor"
	userClient "github.com/sjxiang/mall/service/user/rpc/user"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  userClient.User     
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		// 初始化 user 服务的 rpc 客户端
		UserRPC: userClient.NewUser(zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptor.Prepare))),
	}
}
