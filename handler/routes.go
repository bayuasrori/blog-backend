package handler

import (
	"github.com/gin-gonic/gin"
)

func Routes() *gin.Engine {
	router := gin.Default()

	articles := router.Group("/articles")
	{
		articles.GET("", GetArticles)
		articles.GET("/:slug", GetArticleSlug)
		articles.POST("", CreateArticle)
		articles.DELETE("", DeleteArticle)
	}

	authentication := router.Group("/auth")
	{
		authentication.POST("/login", Login)
		authentication.POST("/refresh", Refresh)
		authentication.POST("/register", Register)

	}

	return router
}
