package routes

import (
	"github.com/gin-gonic/gin"
	"nav/backend/handlers"
	"nav/backend/middleware"
)

func SetupRoutes(router *gin.Engine) {
	// CORS 中间件
	router.Use(CORSMiddleware())

	// API 组
	api := router.Group("/api")

	// 认证相关路由（无需认证）
	auth := api.Group("/auth")
	{
		auth.POST("/register", handlers.Register)
		auth.POST("/login", handlers.Login)
		auth.POST("/logout", handlers.Logout)
	}

	// 导航相关路由（公开访问）
	nav := api.Group("/nav")
	{
		// 分类
		nav.GET("/categories", handlers.GetCategories)
		nav.GET("/categories/:id", handlers.GetCategory)

		// 链接
		nav.GET("/links", handlers.GetLinks)
		nav.GET("/links/:id", handlers.GetLink)
	}

	// 需要认证的路由
	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())
	{
		// 用户相关
		protected.GET("/auth/me", handlers.GetCurrentUser)

		// 导航管理（需要认证）
		navAdmin := protected.Group("/nav")
		{
			// 分类管理
			navAdmin.POST("/categories", handlers.CreateCategory)
			navAdmin.PUT("/categories/:id", handlers.UpdateCategory)
			navAdmin.DELETE("/categories/:id", handlers.DeleteCategory)

			// 链接管理
			navAdmin.POST("/links", handlers.CreateLink)
			navAdmin.PUT("/links/:id", handlers.UpdateLink)
			navAdmin.DELETE("/links/:id", handlers.DeleteLink)
		}
	}
}

// CORS 中间件
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.Request.Header.Get("Origin")

		// 允许的源列表
		allowedOrigins := map[string]bool{
			"http://localhost:5173":  true,
			"http://localhost:3000":  true,
			"http://localhost:8080":  true,
			"http://127.0.0.1:5173":  true,
			"http://127.0.0.1:3000":  true,
			"http://127.0.0.1:8080":  true,
		}

		// 如果请求来自允许的源，则设置该源
		if allowedOrigins[origin] {
			c.Writer.Header().Set("Access-Control-Allow-Origin", origin)
			c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		} else if origin == "" {
			// 如果没有 Origin 头（比如同源请求或某些工具），允许所有
			c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		}

		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}
