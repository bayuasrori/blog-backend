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

	models.CreateArticle(article)

	c.JSON(200, article)
}
