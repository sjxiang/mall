package svc

import (
	"github.com/sjxiang/mall/service/order/api/internal/config"
	userClient "github.com/sjxiang/mall/service/user/rpc/user"
	
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config   config.Config
	UserRPC  userClient.User     
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		UserRPC: userClient.NewUser(zrpc.MustNewClient(c.UserRPC)),
	}
}

