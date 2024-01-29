package interceptor

import (
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

// CtxKey 自定义一个类型
type CtxKey string

const (
	// 使用自定义类型声明 context 中存储的 key，防止被他人覆盖
	CtxKeyAdminID CtxKey = "admin_id"
)

// 添加 client 端，拦截器
func Prepare(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
	// rpc 调用前
	
	adminID := ctx.Value(CtxKeyAdminID).(string)

	md := metadata.Pairs(
		"key1", "val-1",
		"key1", "val-2",  // "key1"的值将会是 []string{"val-1", "val-2"}
		"request_id", "17612345",
		"token", "bearer xxx",
		"admin_id", adminID,  // 从外部获取，借助 ctx 上下文
	)
	
	ctx = metadata.NewOutgoingContext(ctx, md)           // metadata 随 rpc 发送出去
	err := invoker(ctx, method, req, reply, cc, opts...) // 实际的 rpc 调用
	
	// rpc 调用后
	return err
}
