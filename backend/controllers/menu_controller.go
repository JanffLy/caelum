package controllers

// MenuController 菜单控制器
type MenuController struct {
	BaseController
}

// List 获取菜单列表
func (c *MenuController) List() {
	c.Success([]interface{}{}, "获取成功")
}

// Create 创建菜单
func (c *MenuController) Create() {
	c.Success(nil, "创建成功")
}

// Get 获取菜单详情
func (c *MenuController) Get() {
	c.Success(map[string]interface{}{"id": c.GetID()}, "获取成功")
}

// Update 更新菜单
func (c *MenuController) Update() {
	c.Success(nil, "更新成功")
}

// Delete 删除菜单
func (c *MenuController) Delete() {
	c.Success(nil, "删除成功")
}

// GetTree 获取菜单树
func (c *MenuController) GetTree() {
	c.Success([]interface{}{}, "获取成功")
}

// GetByRoleId 根据角色ID获取菜单
func (c *MenuController) GetByRoleId() {
	c.Success([]interface{}{}, "获取成功")
}