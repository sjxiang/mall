package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/sjxiang/mall/service/user/rpc/internal/config"
	"github.com/sjxiang/mall/service/user/rpc/internal/server"
	"github.com/sjxiang/mall/service/user/rpc/internal/svc"
	"github.com/sjxiang/mall/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"github.com/zeromicro/zero-contrib/zrpc/registry/consul"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var configFile = flag.String("f", "etc/user.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterUserServer(grpcServer, server.NewUserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	
	// 注册服务到 consul
	consul.RegisterService(c.ListenOn, c.Consul)
	defer s.Stop()

	// 注册 server 端拦截器
	s.AddUnaryInterceptors(func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
		
		// 取元数据
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Errorf(codes.InvalidArgument, "need metadata")
		}
		if md["token"][0] != "bearer xxx" {
			return nil, status.Errorf(codes.PermissionDenied, "invalid token")
		}

		resp, err = handler(ctx, req)  // 实际的 rpc 方法调用
		
		return
	})

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
