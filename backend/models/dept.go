package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// Dept 部门模型
type Dept struct {
	ID        int64     `orm:"column(id);auto;pk" json:"id"`
	ParentID  int64     `orm:"column(parent_id);default(0)" json:"parent_id"` // 父部门ID，0为顶级
	DeptName  string    `orm:"column(dept_name);size(100)" json:"dept_name"`   // 部门名称
	Sort      int       `orm:"column(sort);default(0)" json:"sort"`           // 排序
	Leader    string    `orm:"column(leader);size(50);null" json:"leader"`     // 负责人
	Phone     string    `orm:"column(phone);size(20);null" json:"phone"`       // 联系电话
	Email     string    `orm:"column(email);size(100);null" json:"email"`      // 邮箱
	Status    int       `orm:"column(status);default(1)" json:"status"`       // 1-正常 0-禁用
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
	DeletedAt time.Time `orm:"column(deleted_at);null;type(datetime)" json:"deleted_at"`
}

// TableName 表名
func (d *Dept) TableName() string {
	return "sys_dept"
}

// DeptQueryParam 部门查询参数
type DeptQueryParam struct {
	DeptName string
	Leader   string
	Status   int
	ParentID int64
}

// DeptTreeNode 部门树节点
type DeptTreeNode struct {
	ID        int64        `json:"id"`
	ParentID  int64        `json:"parent_id"`
	DeptName  string       `json:"dept_name"`
	Sort      int          `json:"sort"`
	Leader    string       `json:"leader"`
	Phone     string       `json:"phone"`
	Email     string       `json:"email"`
	Status    int          `json:"status"`
	Children  []DeptTreeNode `json:"children"`
}

// InitDeptTable 初始化部门表
func InitDeptTable() error {
	orm.RegisterModel(new(Dept))
	return nil
}

// CreateDept 创建部门
func CreateDept(dept *Dept) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(dept)
}

// GetDeptByID 根据ID获取部门
func GetDeptByID(id int64) (*Dept, error) {
	o := orm.NewOrm()
	dept := &Dept{ID: id}
	err := o.Read(dept)
	if err != nil {
		return nil, err
	}
	return dept, nil
}

// UpdateDept 更新部门
func UpdateDept(dept *Dept) error {
	o := orm.NewOrm()
	_, err := o.Update(dept)
	return err
}

// DeleteDept 删除部门
func DeleteDept(id int64) error {
	o := orm.NewOrm()
	
	// 检查是否有子部门
	var count int64
	qs := o.QueryTable(new(Dept)).Filter("parent_id", id).Filter("deleted_at__isnull", true)
	count, _ = qs.Count()
	if count > 0 {
		return fmt.Errorf("该部门下存在子部门，无法删除")
	}
	
	// 检查是否有关联用户
	qsUser := o.QueryTable(new(User)).Filter("dept_id", id).Filter("deleted_at__isnull", true)
	count, _ = qsUser.Count()
	if count > 0 {
		return fmt.Errorf("该部门下存在用户，无法删除")
	}
	
	_, err := o.Delete(&Dept{ID: id})
	return err
}

// GetDeptList 获取部门列表
func GetDeptList(param DeptQueryParam) ([]Dept, int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Dept)).Filter("deleted_at__isnull", true)

	// 构建查询条件
	if param.DeptName != "" {
		qs = qs.Filter("dept_name__contains", param.DeptName)
	}
	if param.Leader != "" {
		qs = qs.Filter("leader__contains", param.Leader)
	}
	if param.Status > 0 {
		qs = qs.Filter("status", param.Status)
	}
	if param.ParentID >= 0 {
		qs = qs.Filter("parent_id", param.ParentID)
	}

	// 总数
	total, _ := qs.Count()

	// 排序查询
	qs = qs.OrderBy("sort", "id")

	var depts []Dept
	_, err := qs.All(&depts)
	if err != nil {
		return nil, 0, err
	}

	return depts, total, nil
}

// GetDeptTree 获取部门树形结构
func GetDeptTree() ([]DeptTreeNode, error) {
	o := orm.NewOrm()
	var depts []Dept
	_, err := o.QueryTable(new(Dept)).Filter("deleted_at__isnull", true).OrderBy("sort", "id").All(&depts)
	if err != nil {
		return nil, err
	}

	// 构建树形结构
	return buildDeptTree(depts, 0), nil
}

// buildDeptTree 递归构建部门树
func buildDeptTree(depts []Dept, parentID int64) []DeptTreeNode {
	var tree []DeptTreeNode
	for _, dept := range depts {
		if dept.ParentID == parentID {
			node := DeptTreeNode{
				ID:       dept.ID,
				ParentID: dept.ParentID,
				DeptName: dept.DeptName,
				Sort:     dept.Sort,
				Leader:   dept.Leader,
				Phone:    dept.Phone,
				Email:    dept.Email,
				Status:   dept.Status,
			}
			// 递归获取子部门
			node.Children = buildDeptTree(depts, dept.ID)
			tree = append(tree, node)
		}
	}
	return tree
}

// GetDeptPath 获取部门路径（用于面包屑）
func GetDeptPath(deptID int64) ([]Dept, error) {
	o := orm.NewOrm()
	var path []Dept
	
	currentID := deptID
	for currentID != 0 {
		dept := &Dept{ID: currentID}
		err := o.Read(dept)
		if err != nil {
			return nil, err
		}
		path = append([]Dept{*dept}, path...)
		currentID = dept.ParentID
	}
	
	return path, nil
}

// GetSubDeptIDs 获取所有子部门ID
func GetSubDeptIDs(parentID int64) ([]int64, error) {
	o := orm.NewOrm()
	var depts []Dept
	_, err := o.QueryTable(new(Dept)).Filter("parent_id", parentID).Filter("deleted_at__isnull", true).All(&depts)
	if err != nil {
		return nil, err
	}

	var ids []int64
	for _, dept := range depts {
		ids = append(ids, dept.ID)
		// 递归获取子部门的子部门
		subIDs, err := GetSubDeptIDs(dept.ID)
		if err != nil {
			return nil, err
		}
		ids = append(ids, subIDs...)
	}

	return ids, nil
}