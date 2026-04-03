# Caelum 中后台管理系统

<p align="center">
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go" alt="Go">
  <img src="https://img.shields.io/badge/Vue-3.4+-4FC08D?style=for-the-badge&logo=vue.js" alt="Vue">
  <img src="https://img.shields.io/badge/Beego-2.0-00ADD8?style=for-the-badge" alt="Beego">
  <img src="https://img.shields.io/badge/Element Plus-2.0+-409EFF?style=for-the-badge" alt="Element Plus">
</p>

## 简介

Caelum 是一套基于 Go + Beego 后端和 Vue3 + Element Plus 前端开发的中后台管理系统，采用前后端分离架构，提供完整的权限管理和系统功能模块。

## 技术栈

### 后端
- **语言**: Go 1.21+
- **框架**: Beego 2.0
- **数据库**: MySQL 8.0
- **缓存**: Redis
- **认证**: JWT + OAuth 2.0

### 前端
- **框架**: Vue 3.4+ (组合式API)
- **UI组件库**: Element Plus
- **状态管理**: Pinia
- **构建工具**: Vite
- **语言**: TypeScript 5.0+

## 功能模块

- 🔐 **用户管理** - 用户CRUD、角色分配、部门分配、岗位分配
- 🏢 **部门管理** - 树形部门结构、级联选择
- 🔑 **角色管理** - 角色权限分配、菜单权限
- 📋 **菜单管理** - 动态菜单、按钮权限
- 💼 **岗位管理** - 岗位CRUD
- 📚 **字典管理** - 字典类型、字典项
- 🔒 **认证模块** - OAuth 2.0 登录、JWT Token

## 项目结构

```
caelum/
├── backend/           # Go后端 (Beego)
│   ├── conf/          # 配置文件
│   ├── controllers/   # 控制器层
│   ├── models/        # 模型层
│   ├── routers/       # 路由配置
│   ├── services/      # 业务逻辑层
│   ├── middleware/    # 中间件
│   ├── utils/         # 工具函数
│   └── main.go        # 入口文件
│
├── frontend/          # Vue3前端
│   ├── src/
│   │   ├── api/       # API接口
│   │   ├── components/# 公共组件
│   │   ├── router/    # 路由配置
│   │   ├── stores/    # 状态管理
│   │   ├── views/     # 页面组件
│   │   └── ...
│   └── package.json
│
├── docs/              # 文档
│   ├── database.sql   # 数据库脚本
│   └── TECHNICAL_PLAN.md # 技术方案
│
└── README.md          # 项目说明
```

## 快速开始

### 环境要求

- Go 1.21+
- Node.js 18+
- MySQL 8.0+
- Redis 6.0+

### 后端启动

```bash
# 进入后端目录
cd backend

# 初始化数据库
mysql -u root -p < ../docs/database.sql

# 运行服务
go run main.go
# 或
./caelum-backend
```

服务将在 http://localhost:8080 启动

### 前端启动

```bash
# 进入前端目录
cd frontend

# 安装依赖
npm install

# 启动开发服务器
npm run dev
```

应用将在 http://localhost:5173 启动

### 默认账号

- 用户名: admin
- 密码: admin123

## API 文档

启动服务后访问: http://localhost:8080/swagger

## 开发指南

详细开发指南请参考 [技术方案文档](docs/TECHNICAL_PLAN.md)

## 许可证

MIT License