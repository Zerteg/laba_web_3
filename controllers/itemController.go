package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"laba_web_3/models"
	"net/http"
	"strconv"
)

// ItemController для работы с товарами
type ItemController struct {
	DB *gorm.DB
}

// GetItemsByCategory — Получить все товары по категории
func (ctrl *ItemController) GetItemsByCategory(c *gin.Context) {
	// Получаем product_id из URL
	productID, err := strconv.Atoi(c.Param("product_id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid product ID"})
		return
	}

	var items []models.Item
	if err := ctrl.DB.Where("product_id = ?", productID).Find(&items).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// GetAllProducts — Получить все категории
func (ctrl *ItemController) GetAllProducts(c *gin.Context) {
	var products []models.Product
	if err := ctrl.DB.Preload("Items").Find(&products).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, products)
}
