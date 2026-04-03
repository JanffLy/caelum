package controllers

// DictController 字典控制器
type DictController struct {
	BaseController
}

// List 获取字典类型列表
func (c *DictController) List() {
	c.Success([]interface{}{}, "获取成功")
}

// Create 创建字典类型
func (c *DictController) Create() {
	c.Success(nil, "创建成功")
}

// Get 获取字典类型详情
func (c *DictController) Get() {
	c.Success(map[string]interface{}{"id": c.GetID()}, "获取成功")
}

// Update 更新字典类型
func (c *DictController) Update() {
	c.Success(nil, "更新成功")
}

// Delete 删除字典类型
func (c *DictController) Delete() {
	c.Success(nil, "删除成功")
}

// GetItems 获取字典项
func (c *DictController) GetItems() {
	c.Success([]interface{}{}, "获取成功")
}

// DictItemController 字典项控制器
type DictItemController struct {
	BaseController
}

// List 获取字典项列表
func (c *DictItemController) List() {
	c.Success([]interface{}{}, "获取成功")
}

// Create 创建字典项
func (c *DictItemController) Create() {
	c.Success(nil, "创建成功")
}

// Get 获取字典项详情
func (c *DictItemController) Get() {
	c.Success(map[string]interface{}{"id": c.GetID()}, "获取成功")
}

// Update 更新字典项
func (c *DictItemController) Update() {
	c.Success(nil, "更新成功")
}

// Delete 删除字典项
func (c *DictItemController) Delete() {
	c.Success(nil, "删除成功")
}