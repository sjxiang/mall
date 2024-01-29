package svc

import (
	"github.com/sjxiang/mall/service/user/model"
	"github.com/sjxiang/mall/service/user/rpc/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config    config.Config
	UserModel model.UserModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	
	conn := sqlx.NewMysql(c.Mysql.DataSource)
	
	return &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, c.CacheRedis),
	}
}
