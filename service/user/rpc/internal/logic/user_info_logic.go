package logic

import (
	"context"
	"errors"
	"fmt"

	"github.com/sjxiang/mall/service/user/model"
	"github.com/sjxiang/mall/service/user/rpc/internal/svc"
	"github.com/sjxiang/mall/service/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *pb.UserInfoRequest) (*pb.UserInfoResp, error) {
	// 参数校验
	violations := validateUserInfoRequest(in)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}
	
	// 根据用户id查询数据库
	user, err := l.svcCtx.UserModel.FindOneByUserId(l.ctx, in.GetUserId())
	if errors.Is(err, model.ErrNotFound) {
		return nil, notFoundError("用户不存在")
	}
	if err != nil {
		logx.Errorw(
			"user_rpc.UserInfo.UserModel.FindOneByUserId failed", 
			logx.Field("err", err),
		)
		return nil, internalError(err)
	}

	// 返回响应
	return &pb.UserInfoResp{
		UserId:   user.UserId,
		Username: user.Username,
		Gender:   user.Gender,
	}, nil
}

func validateUserInfoRequest(req *pb.UserInfoRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := validateUserId(req.GetUserId()); err != nil {
		violations = append(violations, fieldViolation("user_id", err))
	}	
	
	return violations
}

func validateUserId(value int64) error {
	// 无效状态，命中
	if value <= 0 {
		// 提示，每与操反
		return fmt.Errorf("无效 user_id")
	}

	return nil 
}

