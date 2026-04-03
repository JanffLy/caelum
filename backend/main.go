package main

import (
	"fmt"
	_ "caelum-backend/docs"
	"caelum-backend/controllers"
	"caelum-backend/core"
	"caelum-backend/middleware"
	"caelum-backend/models"

	beego "github.com/beego/beego/v2/server/web"
)

func main() {
	// 1. 初始化配置
	if err := core.InitConfig(); err != nil {
		panic("配置初始化失败: " + err.Error())
	}

	// 2. 初始化数据库
	if err := core.InitDatabase(); err != nil {
		panic("数据库初始化失败: " + err.Error())
	}

	// 3. 初始化Redis
	if err := core.InitRedis(); err != nil {
		panic("Redis初始化失败: " + err.Error())
	}

	// 4. 初始化错误处理器
	middleware.InitBeegoErrorHandler()

	// 5. 配置Beego
	beego.BConfig.RunMode = beego.DEV
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.BConfig.WebConfig.StaticDir["swagger"] = "swagger"

	// 6. 注册中间件
	// CORS中间件 - 应用于所有请求
	beego.InsertFilter("*", beego.BeforeRouter, middleware.CORSFilter())
	// JWT认证中间件
	beego.InsertFilter("/api/*", beego.BeforeRouter, middleware.JWTAuth())
	// 错误恢复中间件
	beego.InsertFilter("*", beego.BeforeExec, middleware.RecoveryFilter())

	// 7. 注册认证路由
	beego.Router("/api/v1/auth/login", &controllers.AuthController{}, "post:Login")
	beego.Router("/api/v1/auth/logout", &controllers.AuthController{}, "post:Logout")
	beego.Router("/api/v1/auth/register", &controllers.AuthController{}, "post:Register")
	beego.Router("/api/v1/auth/refresh", &controllers.AuthController{}, "post:RefreshToken")
	beego.Router("/api/v1/auth/user", &controllers.AuthController{}, "get:GetCurrentUser")
	beego.Router("/api/v1/auth/password", &controllers.AuthController{}, "put:ChangePassword")

	// 8. 注册部门路由
	beego.Router("/api/v1/depts", &controllers.DeptController{})
	beego.Router("/api/v1/depts/tree", &controllers.DeptController{}, "get:GetTree")
	beego.Router("/api/v1/depts/:id", &controllers.DeptController{})

	// 9. 初始化模型
	models.InitUserTable()
	models.InitDeptTable()

	// 10. 启动应用
	// 从配置文件读取端口
	port, _ := beego.AppConfig.Int("httpport")
	if port <= 0 {
		port = 8080
	}
	addr := fmt.Sprintf(":%d", port)
	
	// 打印启动信息
	printStartupInfo(port)
	
	beego.Run(addr)
}

func printStartupInfo(port int) {
	println("========================================")
	println("🚀 Caelum 后端服务启动成功")
	println("========================================")
	println(fmt.Sprintf("📍 服务地址: http://localhost:%d", port))
	println(fmt.Sprintf("📍 API文档:  http://localhost:%d/swagger", port))
	println(fmt.Sprintf("📍 健康检查: http://localhost:%d/health", port))
	println("========================================")
}