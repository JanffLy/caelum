package middleware

import (
	"strings"

	"caelum-backend/utils"

	"github.com/beego/beego/v2/server/web/context"
)

// JWTAuth JWT认证中间件
func JWTAuth() func(ctx *context.Context) {
	return func(ctx *context.Context) {
		// 获取Token（从Header或Query）
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			token = ctx.Input.Query("token")
		}

		// 去掉Bearer前缀
		if len(token) > 7 && strings.ToUpper(token[:7]) == "BEARER " {
			token = token[7:]
		}

		// 公开接口不需要验证
		publicPaths := []string{
			"/api/v1/auth/login",
			"/api/v1/auth/register",
			"/health",
			"/",
		}
		
		isPublic := false
		for _, path := range publicPaths {
			if strings.HasPrefix(ctx.Request.URL.Path, path) {
				isPublic = true
				break
			}
		}

		if isPublic {
			return
		}

		// 验证Token
		if token == "" {
			ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.WriteString(`{"code":401,"msg":"未登录或Token已过期","data":null}`)
			return
		}

		// 解析Token
		claims, err := utils.ParseToken(token)
		if err != nil {
			ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
			ctx.ResponseWriter.WriteHeader(401)
			ctx.WriteString(`{"code":401,"msg":"Token无效或已过期","data":null}`)
			return
		}

		// 将用户信息存入Context，供后续使用
		ctx.Input.SetParam("user_id", int64ToString(claims.UserID))
		ctx.Input.SetParam("username", claims.Username)
		ctx.Input.SetParam("role_ids", int64SliceToString(claims.RoleIDs))
	}
}

func int64ToString(i int64) string {
	return string(rune(i + '0'))
}

func int64SliceToString(slice []int64) string {
	result := ""
	for _, v := range slice {
		result += string(rune(v + '0')) + ","
	}
	return result
}