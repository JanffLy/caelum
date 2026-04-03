package middleware

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

// CORSFilter CORS跨域中间件
func CORSFilter() beego.FilterFunc {
	return func(ctx *context.Context) {
		// 允许的来源
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Origin", "*")
		// 允许的方法
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的头部
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		// 允许凭证
		ctx.ResponseWriter.Header().Set("Access-Control-Allow-Credentials", "true")
		// 暴露的头部
		ctx.ResponseWriter.Header().Set("Access-Control-Expose-Headers", "Content-Length, Content-Type")

		// 处理预检请求
		if ctx.Request.Method == "OPTIONS" {
			ctx.ResponseWriter.WriteHeader(204)
			return
		}
	}
}
