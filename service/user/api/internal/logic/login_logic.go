package logic

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/sjxiang/mall/service/user/api/internal/svc"
	"github.com/sjxiang/mall/service/user/api/internal/types"
	"github.com/sjxiang/mall/service/user/model"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	
	// 1. 处理用户发来的请求，拿到用户名和密码
	// req.Username req.Password

	err = validateLoginRequest(req)
	if err != nil {
		return nil, err
	}

	// 2. 判断输入的用户名和密码，跟数据库中的是不是一致的
	// 两种方式：
	// 1.用 用户输入的用户名和密码（加密后）去查数据库
	// select * from user where username=req.Username and password=req.Password
	// 2.用用户名查到结果，再判断密码
	
	user, err := l.svcCtx.UserModel.FindOneByUsername(l.ctx, req.Username)
	if err == model.ErrNotFound {
		return &types.LoginResponse{Message: "用户名不存在"}, nil
	}
	if err != nil {
		logx.Errorw(
			"UserModel.FindOneByUsername failed", 
			logx.Field("err", err),
		)
		return nil, ErrInternal
	}

	err = CheckPassword(user.Password, req.Password)
	if err != nil {
		if err == ErrInvalidCredentials {
			return &types.LoginResponse{
				Message: "用户名或密码错误",
			}, nil	
		}

		return nil, ErrInternal
	}

	// 生成 token 
	var (
		now       = time.Now().Unix()
		expire    = l.svcCtx.Config.Auth.AccessExpire
		secretKey = l.svcCtx.Config.Auth.AccessSecret
	)

	token, err := genToken(secretKey, nil, now, expire,user.UserId)
	if err != nil {
		logx.Errorw(
			"genToken failed", 
			logx.Field("err", err),
		)
		return nil, ErrInternal
	}
	return &types.LoginResponse{
		Message:     "登录成功",
		AccessToken: token,
	}, nil 
}

func validateLoginRequest(req *types.LoginRequest) (err error) {	
	if len(req.Username) == 0 {
		return errors.New("用户名不为空")
	}
	if err := validatePassword(req.Password); err != nil {
		return err
	}
	
	return nil
}

func validatePassword(password string) error {
	return validateString(password, 6, 32)
}

func validateString(value string, minLen, maxLen int) error {
	n := len(value)
	if n < minLen || n > maxLen {
		return fmt.Errorf("字符长度必须在 %d 到 %d 之间", minLen, maxLen)
	}

	return nil 
}

func CheckPassword(hashedPassword, plainText string) (err error) {
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainText))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return ErrInvalidCredentials
		} else {
			return err
		}
	}
	return nil
}


// 生成 JWT token
func genToken(secretKey string, payloads map[string]interface{}, iat, seconds, userId int64) (string, error) {

	claims := make(jwt.MapClaims)
	claims["iat"] = iat            // 当前时间戳
	claims["exp"] = iat + seconds  // 截至日期时间戳
	claims["user_id"] = userId     

	// 其余补充
	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}

