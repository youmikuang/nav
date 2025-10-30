#!/bin/bash

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

echo -e "${GREEN}=== 前端导航网站启动脚本 ===${NC}\n"

# 检查 Node.js 和 npm
if ! command -v node &> /dev/null; then
    echo -e "${RED}❌ 未找到 Node.js，请先安装 Node.js${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Node.js 版本:${NC} $(node -v)"
echo -e "${GREEN}✓ npm 版本:${NC} $(npm -v)"

# 检查 Go
if ! command -v go &> /dev/null; then
    echo -e "${RED}❌ 未找到 Go，请先安装 Go${NC}"
    exit 1
fi

echo -e "${GREEN}✓ Go 版本:${NC} $(go version)"

# 检查 MySQL
if ! command -v mysql &> /dev/null; then
    echo -e "${YELLOW}⚠ 未找到 MySQL 命令行工具，但可以尝试连接${NC}"
else
    echo -e "${GREEN}✓ MySQL 已安装${NC}"
fi

echo ""
echo -e "${YELLOW}选择要启动的部分：${NC}"
echo "1) 启动前端和后端"
echo "2) 仅启动前端"
echo "3) 仅启动后端"
echo "4) 初始化数据库"
echo "5) 退出"

read -p "请输入选项 (1-5): " choice

case $choice in
    1)
        echo -e "\n${YELLOW}启动前端和后端...${NC}\n"

        # 启动前端
        echo -e "${YELLOW}启动前端服务器...${NC}"
        cd homepage
        if [ ! -d "node_modules" ]; then
            echo -e "${YELLOW}安装前端依赖...${NC}"
            npm install
        fi
        npm run dev &
        FRONTEND_PID=$!
        cd ..

        echo -e "${GREEN}前端服务器启动成功 (PID: $FRONTEND_PID)${NC}"
        echo -e "${YELLOW}前端地址: http://localhost:5173${NC}\n"

        # 启动后端
        echo -e "${YELLOW}启动后端服务器...${NC}"
        cd backend
        go mod download
        go run main.go &
        BACKEND_PID=$!
        cd ..

        echo -e "${GREEN}后端服务器启动成功 (PID: $BACKEND_PID)${NC}"
        echo -e "${YELLOW}后端地址: http://localhost:8080${NC}\n"

        echo -e "${GREEN}✓ 所有服务已启动${NC}"
        echo -e "${YELLOW}按 Ctrl+C 停止所有服务${NC}"

        # 等待用户中断
        wait
        ;;

    2)
        echo -e "\n${YELLOW}启动前端服务器...${NC}\n"
        cd homepage
        if [ ! -d "node_modules" ]; then
            echo -e "${YELLOW}安装前端依赖...${NC}"
            npm install
        fi
        npm run dev
        ;;

    3)
        echo -e "\n${YELLOW}启动后端服务器...${NC}\n"
        cd backend

        # 创建 .env 文件如果不存在
        if [ ! -f ".env" ]; then
            echo -e "${YELLOW}创建 .env 文件...${NC}"
            cp .env.example .env
            echo -e "${YELLOW}请编辑 .env 文件，配置数据库连接${NC}"
        fi

        go mod download
        go run main.go
        ;;

    4)
        echo -e "\n${YELLOW}初始化数据库...${NC}\n"

        read -p "请输入 MySQL 用户名 (默认: root): " db_user
        db_user=${db_user:-root}

        read -s -p "请输入 MySQL 密码: " db_password
        echo ""

        # 执行数据库初始化脚本
        if [ -z "$db_password" ]; then
            mysql -u "$db_user" < backend/database.sql
        else
            mysql -u "$db_user" -p"$db_password" < backend/database.sql
        fi

        if [ $? -eq 0 ]; then
            echo -e "\n${GREEN}✓ 数据库初始化成功${NC}"
        else
            echo -e "\n${RED}❌ 数据库初始化失败${NC}"
            exit 1
        fi
        ;;

    5)
        echo -e "\n${YELLOW}退出${NC}"
        exit 0
        ;;

    *)
        echo -e "${RED}❌ 无效的选项${NC}"
        exit 1
        ;;
esac
