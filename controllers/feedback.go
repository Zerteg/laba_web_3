package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"laba_web_3/middlewares"
	"laba_web_3/models"
	"net/http"
)

type FeedbackController struct {
	DB *gorm.DB
}

// HandleFeedback обработчик для маршрута "/feedback"

// FeedbackHandler - метод для получения обратной связи
// @Summary Submit feedback
// @Description This endpoint allows a user to submit feedback
// @Tags Users
// @Accept json
// @Produce json
// @Param feedback body models.Feedback true "Feedback data"
// @Success 200 {object} models.Feedback
// @Failure 400 {object} models.Error
// @Router /feedback [post]
func (ctrl *FeedbackController) FeedbackHandler(c *gin.Context) {
	// Структура для входных данных
	var input models.Feedback
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, что все поля заполнены
	if input.Name == "" || input.Email == "" || input.Message == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "All fields are required"})
		return
	}

	//// Сохраняем в базу данных
	//if err := ctrl.DB.Create(&input).Error; err != nil {
	//	c.JSON(http.StatusInternalServerError, gin.H{"error": "Не удалось сохранить отзыв"})
	//	return
	//}

	feedback := models.Feedback{
		Name:    input.Name,
		Email:   input.Email,
		Message: input.Message,
	}

	models.CreateFeedback(middlewares.DB, &feedback)

	fmt.Printf("Отзыв получен! Имя: %s, Почта: %s, Сообщение: %s\n", input.Name, input.Email, input.Message)

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted successfully"})
}
