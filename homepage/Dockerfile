# Build stage
FROM node:20-alpine AS builder

WORKDIR /app

# 复制 package.json 和 yarn.lock
COPY package.json yarn.lock ./

# 安装依赖
RUN yarn install --frozen-lockfile

# 复制源代码
COPY . .

# 构建
RUN yarn build

# Production stage
FROM node:20-alpine

WORKDIR /app

# 安装 serve 来提供静态文件
RUN npm install -g serve

# 从 builder 阶段复制构建产物
COPY --from=builder /app/dist ./dist

EXPOSE 5173

CMD ["serve", "-s", "dist", "-l", "5173"]
