package handler

import (
	"goblog/middleware"

	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	articles := router.Group("/articles")
	{
		articles.GET("", GetArticles)
		articles.GET("/category/:name", GetArticlesByCategory)
		articles.GET("/category", GetCategories)
		articles.GET("/:slug", GetArticleSlug)
		author := articles.Group("")
		author.Use(middleware.AuthMiddleware())
		{
			author.POST("", CreateArticle)
			author.PUT("/:slug", CreateArticle)
			author.DELETE("", DeleteArticle)

			author.POST("/category", CreateCategory)
			author.PUT("/category/:name", GetArticles)
			author.DELETE("/category", DeleteCategory)

		}
	}

	authentication := router.Group("/auth")
	{
		authentication.POST("/login", Login)
		authentication.POST("/refresh", Refresh)
		authentication.POST("/register", Register)

	}

	return router
}
