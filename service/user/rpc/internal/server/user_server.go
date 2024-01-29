// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"github.com/sjxiang/mall/service/user/rpc/internal/logic"
	"github.com/sjxiang/mall/service/user/rpc/internal/svc"
	"github.com/sjxiang/mall/service/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

func (s *UserServer) UserInfo(ctx context.Context, in *pb.UserInfoRequest) (*pb.UserInfoResp, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}
