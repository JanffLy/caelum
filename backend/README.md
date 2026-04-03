# Caelum 后端服务

基于 Beego 2.0 框架开发的企业级中后台管理系统后端 API 服务。

## 技术栈

- **语言**: Go 1.21+
- **框架**: Beego 2.0
- **ORM**: Beego ORM
- **数据库**: MySQL 8.0
- **缓存**: Redis

## 项目结构

```
backend/
├── conf/                # 配置文件
│   └── app.conf        # Beego 主配置
├── controllers/        # 控制器层
│   ├── base.go         # 基础控制器
│   ├── auth_controller.go    # 认证模块
│   ├── user_controller.go    # 用户管理
│   ├── dept_controller.go    # 部门管理
│   ├── menu_controller.go    # 菜单管理
│   ├── role_controller.go    # 角色管理
│   ├── post_controller.go    # 岗位管理
│   └── dict_controller.go    # 字典管理
├── models/             # 模型层（待实现）
├── routers/            # 路由配置
│   └── router.go       # 路由注册
├── services/           # 业务逻辑层（待实现）
├── middleware/         # 中间件（待实现）
├── utils/              # 工具函数（待实现）
├── docs/               # Swagger 文档
├── main.go             # 入口文件
└── go.mod              # Go 模块
```

## 配置说明

配置文件位于 `conf/app.conf`，主要配置项：

```ini
# 应用配置
appname = caelum-backend
httpport = 8080
runmode = dev

# 数据库配置
db.type = mysql
db.host = 127.0.0.1
db.port = 3306
db.database = caelum
db.username = root
db.password = 
db.charset = utf8mb4

# Redis 配置
redis.host = 127.0.0.1
redis.port = 6379
redis.password = 
redis.db = 0

# JWT 配置
jwt.secret = caelum-secret-key-2026
jwt.expire = 720
```

## 快速开始

### 1. 初始化数据库

```bash
mysql -u root -p < ../docs/database.sql
```

### 2. 修改配置文件

编辑 `conf/app.conf`，配置正确的数据库和 Redis 连接信息。

### 3. 运行服务

```bash
# 开发模式
go run main.go

# 或运行编译后的可执行文件
./caelum-backend
```

服务将在 http://localhost:8080 启动。

## API 列表

### 认证模块
| 方法 | 路径 | 说明 |
|------|------|------|
| POST | /api/v1/auth/login | 用户登录 |
| POST | /api/v1/auth/logout | 用户登出 |
| POST | /api/v1/auth/register | 用户注册 |
| POST | /api/v1/auth/refresh | 刷新Token |

### 用户管理
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/users | 获取用户列表 |
| POST | /api/v1/users | 创建用户 |
| GET | /api/v1/users/:id | 获取用户详情 |
| PUT | /api/v1/users/:id | 更新用户 |
| DELETE | /api/v1/users/:id | 删除用户 |
| PUT | /api/v1/users/:id/reset-password | 重置密码 |
| PUT | /api/v1/users/:id/assign-roles | 分配角色 |
| PUT | /api/v1/users/:id/assign-dept | 分配部门 |
| PUT | /api/v1/users/:id/assign-post | 分配岗位 |

### 部门管理
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/depts | 获取部门列表 |
| POST | /api/v1/depts | 创建部门 |
| GET | /api/v1/depts/:id | 获取部门详情 |
| PUT | /api/v1/depts/:id | 更新部门 |
| DELETE | /api/v1/depts/:id | 删除部门 |
| GET | /api/v1/depts/tree | 获取部门树 |

### 菜单管理
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/menus | 获取菜单列表 |
| POST | /api/v1/menus | 创建菜单 |
| GET | /api/v1/menus/:id | 获取菜单详情 |
| PUT | /api/v1/menus/:id | 更新菜单 |
| DELETE | /api/v1/menus/:id | 删除菜单 |
| GET | /api/v1/menus/tree | 获取菜单树 |
| GET | /api/v1/menus/role/:roleId | 获取角色菜单 |

### 角色管理
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/roles | 获取角色列表 |
| POST | /api/v1/roles | 创建角色 |
| GET | /api/v1/roles/:id | 获取角色详情 |
| PUT | /api/v1/roles/:id | 更新角色 |
| DELETE | /api/v1/roles/:id | 删除角色 |
| GET | /api/v1/roles/:id/menus | 获取角色菜单 |
| PUT | /api/v1/roles/:id/menus | 分配菜单 |

### 岗位管理
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/posts | 获取岗位列表 |
| POST | /api/v1/posts | 创建岗位 |
| GET | /api/v1/posts/:id | 获取岗位详情 |
| PUT | /api/v1/posts/:id | 更新岗位 |
| DELETE | /api/v1/posts/:id | 删除岗位 |

### 字典管理
| 方法 | 路径 | 说明 |
|------|------|------|
| GET | /api/v1/dicts | 获取字典类型列表 |
| POST | /api/v1/dicts | 创建字典类型 |
| GET | /api/v1/dicts/:id | 获取字典类型详情 |
| PUT | /api/v1/dicts/:id | 更新字典类型 |
| DELETE | /api/v1/dicts/:id | 删除字典类型 |
| GET | /api/v1/dicts/:id/items | 获取字典项 |
| GET | /api/v1/dict-items | 获取字典项列表 |
| POST | /api/v1/dict-items | 创建字典项 |
| GET | /api/v1/dict-items/:id | 获取字典项详情 |
| PUT | /api/v1/dict-items/:id | 更新字典项 |
| DELETE | /api/v1/dict-items/:id | 删除字典项 |

## 统一响应格式

```json
{
  "code": 200,
  "msg": "操作成功",
  "data": {}
}
```

## 开发指南

1. 控制器位于 `controllers/` 目录，继承 `BaseController`
2. 业务逻辑建议放在 `services/` 目录
3. 数据模型定义在 `models/` 目录
4. 中间件放在 `middleware/` 目录
5. 工具函数放在 `utils/` 目录

## 许可证

MIT License