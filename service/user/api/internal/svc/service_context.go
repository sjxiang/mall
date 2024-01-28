package svc

import (
	"github.com/sjxiang/mall/service/user/api/internal/config"
	"github.com/sjxiang/mall/service/user/api/internal/middleware"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config config.Config
	Cost   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Cost:   middleware.NewCostMiddleware().Handle,
	}
}
