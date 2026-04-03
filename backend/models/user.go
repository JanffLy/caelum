package models

import (
	"time"

	"github.com/beego/beego/v2/client/orm"
)

// User 用户模型
type User struct {
	ID        int64     `orm:"column(id);auto;pk" json:"id"`
	Username  string    `orm:"column(username);size(50);unique" json:"username"`
	Password  string    `orm:"column(password);size(255)" json:"-"`
	Nickname  string    `orm:"column(nickname);size(50);null" json:"nickname"`
	Email     string    `orm:"column(email);size(100);null" json:"email"`
	Phone     string    `orm:"column(phone);size(20);null" json:"phone"`
	Avatar    string    `orm:"column(avatar);size(255);null" json:"avatar"`
	Status    int       `orm:"column(status);default(1)" json:"status"` // 1-正常 0-禁用
	DeptID    int64     `orm:"column(dept_id);null" json:"dept_id"`
	PostID    int64     `orm:"column(post_id);null" json:"post_id"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
	UpdatedAt time.Time `orm:"column(updated_at);auto_now;type(datetime)" json:"updated_at"`
	DeletedAt time.Time `orm:"column(deleted_at);null;type(datetime)" json:"deleted_at"`
}

// TableName 表名
func (u *User) TableName() string {
	return "sys_user"
}

// UserRole 用户-角色关联模型
type UserRole struct {
	ID        int64     `orm:"column(id);auto;pk" json:"id"`
	UserID    int64     `orm:"column(user_id)" json:"user_id"`
	RoleID    int64     `orm:"column(role_id)" json:"role_id"`
	CreatedAt time.Time `orm:"column(created_at);auto_now_add;type(datetime)" json:"created_at"`
}

// TableName 表名
func (ur *UserRole) TableName() string {
	return "sys_user_role"
}

// UserQueryParam 用户查询参数
type UserQueryParam struct {
	Username  string
	Nickname  string
	Status    int
	DeptID    int64
	PostID    int64
	Page      int
	PageSize  int
}

// UserView 用户视图对象
type UserView struct {
	ID        int64     `json:"id"`
	Username  string    `json:"username"`
	Nickname  string    `json:"nickname"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Avatar    string    `json:"avatar"`
	Status    int       `json:"status"`
	DeptID    int64     `json:"dept_id"`
	DeptName  string    `json:"dept_name"`
	PostID    int64     `json:"post_id"`
	PostName  string    `json:"post_name"`
	CreatedAt time.Time `json:"created_at"`
}

// InitUserTable 初始化用户表
func InitUserTable() error {
	// 注册模型
	orm.RegisterModel(new(User), new(UserRole))
	return nil
}

// CreateUser 创建用户
func CreateUser(user *User) (int64, error) {
	o := orm.NewOrm()
	return o.Insert(user)
}

// GetUserByID 根据ID获取用户
func GetUserByID(id int64) (*User, error) {
	o := orm.NewOrm()
	user := &User{ID: id}
	err := o.Read(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(username string) (*User, error) {
	o := orm.NewOrm()
	user := &User{Username: username}
	err := o.Read(user, "username")
	if err != nil {
		return nil, err
	}
	return user, nil
}

// UpdateUser 更新用户
func UpdateUser(user *User) error {
	o := orm.NewOrm()
	_, err := o.Update(user)
	return err
}

// DeleteUser 删除用户
func DeleteUser(id int64) error {
	o := orm.NewOrm()
	_, err := o.Delete(&User{ID: id})
	return err
}

// GetUserList 获取用户列表
func GetUserList(param UserQueryParam) ([]UserView, int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(User)).Filter("deleted_at__isnull", true)

	// 构建查询条件
	if param.Username != "" {
		qs = qs.Filter("username__contains", param.Username)
	}
	if param.Nickname != "" {
		qs = qs.Filter("nickname__contains", param.Nickname)
	}
	if param.Status > 0 {
		qs = qs.Filter("status", param.Status)
	}
	if param.DeptID > 0 {
		qs = qs.Filter("dept_id", param.DeptID)
	}
	if param.PostID > 0 {
		qs = qs.Filter("post_id", param.PostID)
	}

	// 总数
	total, _ := qs.Count()

	// 分页
	if param.Page <= 0 {
		param.Page = 1
	}
	if param.PageSize <= 0 {
		param.PageSize = 10
	}
	qs = qs.Limit(param.PageSize, (param.Page-1)*param.PageSize)

	// 查询
	var users []User
	_, err := qs.All(&users)
	if err != nil {
		return nil, 0, err
	}

	// 转换为视图对象
	var views []UserView
	for _, user := range users {
		views = append(views, UserView{
			ID:        user.ID,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Email:     user.Email,
			Phone:     user.Phone,
			Avatar:    user.Avatar,
			Status:    user.Status,
			DeptID:    user.DeptID,
			PostID:    user.PostID,
			CreatedAt: user.CreatedAt,
		})
	}

	return views, total, nil
}

// AssignRoles 分配角色
func AssignRoles(userID int64, roleIDs []int64) error {
	o := orm.NewOrm()

	// 删除原有角色关联
	_, err := o.QueryTable(new(UserRole)).Filter("user_id", userID).Delete()
	if err != nil {
		return err
	}

	// 添加新的角色关联
	for _, roleID := range roleIDs {
		userRole := &UserRole{
			UserID: userID,
			RoleID: roleID,
		}
		_, err := o.Insert(userRole)
		if err != nil {
			return err
		}
	}

	return nil
}

// GetUserRoles 获取用户角色ID列表
func GetUserRoles(userID int64) ([]int64, error) {
	o := orm.NewOrm()
	var userRoles []UserRole
	_, err := o.QueryTable(new(UserRole)).Filter("user_id", userID).All(&userRoles)
	if err != nil {
		return nil, err
	}

	var roleIDs []int64
	for _, ur := range userRoles {
		roleIDs = append(roleIDs, ur.RoleID)
	}
	return roleIDs, nil
}

// ChangePassword 修改密码
func ChangePassword(userID int64, newPassword string) error {
	o := orm.NewOrm()
	user := &User{ID: userID}
	if err := o.Read(user); err != nil {
		return err
	}
	user.Password = newPassword
	_, err := o.Update(user, "password")
	return err
}