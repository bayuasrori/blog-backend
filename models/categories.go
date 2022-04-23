package models

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	gorm.Model
	ID          uint64     `gorm:"primaryKey" json:"id"`
	Name        string     `json:"name" gorm:"unique"`
	Description string     `gorm:"type:text" json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
	Articles    []*Article `gorm:"many2many:article_categories;"`
}

type ArticleCategory struct {
	gorm.Model
	CategoryID int `gorm:"primaryKey"`
	ArticleID  int `gorm:"primaryKey"`
	CreatedAt  time.Time
}

func GetCategories() []Category {
	db := Connect()
	categories := []Category{}
	db.Find(&categories)

	return categories
}

func GetCategoryName(name string) Category {
	db := Connect()
	category := Category{Name: name}
	db.Find(&category)

	return category
}

func CreateCategory(category Category) Category {
	db := Connect()
	db.Create(&category)

	return category
}

func DeleteCategory(id uint64) Category {
	db := Connect()
	category := Category{ID: id}
	db.Delete(&category)
	return category
}
