package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"nav/backend/db"
	"nav/backend/models"
	"nav/backend/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "missing authorization header"})
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid authorization header"})
			c.Abort()
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}

		// 获取用户信息
		var user models.User
		if err := db.DB.First(&user, claims.UserID).Error; err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "user not found"})
			c.Abort()
			return
		}

		// 设置用户信息到 context
		c.Set("user", &user)
		c.Set("userID", user.ID)
		c.Next()
	}
}

// GetCurrentUser 获取当前用户
func GetCurrentUser(c *gin.Context) *models.User {
	user, _ := c.Get("user")
	return user.(*models.User)
}

// GetCurrentUserID 获取当前用户ID
func GetCurrentUserID(c *gin.Context) uint {
	userID, _ := c.Get("userID")
	return userID.(uint)
}
