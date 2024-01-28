package svc

import (
	"github.com/sjxiang/mall/service/user/api/internal/config"
	"github.com/sjxiang/mall/service/user/api/internal/middleware"
	"github.com/sjxiang/mall/service/user/model"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

// 依赖倒置
type ServiceContext struct {
	Config    config.Config
	Cost      rest.Middleware
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	// 考虑切换成 gorm，✅

	conn := sqlx.NewMysql(c.Mysql.DataSource)
	
	return &ServiceContext{
		Config:    c,
		Cost:      middleware.NewCostMiddleware().Handle,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
