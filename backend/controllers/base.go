package controllers

import (
	"encoding/json"

	beego "github.com/beego/beego/v2/server/web"
)

// BaseController 基础控制器
type BaseController struct {
	beego.Controller
}

// Index 首页
func (c *BaseController) Index() {
	c.Ctx.WriteString("Welcome to Caelum Backend API")
}

// Response 统一响应
func (c *BaseController) Response(code int, msg string, data interface{}) {
	c.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json")
	c.Data["json"] = map[string]interface{}{
		"code": code,
		"msg":  msg,
		"data": data,
	}
	c.ServeJSON()
}

// Success 成功响应
func (c *BaseController) Success(data interface{}, msg string) {
	if msg == "" {
		msg = "操作成功"
	}
	c.Response(200, msg, data)
}

// Error 错误响应
func (c *BaseController) Error(code int, msg string) {
	c.Response(code, msg, nil)
}

// GetJSONBody 解析JSON请求体
func (c *BaseController) GetJSONBody(v interface{}) error {
	body := c.Ctx.Request.Body
	return json.NewDecoder(body).Decode(v)
}

// GetPage 获取分页参数
func (c *BaseController) GetPage() (int, int) {
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("pageSize", 10)
	if page < 1 {
		page = 1
	}
	if pageSize < 1 {
		pageSize = 10
	}
	if pageSize > 100 {
		pageSize = 100
	}
	return page, pageSize
}

// GetID 获取URL参数ID
func (c *BaseController) GetID() int64 {
	id, _ := c.GetInt64(":id")
	return id
}

// HealthController 健康检查控制器
type HealthController struct {
	beego.Controller
}

// Check 健康检查
func (c *HealthController) Check() {
	c.Ctx.WriteString(`{"status":"ok","version":"1.0.0"}`)
}
