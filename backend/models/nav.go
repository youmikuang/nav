package models

// NavCategory 前端展示模型
type NavCategory struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"not null"`
	SortOrder int       `json:"-" gorm:"default:0;column:sort_order"`
	CreatedAt int64     `json:"-" gorm:"column:created_at"`
	UpdatedAt int64     `json:"-" gorm:"column:updated_at"`
	Items     []NavLink `json:"items,omitempty" gorm:"foreignKey:CategoryID"`
}

// NavLink 前端展示模型
type NavLink struct {
	ID         uint   `json:"-" gorm:"primaryKey"`
	CategoryID uint   `json:"-" gorm:"not null;column:category_id"`
	Icon       string `json:"icon" gorm:"column:icon"`
	Title      string `json:"title" gorm:"not null;column:title"`
	Desc       string `json:"desc" gorm:"column:description"` // 数据库是description，JSON是desc
	Link       string `json:"link" gorm:"not null;column:link"`
	SortOrder  int    `json:"-" gorm:"default:0;column:sort_order"`
	CreatedAt  int64  `json:"-" gorm:"column:created_at"`
	UpdatedAt  int64  `json:"-" gorm:"column:updated_at"`
}

// 指定表名
func (NavCategory) TableName() string {
	return "nav_categories"
}

func (NavLink) TableName() string {
	return "nav_links"
}

// Admin 请求结构体（用于管理接口，包含完整字段）
type CreateCategoryRequest struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateCategoryRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	SortOrder   *int   `json:"sort_order"`
}

type CreateLinkRequest struct {
	CategoryID  uint   `json:"category_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Link        string `json:"link" binding:"required"`
	SortOrder   int    `json:"sort_order"`
}

type UpdateLinkRequest struct {
	CategoryID  *uint  `json:"category_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	Link        string `json:"link"`
	SortOrder   *int   `json:"sort_order"`
}
