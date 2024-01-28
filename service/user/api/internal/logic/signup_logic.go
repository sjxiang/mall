package logic

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/sjxiang/mall/service/user/api/internal/svc"
	"github.com/sjxiang/mall/service/user/api/internal/types"
	"github.com/sjxiang/mall/service/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type SignupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSignupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SignupLogic {
	return &SignupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

var (
	
	ErrInternal           = errors.New("系统内部错误")
	ErrNoRecord           = errors.New("no matching record found，用户不存在")
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrDuplicateUsername  = errors.New("duplicate username，用户名已经存在")
)


// 业务逻辑
func (l *SignupLogic) Signup(req *types.SignupRequest) (resp *types.SignupResponse, err error) {

	err = validateSignupRequest(req)
	if err != nil {
		return nil, err
	}	

	// 把用户的注册信息保存到数据库中
	// 1. 查询 username 是否已经注册（事务？）
	
	u, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	// 1.1 查询数据库失败了
	if err != nil && err != model.ErrNotFound {
		logx.Errorw(
			"User_Signup_UserModel.FindOneByUsername failed", 
			logx.Field("err", err),
		)
		return nil, ErrInternal
	}

	// 1.2 查到记录，表示该用户名已经被注册
	if u != nil {
		return nil, errors.New("用户名已存在")
	}

	// 1.3 没查到记录，那就走流程
	// 2. 生成 userId（雪花算法）
	// 3. 密码哈希（加盐，md5）
	hashedPassword, err := HashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("加密失败：%w", err) 
	}

	user := &model.User{
		UserId:   time.Now().Unix(),  // 简化，设置为当前时间戳
		Username: req.Username,      
		Password: hashedPassword,     // 不能存明文
		Gender:   int64(req.Gender),
	}
	_, err = l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		logx.Errorf("User_Signup_UserModel.Insert failed, err: %v", err)
		return nil, err
	}
	
	return &types.SignupResponse{Message: "注册成功"}, nil 
}

func validateSignupRequest(req *types.SignupRequest) (err error) {
	logx.Debugv(req)  // json.Marshal(req)
	logx.Debugf("req: %#v", req)

	if req.Password != req.RePassword {
		return errors.New("两次输入的密码不一致")
	}

	return nil
}


func HashPassword(plainText string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)
	}

	return string(hashedPassword), nil
}


func encrypt(plainText string, salt string) string {
	h := md5.New()

	h.Write([]byte(plainText))  // 明文
	h.Write([]byte(salt))       // 加盐

	return hex.EncodeToString(h.Sum(nil)) // 编码
}