#!/bin/bash

# Next.js 前端部署脚本
set -e  # 遇到错误立即退出

# 配置变量
SERVER="root@140.143.139.59"  # 替换为你的服务器 IP
DIR="/www/www.gaoheng.top"       # 替换为你的部署目录
APP_NAME="nav"               # 应用名称

# 颜色输出
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
RED='\033[0;31m'
NC='\033[0m' # No Color

echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}开始部署 ${APP_NAME}${NC}"
echo -e "${GREEN}================================${NC}"

# 1. 安装依赖并构建
echo -e "${YELLOW}[1/4] 安装依赖并构建...${NC}"
yarn
yarn build
echo -e "${GREEN}✓ 构建完成${NC}"

# 2. 初始化服务器目录
echo -e "${YELLOW}[2/4] 初始化服务器目录...${NC}"
ssh ${SERVER} "mkdir -p ${DIR}"
echo -e "${GREEN}✓ 目录初始化完成${NC}"

# 3. 打包并上传
echo -e "${YELLOW}[3/4] 打包并上传文件到服务器...${NC}"
tar --exclude='.git' \
    --exclude='node_modules' \
    --exclude='.next/cache' \
    -czf ${APP_NAME}.tar.gz .

scp ${APP_NAME}.tar.gz ${SERVER}:${DIR}/
rm ${APP_NAME}.tar.gz

# 4. 在服务器上部署
echo -e "${YELLOW}[4/4] 在服务器上部署...${NC}"
ssh ${SERVER} "cd ${DIR} && \
    tar -xzf ${APP_NAME}.tar.gz && \
    pnpm install --prod && \
    pm2 restart ${APP_NAME} || pm2 start 'pnpm start' --name ${APP_NAME}"

echo -e "${GREEN}✓ 部署完成${NC}"

echo -e "${GREEN}================================${NC}"
echo -e "${GREEN}部署成功！${NC}"
echo -e "${GREEN}================================${NC}"
echo ""
echo "提示："
echo "  - 使用 'pm2 logs ${APP_NAME}' 查看应用日志"
echo "  - 使用 'pm2 monit' 监控应用状态"
echo "  - 如果遇到问题，检查 Nginx 配置和防火墙设置"