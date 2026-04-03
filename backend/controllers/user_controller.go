package controllers

// UserController 用户控制器
type UserController struct {
	BaseController
}

// List 获取用户列表
func (c *UserController) List() {
	c.Success([]interface{}{}, "获取成功")
}

// Create 创建用户
func (c *UserController) Create() {
	c.Success(nil, "创建成功")
}

// Get 获取用户详情
func (c *UserController) Get() {
	c.Success(map[string]interface{}{"id": c.GetID()}, "获取成功")
}

// Update 更新用户
func (c *UserController) Update() {
	c.Success(nil, "更新成功")
}

// Delete 删除用户
func (c *UserController) Delete() {
	c.Success(nil, "删除成功")
}

// ResetPassword 重置密码
func (c *UserController) ResetPassword() {
	c.Success(nil, "密码重置成功")
}

// AssignRoles 分配角色
func (c *UserController) AssignRoles() {
	c.Success(nil, "角色分配成功")
}

// AssignDept 分配部门
func (c *UserController) AssignDept() {
	c.Success(nil, "部门分配成功")
}

// AssignPost 分配岗位
func (c *UserController) AssignPost() {
	c.Success(nil, "岗位分配成功")
}