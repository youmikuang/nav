package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"nav/backend/db"
	"nav/backend/routes"
)

func main() {
	// 加载 .env 文件
	_ = godotenv.Load()

	// 初始化数据库
	db.Init()

	// 创建 Gin 路由
	router := gin.Default()

	// 设置信任的代理
	// 在开发环境中，只信任 localhost
	// 在生产环境中，应该设置为实际的代理 IP
	router.SetTrustedProxies([]string{"127.0.0.1", "::1"})

	// 设置路由
	routes.SetupRoutes(router)

	// 启动服务器
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	addr := fmt.Sprintf(":%s", port)
	log.Printf("Server started on %s", addr)

	if err := router.Run(addr); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
