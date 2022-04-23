package handler

import (
	"goblog/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCategories(c *gin.Context) {
	category := models.GetCategories

	c.JSON(200, category)
}

func CreateCategory(c *gin.Context) {
	var category models.Category
	c.BindJSON(&category)

	new_category := models.CreateCategory(category)

	c.JSON(200, new_category)
}

func EditCategory(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	category := models.Category{ID: id}
	c.BindJSON(&category)

	new_category := models.CreateCategory(category)

	c.JSON(200, new_category)
}

func DeleteCategory(c *gin.Context) {
	var category models.Category
	c.BindJSON(&category)
	models.DeleteCategory(category.ID)

	c.JSON(200, category)
}
