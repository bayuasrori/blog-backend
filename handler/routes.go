package handler

import "github.com/gin-gonic/gin"

func Routes() *gin.Engine {
	router := gin.Default()

	router.GET("/articles", GetArticles)
	router.GET("/articles/:slug", GetArticleSlug)
	router.POST("/articles", CreateArticle)

	return router
}
