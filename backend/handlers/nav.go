package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"nav/backend/db"
	"nav/backend/models"
)

// GetCategories 获取所有分类（包含链接）- 前端展示用
func GetCategories(c *gin.Context) {
	var categories []models.NavCategory

	// 预加载链接，并按 sort_order 排序
	if err := db.DB.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		Order("sort_order ASC").
		Find(&categories).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch categories"})
		return
	}

	c.JSON(http.StatusOK, categories)
}

// GetCategory 获取单个分类 - 前端展示用
func GetCategory(c *gin.Context) {
	id := c.Param("id")
	var category models.NavCategory

	if err := db.DB.
		Preload("Items", func(db *gorm.DB) *gorm.DB {
			return db.Order("sort_order ASC")
		}).
		First(&category, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, category)
}

// CreateCategory 创建分类
func CreateCategory(c *gin.Context) {
	var req models.CreateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 直接使用数据库原生字段插入
	if err := db.DB.Exec(
		"INSERT INTO nav_categories (title, description, sort_order) VALUES (?, ?, ?)",
		req.Title, req.Description, req.SortOrder,
	).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Category created successfully"})
}

// UpdateCategory 更新分类
func UpdateCategory(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateCategoryRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.SortOrder != nil {
		updates["sort_order"] = *req.SortOrder
	}

	result := db.DB.Table("nav_categories").Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category updated successfully"})
}

// DeleteCategory 删除分类
func DeleteCategory(c *gin.Context) {
	id := c.Param("id")

	result := db.DB.Delete(&models.NavCategory{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete category"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Category deleted successfully"})
}

// GetLinks 获取所有链接
func GetLinks(c *gin.Context) {
	var links []models.NavLink

	// 可选的分类过滤
	categoryID := c.Query("category_id")
	query := db.DB.Order("sort_order ASC")

	if categoryID != "" {
		query = query.Where("category_id = ?", categoryID)
	}

	if err := query.Find(&links).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch links"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": links})
}

// GetLink 获取单个链接
func GetLink(c *gin.Context) {
	id := c.Param("id")
	var link models.NavLink

	if err := db.DB.First(&link, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": link})
}

// CreateLink 创建链接
func CreateLink(c *gin.Context) {
	var req models.CreateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 验证分类是否存在
	var count int64
	if err := db.DB.Model(&models.NavCategory{}).Where("id = ?", req.CategoryID).Count(&count).Error; err != nil || count == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
		return
	}

	// 直接使用数据库原生字段插入
	if err := db.DB.Exec(
		"INSERT INTO nav_links (category_id, title, description, icon, link, sort_order) VALUES (?, ?, ?, ?, ?, ?)",
		req.CategoryID, req.Title, req.Description, req.Icon, req.Link, req.SortOrder,
	).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create link"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Link created successfully"})
}

// UpdateLink 更新链接
func UpdateLink(c *gin.Context) {
	id := c.Param("id")

	var req models.UpdateLinkRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updates := make(map[string]interface{})
	if req.CategoryID != nil {
		// 验证分类是否存在
		var count int64
		if err := db.DB.Model(&models.NavCategory{}).Where("id = ?", *req.CategoryID).Count(&count).Error; err != nil || count == 0 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category not found"})
			return
		}
		updates["category_id"] = *req.CategoryID
	}
	if req.Title != "" {
		updates["title"] = req.Title
	}
	if req.Description != "" {
		updates["description"] = req.Description
	}
	if req.Icon != "" {
		updates["icon"] = req.Icon
	}
	if req.Link != "" {
		updates["link"] = req.Link
	}
	if req.SortOrder != nil {
		updates["sort_order"] = *req.SortOrder
	}

	result := db.DB.Model(&models.NavLink{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update link"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link updated successfully"})
}

// DeleteLink 删除链接
func DeleteLink(c *gin.Context) {
	id := c.Param("id")

	result := db.DB.Delete(&models.NavLink{}, id)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete link"})
		return
	}
	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, gin.H{"error": "Link not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Link deleted successfully"})
}
