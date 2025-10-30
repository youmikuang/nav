-- 创建数据库
CREATE DATABASE IF NOT EXISTS nav DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE nav;

-- 创建用户表
CREATE TABLE IF NOT EXISTS users (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  username VARCHAR(255) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password VARCHAR(255) NOT NULL,
  created_at BIGINT UNSIGNED,
  updated_at BIGINT UNSIGNED,
  PRIMARY KEY (id),
  KEY idx_username (username),
  KEY idx_email (email)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建导航分类表
CREATE TABLE IF NOT EXISTS nav_categories (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  sort_order INT DEFAULT 0,
  created_at BIGINT UNSIGNED,
  updated_at BIGINT UNSIGNED,
  PRIMARY KEY (id),
  KEY idx_sort_order (sort_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 创建导航链接表
CREATE TABLE IF NOT EXISTS nav_links (
  id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  category_id BIGINT UNSIGNED NOT NULL,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  icon VARCHAR(255),
  link VARCHAR(255) NOT NULL,
  sort_order INT DEFAULT 0,
  created_at BIGINT UNSIGNED,
  updated_at BIGINT UNSIGNED,
  PRIMARY KEY (id),
  FOREIGN KEY (category_id) REFERENCES nav_categories(id) ON DELETE CASCADE,
  KEY idx_category_id (category_id),
  KEY idx_sort_order (sort_order)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

-- 插入示例数据
INSERT INTO nav_categories (title, sort_order) VALUES
('个人网站', 1),
('常用工具', 2),
('React', 3);

INSERT INTO nav_links (category_id, title, description, icon, link, sort_order) VALUES
(1, 'TODO List', '一个简单好用的待办事项应用', 'https://todo.gaoheng.top/assets/favicon.png', 'https://todo.gaoheng.top/', 1),
(1, '个人简历', '个人在线简历', 'https://resume.gaoheng.top/avatar.jpg', 'https://resume.gaoheng.top/', 2),
(2, 'Can I use', '前端 API 兼容性查询', 'https://caniuse.com/img/favicon-128.png', 'https://caniuse.com', 1),
(2, 'TinyPNG', '在线图片压缩工具', 'https://tinypng.com/images/apple-touch-icon.png', 'https://tinypng.com', 2),
(3, 'React', '用于构建用户界面的 JavaScript 库', 'https://zh-hans.reactjs.org/favicon.ico', 'https://zh-hans.reactjs.org', 1),
(3, 'React Router', 'React 的声明式路由', 'https://reactrouter.com/favicon-light.png', 'https://reactrouter.com', 2);
