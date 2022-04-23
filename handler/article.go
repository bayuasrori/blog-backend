package handler

import (
	"goblog/models"

	"github.com/gin-gonic/gin"
)

func GetArticles(c *gin.Context) {
	articles := models.GetArticles()

	c.JSON(200, articles)
}

func GetArticleSlug(c *gin.Context) {
	slug := c.Param("slug")
	article := models.GetArticleSlug(slug)
	c.JSON(200, article)
}

func CreateArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)

	new_article := models.CreateArticle(article)

	c.JSON(200, new_article)
}

func DeleteArticle(c *gin.Context) {
	var article models.Article
	c.BindJSON(&article)
	models.DeleteArticle(article.ID)

	c.JSON(200, article)
}

func GetArticlesByCategory(c *gin.Context) {
	name := c.Param("name")
	articles := models.GetArticlesByCategory(name)

	c.JSON(200, articles)
}
