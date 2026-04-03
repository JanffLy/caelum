package controllers

// PostController 岗位控制器
type PostController struct {
	BaseController
}

// List 获取岗位列表
func (c *PostController) List() {
	c.Success([]interface{}{}, "获取成功")
}

// Create 创建岗位
func (c *PostController) Create() {
	c.Success(nil, "创建成功")
}

// Get 获取岗位详情
func (c *PostController) Get() {
	c.Success(map[string]interface{}{"id": c.GetID()}, "获取成功")
}

// Update 更新岗位
func (c *PostController) Update() {
	c.Success(nil, "更新成功")
}

// Delete 删除岗位
func (c *PostController) Delete() {
	c.Success(nil, "删除成功")
}