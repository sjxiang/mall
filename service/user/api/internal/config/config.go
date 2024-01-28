package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	
	Auth struct {  // jwt 鉴权配置
		AccessSecret string  // jwt 密钥
		AccessExpire int64   // 有效期，单位：秒
	}

	Mysql struct {  // 数据库配置，除 mysql 外，可能还有 mongo 等其他数据库
		DataSource string  // mysql 链接地址，满足 $user:$password@tcp($ip:$port)/$db?$queries 格式即可
	}

	CacheRedis cache.CacheConf
}
