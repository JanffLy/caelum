package controllers

// RoleController 角色控制器
type RoleController struct {
	BaseController
}

// List 获取角色列表
func (c *RoleController) List() {
	c.Success([]interface{}{}, "获取成功")
}

// Create 创建角色
func (c *RoleController) Create() {
	c.Success(nil, "创建成功")
}

// Get 获取角色详情
func (c *RoleController) Get() {
	c.Success(map[string]interface{}{"id": c.GetID()}, "获取成功")
}

// Update 更新角色
func (c *RoleController) Update() {
	c.Success(nil, "更新成功")
}

// Delete 删除角色
func (c *RoleController) Delete() {
	c.Success(nil, "删除成功")
}

// GetMenus 获取角色菜单
func (c *RoleController) GetMenus() {
	c.Success([]interface{}{}, "获取成功")
}

// AssignMenus 分配菜单
func (c *RoleController) AssignMenus() {
	c.Success(nil, "菜单分配成功")
}