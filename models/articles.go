package models

import (
	"time"

	"gorm.io/gorm"

	"github.com/gosimple/slug"
)

type Article struct {
	gorm.Model
	ID        uint      `gorm:"primaryKey" json:"id"`
	Slug      string    `json:"slug"`
	Title     string    `json:"title"`
	Content   string    `gorm:"type:text" json:"content"`
	Likes     uint8     `json:"likse"`
	CreatedAt time.Time `gorm:"primaryKey" json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func GetArticles() []Article {
	db := Connect()
	articles := []Article{}
	db.Find(&articles)

	return articles
}

func GetArticleSlug(slug string) Article {
	db := Connect()
	article := Article{Slug: slug}
	db.Find(&article)

	return article
}

func CreateArticle(article Article) Article {
	db := Connect()
	article.Slug = slug.Make(article.Title)
	db.Create(&article)

	return article
}

func DeleteArticle(id uint) Article {
	db := Connect()
	article := Article{ID: id}
	db.Delete(&article)
	return article
}
