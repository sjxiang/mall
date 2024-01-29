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
	Config    config.Config    // 配置
	Cost      rest.Middleware  // web 中间件，要与 api 文件中命名一致
	UserModel model.UserModel  // db
}

func NewServiceContext(c config.Config) *ServiceContext {
	
	// 考虑切换成 gorm，✅

	sqlxConn := sqlx.NewMysql(c.Mysql.DataSource)
	
	return &ServiceContext{
		Config:    c,
		Cost:      middleware.NewCostMiddleware().Handle,
		UserModel: model.NewUserModel(sqlxConn, c.CacheRedis),
	}
}
