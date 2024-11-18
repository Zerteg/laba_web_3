package controllers

import (
	"laba_web_3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ProductController struct {
	DB *gorm.DB
}

// Получить список товаров

// GetProducts - Получить список товаров
// @Summary Получить все товары
// @Description Возвращает список всех товаров с изображениями
// @Tags Products
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {array} models.Product "Список товаров"
// @Failure 401 {object} models.Error "Неавторизованный доступ"
// @Failure 500 {object} models.Error "Ошибка сервера"
// @Router /protected/products [get]
func (ctrl *ProductController) GetProducts(c *gin.Context) {
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
