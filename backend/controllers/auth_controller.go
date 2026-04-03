package controllers

import (
	"caelum-backend/models"
	"caelum-backend/utils"

	beego "github.com/beego/beego/v2/server/web"
)

// LoginRequest 登录请求
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest 注册请求
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
}

// TokenResponse Token响应
type TokenResponse struct {
	Token        string `json:"token"`
	RefreshToken string `json:"refresh_token"`
	Expire       int    `json:"expire"`
}

// AuthController 认证控制器
type AuthController struct {
	BaseController
}

// Login 登录
func (c *AuthController) Login() {
	var req LoginRequest
	if err := c.GetJSONBody(&req); err != nil {
		c.Error(400, "请求参数错误")
		return
	}

	// 参数验证
	if req.Username == "" || req.Password == "" {
		c.Error(400, "用户名和密码不能为空")
		return
	}

	// 获取用户
	user, err := models.GetUserByUsername(req.Username)
	if err != nil {
		c.Error(401, "用户名或密码错误")
		return
	}

	// 验证密码
	if !utils.CheckPassword(req.Password, user.Password) {
		c.Error(401, "用户名或密码错误")
		return
	}

	// 检查用户状态
	if user.Status == 0 {
		c.Error(403, "账号已被禁用")
		return
	}

	// 获取用户角色
	roleIDs, _ := models.GetUserRoles(user.ID)

	// 生成Token
	token, err := utils.GenerateToken(user.ID, user.Username, roleIDs)
	if err != nil {
		c.Error(500, "生成Token失败")
		return
	}

	// 生成刷新Token
	refreshToken, err := utils.RefreshToken(token)
	if err != nil {
		c.Error(500, "生成刷新Token失败")
		return
	}

	// 返回响应
	c.Success(TokenResponse{
		Token:        token,
		RefreshToken: refreshToken,
		Expire:       24 * 3600, // 24小时
	}, "登录成功")
}

// Logout 登出
func (c *AuthController) Logout() {
	// 从Header获取Token
	token := c.Ctx.Request.Header.Get("Authorization")
	if token != "" {
		// 可以在这里将Token加入黑名单（需要Redis支持）
		// 目前简化处理，直接返回成功
	}

	c.Success(nil, "登出成功")
}

// Register 注册
func (c *AuthController) Register() {
	var req RegisterRequest
	if err := c.GetJSONBody(&req); err != nil {
		c.Error(400, "请求参数错误")
		return
	}

	// 参数验证
	if req.Username == "" || req.Password == "" {
		c.Error(400, "用户名和密码不能为空")
		return
	}

	// 检查用户名是否已存在
	_, err := models.GetUserByUsername(req.Username)
	if err == nil {
		c.Error(400, "用户名已存在")
		return
	}

	// 密码加密
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.Error(500, "密码加密失败")
		return
	}

	// 创建用户
	user := &models.User{
		Username: req.Username,
		Password: hashedPassword,
		Nickname: req.Nickname,
		Email:    req.Email,
		Status:   1,
	}

	_, err = models.CreateUser(user)
	if err != nil {
		c.Error(500, "创建用户失败")
		return
	}

	c.Success(nil, "注册成功")
}

// RefreshToken 刷新Token
func (c *AuthController) RefreshToken() {
	// 从Header获取Token
	token := c.Ctx.Request.Header.Get("Authorization")
	if token == "" {
		c.Error(401, "Token不能为空")
		return
	}

	// 解析Token（去掉Bearer前缀）
	if len(token) > 7 && token[:7] == "Bearer " {
		token = token[7:]
	}

	// 刷新Token
	newToken, err := utils.RefreshToken(token)
	if err != nil {
		c.Error(401, "Token刷新失败")
		return
	}

	c.Success(TokenResponse{
		Token:        newToken,
		RefreshToken: "",
		Expire:       24 * 3600,
	}, "Token刷新成功")
}

// GetCurrentUser 获取当前用户信息
func (c *AuthController) GetCurrentUser() {
	// 从Context获取当前用户（由中间件设置）
	userID, _ := c.GetInt64("user_id")

	if userID == 0 {
		c.Error(401, "未登录")
		return
	}

	user, err := models.GetUserByID(userID)
	if err != nil {
		c.Error(404, "用户不存在")
		return
	}

	c.Success(map[string]interface{}{
		"id":       user.ID,
		"username": user.Username,
		"nickname": user.Nickname,
		"email":    user.Email,
		"phone":    user.Phone,
		"avatar":   user.Avatar,
	}, "获取成功")
}

// ChangePassword 修改密码
func (c *AuthController) ChangePassword() {
	userID, _ := c.GetInt64("user_id")
	
	var req struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	
	if err := c.GetJSONBody(&req); err != nil {
		c.Error(400, "请求参数错误")
		return
	}

	if req.OldPassword == "" || req.NewPassword == "" {
		c.Error(400, "密码不能为空")
		return
	}

	// 获取用户
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.Error(404, "用户不存在")
		return
	}

	// 验证旧密码
	if !utils.CheckPassword(req.OldPassword, user.Password) {
		c.Error(400, "原密码错误")
		return
	}

	// 加密新密码
	hashedPassword, err := utils.HashPassword(req.NewPassword)
	if err != nil {
		c.Error(500, "密码加密失败")
		return
	}

	// 更新密码
	err = models.ChangePassword(userID, hashedPassword)
	if err != nil {
		c.Error(500, "修改密码失败")
		return
	}

	c.Success(nil, "密码修改成功")
}

// InitAuth 初始化认证相关表
func InitAuth() {
	models.InitUserTable()
}

// InitBeegoAuthController 初始化Beego控制器
func init() {
	beego.Router("/api/v1/auth/login", &AuthController{}, "post:Login")
	beego.Router("/api/v1/auth/logout", &AuthController{}, "post:Logout")
	beego.Router("/api/v1/auth/register", &AuthController{}, "post:Register")
	beego.Router("/api/v1/auth/refresh", &AuthController{}, "post:RefreshToken")
	beego.Router("/api/v1/auth/user", &AuthController{}, "get:GetCurrentUser")
	beego.Router("/api/v1/auth/password", &AuthController{}, "put:ChangePassword")
}