package models

import (
	"log"
	"strconv"
	"time"

	"gorm.io/gorm"

	"github.com/gosimple/slug"
)

type Article struct {
	gorm.Model
	ID         uint        `gorm:"primaryKey" json:"id"`
	Slug       string      `json:"slug" gorm:"index:idx_slug,unique"`
	Title      string      `json:"title"`
	Content    string      `gorm:"type:text" json:"content"`
	Likes      uint8       `json:"likse"`
	CreatedAt  time.Time   `json:"created_at"`
	UpdatedAt  time.Time   `json:"updated_at"`
	Categories []*Category `gorm:"many2many:article_categories;"`
	Author     User
}

func AddArticleCategory(category Category, article Article) {
	db := Connect()
	db.Model(&article).Association("Categories").Append([]Category{category})
}

func GetArticlesByCategory(name string) []Article {
	db := Connect()
	articles := []Article{}
	db.Preload("Categories").Where("id = (SELECT article_id FROM article_categories, categories WHERE categories.name = ?)", name).Find(&articles)

	return articles
}

func GetArticles() []Article {
	db := Connect()
	articles := []Article{}
	db.Find(&articles)

	for a := range articles {
		if len(articles[a].Content) > 40 {
			articles[a].Content = articles[a].Content[0:40] + "..."
		}
	}
	return articles
}

func GetArticleSlug(slug string) Article {
	db := Connect()
	var article Article
	log.Printf(slug)
	db.Where("slug = ?", slug).First(&article)
	log.Printf(article.Title)

	return article
}

func CreateArticle(article Article) Article {
	db := Connect()
	article.Slug = slug.Make(article.Title) + strconv.Itoa(time.Now().Nanosecond())
	db.Create(&article)

	return article
}

func DeleteArticle(id uint) Article {
	db := Connect()
	article := Article{ID: id}
	db.Delete(&Article{}, id)
	return article
}
