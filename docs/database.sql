-- Caelum 中后台管理系统数据库初始化脚本

-- 创建数据库
CREATE DATABASE IF NOT EXISTS caelum DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE caelum;

-- 用户表
CREATE TABLE IF NOT EXISTS sys_user (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    username VARCHAR(50) NOT NULL UNIQUE COMMENT '用户名',
    password VARCHAR(255) NOT NULL COMMENT '密码',
    nickname VARCHAR(50) COMMENT '昵称',
    email VARCHAR(100) COMMENT '邮箱',
    phone VARCHAR(20) COMMENT '手机号',
    avatar VARCHAR(255) COMMENT '头像',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    dept_id BIGINT COMMENT '部门ID',
    post_id BIGINT COMMENT '岗位ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_username (username),
    INDEX idx_dept_id (dept_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户表';

-- 部门表
CREATE TABLE IF NOT EXISTS sys_dept (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    parent_id BIGINT DEFAULT 0 COMMENT '父部门ID',
    dept_name VARCHAR(50) NOT NULL COMMENT '部门名称',
    dept_code VARCHAR(50) COMMENT '部门编码',
    sort INT DEFAULT 0 COMMENT '排序',
    leader VARCHAR(50) COMMENT '负责人',
    phone VARCHAR(20) COMMENT '联系电话',
    email VARCHAR(100) COMMENT '邮箱',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_parent_id (parent_id),
    INDEX idx_dept_name (dept_name)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='部门表';

-- 角色表
CREATE TABLE IF NOT EXISTS sys_role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    role_name VARCHAR(50) NOT NULL COMMENT '角色名称',
    role_code VARCHAR(50) NOT NULL UNIQUE COMMENT '角色编码',
    description VARCHAR(255) COMMENT '描述',
    sort INT DEFAULT 0 COMMENT '排序',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_role_code (role_code),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色表';

-- 菜单表
CREATE TABLE IF NOT EXISTS sys_menu (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    parent_id BIGINT DEFAULT 0 COMMENT '父菜单ID',
    menu_type VARCHAR(1) DEFAULT 'C' COMMENT '菜单类型: C-目录 M-菜单 B-按钮',
    menu_name VARCHAR(50) NOT NULL COMMENT '菜单名称',
    icon VARCHAR(50) COMMENT '图标',
    path VARCHAR(255) COMMENT '路由路径',
    component VARCHAR(255) COMMENT '组件路径',
    permission VARCHAR(100) COMMENT '权限标识',
    sort INT DEFAULT 0 COMMENT '排序',
    visible TINYINT DEFAULT 1 COMMENT '是否可见: 1-是 0-否',
    keep_alive TINYINT DEFAULT 1 COMMENT '是否缓存: 1-是 0-否',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_parent_id (parent_id),
    INDEX idx_menu_type (menu_type),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='菜单表';

-- 岗位表
CREATE TABLE IF NOT EXISTS sys_post (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    post_name VARCHAR(50) NOT NULL COMMENT '岗位名称',
    post_code VARCHAR(50) NOT NULL UNIQUE COMMENT '岗位编码',
    sort INT DEFAULT 0 COMMENT '排序',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_post_code (post_code),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='岗位表';

-- 字典类型表
CREATE TABLE IF NOT EXISTS sys_dict (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    dict_name VARCHAR(100) NOT NULL COMMENT '字典名称',
    dict_code VARCHAR(100) NOT NULL UNIQUE COMMENT '字典编码',
    description VARCHAR(255) COMMENT '描述',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_dict_code (dict_code),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典类型表';

-- 字典项表
CREATE TABLE IF NOT EXISTS sys_dict_item (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    dict_id BIGINT NOT NULL COMMENT '字典类型ID',
    item_text VARCHAR(100) NOT NULL COMMENT '字典项文本',
    item_value VARCHAR(100) NOT NULL COMMENT '字典项值',
    sort INT DEFAULT 0 COMMENT '排序',
    status TINYINT DEFAULT 1 COMMENT '状态: 1-正常 0-禁用',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    deleted_at DATETIME DEFAULT NULL,
    INDEX idx_dict_id (dict_id),
    INDEX idx_status (status)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='字典项表';

-- 用户-角色关联表
CREATE TABLE IF NOT EXISTS sys_user_role (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    user_id BIGINT NOT NULL COMMENT '用户ID',
    role_id BIGINT NOT NULL COMMENT '角色ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_user_role (user_id, role_id),
    INDEX idx_user_id (user_id),
    INDEX idx_role_id (role_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='用户-角色关联表';

-- 角色-菜单关联表
CREATE TABLE IF NOT EXISTS sys_role_menu (
    id BIGINT PRIMARY KEY AUTO_INCREMENT,
    role_id BIGINT NOT NULL COMMENT '角色ID',
    menu_id BIGINT NOT NULL COMMENT '菜单ID',
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    UNIQUE KEY uk_role_menu (role_id, menu_id),
    INDEX idx_role_id (role_id),
    INDEX idx_menu_id (menu_id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色-菜单关联表';

-- 初始化管理员账号 (密码: admin123)
INSERT INTO sys_user (username, password, nickname, email, status) 
VALUES ('admin', '$2a$10$N.zmdr9k7uOCQb376NoUnuTJ8iAt6Z5EHsM8lE9lBOsl7iAt.RNxu', '超级管理员', 'admin@caelum.com', 1);

-- 初始化顶级部门
INSERT INTO sys_dept (dept_name, dept_code, sort) VALUES ('Caelum集团', 'CAELUM', 1);

-- 初始化角色
INSERT INTO sys_role (role_name, role_code, description, sort) VALUES 
('超级管理员', 'super_admin', '拥有所有权限', 1),
('普通角色', 'common', '普通用户角色', 2);

-- 初始化菜单
INSERT INTO sys_menu (menu_name, menu_type, path, component, permission, sort) VALUES 
('系统管理', 'C', '/system', 'layouts/SystemLayout', '', 1),
('用户管理', 'C', '/system/user', 'system/user/index', 'system:user:list', 2),
('角色管理', 'C', '/system/role', 'system/role/index', 'system:role:list', 3),
('菜单管理', 'C', '/system/menu', 'system/menu/index', 'system:menu:list', 4),
('部门管理', 'C', '/system/dept', 'system/dept/index', 'system:dept:list', 5),
('岗位管理', 'C', '/system/post', 'system/post/index', 'system:post:list', 6),
('字典管理', 'C', '/system/dict', 'system/dict/index', 'system:dict:list', 7);

-- 给超级管理员分配所有菜单
INSERT INTO sys_role_menu (role_id, menu_id) 
SELECT 1, id FROM sys_menu;

-- 给管理员分配超级管理员角色
INSERT INTO sys_user_role (user_id, role_id) VALUES (1, 1);