package routers

import (
	"caelum-backend/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func InitRouters() {
	// 基础路由
	beego.Router("/", &controllers.BaseController{}, "get:Index")

	// 健康检查
	beego.Router("/health", &controllers.HealthController{}, "get:Check")

	// API路由组
	api := beego.NewNamespace("/api/v1",
		// 认证模块
		beego.NSRouter("/auth/login", &controllers.AuthController{}, "post:Login"),
		beego.NSRouter("/auth/logout", &controllers.AuthController{}, "post:Logout"),
		beego.NSRouter("/auth/register", &controllers.AuthController{}, "post:Register"),
		beego.NSRouter("/auth/refresh", &controllers.AuthController{}, "post:RefreshToken"),

		// 用户管理
		beego.NSRouter("/users", &controllers.UserController{}, "get:List;post:Create"),
		beego.NSRouter("/users/:id", &controllers.UserController{}, "get:Get;put:Update;delete:Delete"),
		beego.NSRouter("/users/:id/reset-password", &controllers.UserController{}, "put:ResetPassword"),
		beego.NSRouter("/users/:id/assign-roles", &controllers.UserController{}, "put:AssignRoles"),
		beego.NSRouter("/users/:id/assign-dept", &controllers.UserController{}, "put:AssignDept"),
		beego.NSRouter("/users/:id/assign-post", &controllers.UserController{}, "put:AssignPost"),

		// 部门管理
		beego.NSRouter("/depts", &controllers.DeptController{}, "get:List;post:Create"),
		beego.NSRouter("/depts/:id", &controllers.DeptController{}, "get:Get;put:Update;delete:Delete"),
		beego.NSRouter("/depts/tree", &controllers.DeptController{}, "get:GetTree"),

		// 菜单管理
		beego.NSRouter("/menus", &controllers.MenuController{}, "get:List;post:Create"),
		beego.NSRouter("/menus/:id", &controllers.MenuController{}, "get:Get;put:Update;delete:Delete"),
		beego.NSRouter("/menus/tree", &controllers.MenuController{}, "get:GetTree"),
		beego.NSRouter("/menus/role/:roleId", &controllers.MenuController{}, "get:GetByRoleId"),

		// 角色管理
		beego.NSRouter("/roles", &controllers.RoleController{}, "get:List;post:Create"),
		beego.NSRouter("/roles/:id", &controllers.RoleController{}, "get:Get;put:Update;delete:Delete"),
		beego.NSRouter("/roles/:id/menus", &controllers.RoleController{}, "get:GetMenus;put:AssignMenus"),

		// 岗位管理
		beego.NSRouter("/posts", &controllers.PostController{}, "get:List;post:Create"),
		beego.NSRouter("/posts/:id", &controllers.PostController{}, "get:Get;put:Update;delete:Delete"),

		// 字典管理
		beego.NSRouter("/dicts", &controllers.DictController{}, "get:List;post:Create"),
		beego.NSRouter("/dicts/:id", &controllers.DictController{}, "get:Get;put:Update;delete:Delete"),
		beego.NSRouter("/dicts/:id/items", &controllers.DictController{}, "get:GetItems"),
		beego.NSRouter("/dict-items", &controllers.DictItemController{}, "get:List;post:Create"),
		beego.NSRouter("/dict-items/:id", &controllers.DictItemController{}, "get:Get;put:Update;delete:Delete"),
	)

	// 注册命名空间
	beego.AddNamespace(api)

	// 启用自动路由
	beego.AutoRouter(&controllers.UserController{})
}