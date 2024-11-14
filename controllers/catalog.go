package controllers

import (
	"laba_web_3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Controller struct {
	DB *gorm.DB
}

// Получить список товаров
func (ctrl *Controller) GetProducts(c *gin.Context) {
	var products []models.Product

	if err := ctrl.DB.Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось получить товары"})
		return
	}

	// Добавляем полный путь к картинке для каждого товара
	for i, product := range products {
		// Преобразуем относительный путь к изображению в полный URL
		// Например, если изображения хранятся в папке "images", то полный путь будет /images/имя_файла
		products[i].ImageURL = product.ImageURL
	}

	c.JSON(http.StatusOK, products)
}
