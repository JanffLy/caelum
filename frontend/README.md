# Caelum 前端应用

基于 Vue 3 + Element Plus 开发的中后台管理系统前端应用。

## 技术栈

- **框架**: Vue 3.4+ (组合式 API)
- **UI 组件库**: Element Plus
- **状态管理**: Pinia
- **路由**: Vue Router 4
- **构建工具**: Vite
- **语言**: TypeScript 5.0+
- **HTTP 客户端**: Axios

## 项目结构

```
frontend/
├── public/                 # 静态资源
├── src/
│   ├── api/                # API 接口
│   │   └── ...
│   ├── assets/             # 资源文件
│   │   ├── styles/         # 全局样式
│   │   └── ...
│   ├── components/        # 公共组件
│   │   ├── common/         # 通用组件
│   │   └── ...
│   ├── directive/          # 自定义指令
│   │   └── ...
│   ├── layouts/            # 布局组件
│   │   └── ...
│   ├── router/             # 路由配置
│   │   └── index.ts
│   ├── stores/             # 状态管理
│   │   ├── auth.ts         # 认证状态
│   │   ├── user.ts         # 用户状态
│   │   └── ...
│   ├── types/              # TypeScript 类型
│   │   └── ...
│   ├── utils/              # 工具函数
│   │   ├── request.ts      # axios 封装
│   │   ├── auth.ts         # 认证工具
│   │   └── ...
│   ├── views/              # 页面组件
│   │   ├── login/          # 登录页
│   │   ├── dashboard/      # 仪表盘
│   │   ├── system/         # 系统管理
│   │   │   ├── user/       # 用户管理
│   │   │   ├── role/       # 角色管理
│   │   │   ├── menu/       # 菜单管理
│   │   │   ├── dept/       # 部门管理
│   │   │   ├── post/       # 岗位管理
│   │   │   └── dict/       # 字典管理
│   │   └── ...
│   ├── App.vue             # 根组件
│   └── main.ts             # 入口文件
├── index.html              # HTML 模板
├── package.json            # 项目依赖
├── tsconfig.json           # TypeScript 配置
├── vite.config.ts          # Vite 配置
└── env.d.ts                # 环境变量类型
```

## 快速开始

### 安装依赖

```bash
npm install
```

### 开发模式

```bash
npm run dev
```

应用将在 http://localhost:5173 启动。

### 构建生产版本

```bash
npm run build
```

### 代码检查

```bash
npm run lint
```

## 功能模块

### 认证模块

- 用户登录
- 用户登出
- Token 刷新
- 权限验证

### 系统管理

- **用户管理**: 用户列表、创建用户、编辑用户、删除用户、分配角色、分配部门、分配岗位
- **角色管理**: 角色列表、创建角色、编辑角色、删除角色、分配菜单权限
- **菜单管理**: 菜单列表、创建菜单、编辑菜单、删除菜单、菜单树
- **部门管理**: 部门列表、创建部门、编辑部门、删除部门、部门树
- **岗位管理**: 岗位列表、创建岗位、编辑岗位、删除岗位
- **字典管理**: 字典类型管理、字典项管理

## 页面路由

| 路径         | 名称     | 说明     |
| ------------ | -------- | -------- |
| /login       | 登录页   | 用户登录 |
| /            | 首页     | 仪表盘   |
| /system/user | 用户管理 | 用户CRUD |
| /system/role | 角色管理 | 角色CRUD |
| /system/menu | 菜单管理 | 菜单CRUD |
| /system/dept | 部门管理 | 部门CRUD |
| /system/post | 岗位管理 | 岗位CRUD |
| /system/dict | 字典管理 | 字典CRUD |

## 开发指南

### 创建新页面

1. 在 `src/views/` 下创建页面目录
2. 创建 `index.vue` 页面组件
3. 在 `src/router/index.ts` 中配置路由
4. 在菜单管理中添加对应菜单

### API 接口封装

在 `src/api/` 目录下创建对应的 API 文件：

```typescript
// src/api/user.ts
import request from "@/utils/request";
import type { User, UserQuery } from "@/types/user";

export function getUserList(params: UserQuery) {
  return request({
    url: "/api/v1/users",
    method: "get",
    params,
  });
}

export function createUser(data: User) {
  return request({
    url: "/api/v1/users",
    method: "post",
    data,
  });
}
```

### 状态管理

使用 Pinia 管理全局状态：

```typescript
// src/stores/user.ts
import { defineStore } from "pinia";
import { ref } from "vue";

export const useUserStore = defineStore("user", () => {
  const token = ref("");
  const userInfo = ref(null);

  function setToken(newToken: string) {
    token.value = newToken;
  }

  return {
    token,
    userInfo,
    setToken,
  };
});
```

## 环境变量

| 变量名            | 说明         | 默认值                |
| ----------------- | ------------ | --------------------- |
| VITE_API_BASE_URL | API 基础路径 | http://localhost:8080 |

## 浏览器支持

- Chrome >= 90
- Firefox >= 90
- Safari >= 14
- Edge >= 90

## 许可证

MIT License
