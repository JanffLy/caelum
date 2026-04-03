package controllers

import (
	"caelum-backend/models"

	"github.com/beego/beego/v2/server/web"
)

// DeptController 部门控制器
type DeptController struct {
	BaseController
}

// URLMapping URL映射
func (c *DeptController) URLMapping() {
	c.Mapping("List", c.List)
	c.Mapping("GetTree", c.GetTree)
	c.Mapping("Get", c.Get)
	c.Mapping("Create", c.Create)
	c.Mapping("Update", c.Update)
	c.Mapping("Delete", c.Delete)
}

// List 获取部门列表
// @Title 获取部门列表
// @Description 获取部门列表
// @Param	dept_name	query	string	false	"部门名称"
// @Param	leader		query	string	false	"负责人"
// @Param	status		query	int		false	"状态"
// @Param	page		query	int		false	"页码"
// @Param	page_size	query	int		false	"每页数量"
// @Success 200 {object} controllers.Resp
// @router / [get]
func (c *DeptController) List() {
	// 获取查询参数
	deptName := c.GetString("dept_name")
	leader := c.GetString("leader")
	status, _ := c.GetInt("status")
	page, _ := c.GetInt("page", 1)
	pageSize, _ := c.GetInt("page_size", 10)

	param := models.DeptQueryParam{
		DeptName: deptName,
		Leader:   leader,
		Status:   status,
	}

	depts, total, err := models.GetDeptList(param)
	if err != nil {
		c.Error(500, "获取部门列表失败")
		return
	}

	// 构建分页响应
	pageData := map[string]interface{}{
		"list":      depts,
		"total":     total,
		"page":      page,
		"page_size": pageSize,
	}

	c.Success(pageData, "获取成功")
}

// GetTree 获取部门树形结构
// @Title 获取部门树
// @Description 获取部门树形结构
// @Success 200 {object} controllers.Resp
// @router /tree [get]
func (c *DeptController) GetTree() {
	tree, err := models.GetDeptTree()
	if err != nil {
		c.Error(500, "获取部门树失败")
		return
	}

	c.Success(tree, "获取成功")
}

// Get 获取单个部门
// @Title 获取部门
// @Description 根据ID获取部门详情
// @Param	id		path	int	true	"部门ID"
// @Success 200 {object} controllers.Resp
// @router /:id [get]
func (c *DeptController) Get() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.Error(400, "无效的部门ID")
		return
	}

	dept, err := models.GetDeptByID(id)
	if err != nil {
		c.Error(404, "部门不存在")
		return
	}

	c.Success(dept, "获取成功")
}

// Create 创建部门
// @Title 创建部门
// @Description 创建新部门
// @Param	body	body	models.Dept	true	"部门信息"
// @Success 200 {object} controllers.Resp
// @router / [post]
func (c *DeptController) Create() {
	var dept models.Dept
	if err := c.GetJSONBody(&dept); err != nil {
		c.Error(400, "请求参数错误")
		return
	}

	// 参数验证
	if dept.DeptName == "" {
		c.Error(400, "部门名称不能为空")
		return
	}

	// 设置默认值
	if dept.Status == 0 {
		dept.Status = 1
	}

	id, err := models.CreateDept(&dept)
	if err != nil {
		c.Error(500, "创建部门失败")
		return
	}

	dept.ID = id
	c.Success(dept, "创建成功")
}

// Update 更新部门
// @Title 更新部门
// @Description 更新部门信息
// @Param	id		path	int			true	"部门ID"
// @Param	body	body	models.Dept	true	"部门信息"
// @Success 200 {object} controllers.Resp
// @router /:id [put]
func (c *DeptController) Update() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.Error(400, "无效的部门ID")
		return
	}

	var dept models.Dept
	if err := c.GetJSONBody(&dept); err != nil {
		c.Error(400, "请求参数错误")
		return
	}

	// 参数验证
	if dept.DeptName == "" {
		c.Error(400, "部门名称不能为空")
		return
	}

	// 检查部门是否存在
	existing, err := models.GetDeptByID(id)
	if err != nil {
		c.Error(404, "部门不存在")
		return
	}

	// 更新字段
	existing.DeptName = dept.DeptName
	existing.ParentID = dept.ParentID
	existing.Sort = dept.Sort
	existing.Leader = dept.Leader
	existing.Phone = dept.Phone
	existing.Email = dept.Email
	existing.Status = dept.Status

	err = models.UpdateDept(existing)
	if err != nil {
		c.Error(500, "更新部门失败")
		return
	}

	c.Success(existing, "更新成功")
}

// Delete 删除部门
// @Title 删除部门
// @Description 删除部门
// @Param	id		path	int	true	"部门ID"
// @Success 200 {object} controllers.Resp
// @router /:id [delete]
func (c *DeptController) Delete() {
	id, err := c.GetInt64(":id")
	if err != nil {
		c.Error(400, "无效的部门ID")
		return
	}

	err = models.DeleteDept(id)
	if err != nil {
		c.Error(500, err.Error())
		return
	}

	c.Success(nil, "删除成功")
}

// InitDeptController 初始化部门控制器路由
func InitDeptController() {
	web.Router("/api/v1/depts", &DeptController{})
	web.Router("/api/v1/depts/tree", &DeptController{}, "get:GetTree")
	web.Router("/api/v1/depts/:id", &DeptController{})
}