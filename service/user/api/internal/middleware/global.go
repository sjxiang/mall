package middleware

import (
	"bytes"
	"fmt"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

// 定义全局中间件
// feature 记录所有请求的响应信息

// rest.Middleware
// type Middleware func(next http.HandlerFunc) http.HandlerFunc
// type HandlerFunc func(ResponseWriter, *Request)

// 拷贝一份响应
type responseWithRecorder struct {
	http.ResponseWriter                // 结构体嵌入接口类型
	body                *bytes.Buffer  // 用来记录响应体内容
}

func NewResponseWithRecorder() *responseWithRecorder {
	return &responseWithRecorder{}
}

// 满足 http.ResponseWriter 接口类型
func (rr responseWithRecorder) Write(b []byte) (int, error) {
	// 1. 先拷贝响应内容
	rr.body.Write(b)
	// 2. 再往HTTP响应里写响应内容
	return rr.ResponseWriter.Write(b)
}

// CopyResp 复制请求的响应体
func (rr *responseWithRecorder) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 构造
		nw := responseWithRecorder{
			ResponseWriter: w,
			body:           bytes.NewBuffer([]byte{}),
		}

		// 处理业务逻辑
		next(nw, r)

		// 打印日志
		fmt.Printf("请求路径，%v；请求响应，%v\n", r.URL, nw.body.String())

	}
}

func MiddlewareWithAnotherService(ok bool) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if ok {
				fmt.Println("ok!")
			}
			next(w, r)
		}
	}
}
