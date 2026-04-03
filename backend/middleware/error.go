package middleware

import (
	"fmt"
	"runtime/debug"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// ErrorHandler 全局错误处理中间件
func ErrorHandler() beego.FilterFunc {
	return func(ctx *context.Context) {
		// 包装panic处理
		defer func() {
			if err := recover(); err != nil {
				// 记录错误日志
				fmt.Printf("❌ [PANIC] %v\n", err)
				fmt.Printf("Stack: %s\n", debug.Stack())

				// 返回错误响应
				ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
				ctx.ResponseWriter.WriteHeader(500)
				ctx.WriteString(`{"code":500,"msg":"服务器内部错误","data":null}`)
			}
		}()
	}
}

// NotFoundHandler 404处理
func NotFoundHandler() beego.FilterFunc {
	return func(ctx *context.Context) {
		// 如果没有匹配的路由
		status := ctx.ResponseWriter.Status
		if status == 0 || status == 404 {
			ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
			ctx.WriteString(`{"code":404,"msg":"请求的资源不存在","data":null}`)
		}
	}
}

// RecoveryFilter 恢复过滤器
func RecoveryFilter() beego.FilterFunc {
	return func(ctx *context.Context) {
		defer func() {
			if err := recover(); err != nil {
				fmt.Printf("❌ [RECOVERY] %v\n", err)
				fmt.Printf("Stack: %s\n", debug.Stack())
				
				ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
				ctx.ResponseWriter.WriteHeader(500)
				ctx.WriteString(`{"code":500,"msg":"服务器内部错误","data":null}`)
			}
		}()
	}
}

// InitBeegoErrorHandler 初始化Beego错误处理器
func InitBeegoErrorHandler() {
	// 设置Panic处理函数
	beego.BConfig.RecoverFunc = func(ctx *context.Context, cfg *beego.Config) {
		if err := recover(); err != nil {
			fmt.Printf("❌ [BEEGO PANIC] %v\n", err)
			fmt.Printf("Stack: %s\n", debug.Stack())
		}
	}
}