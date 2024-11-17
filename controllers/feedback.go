package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"laba_web_3/models"
	"net/http"
)

// HandleFeedback обработчик для маршрута "/feedback"
func FeedbackHandler(c *gin.Context) {
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

	fmt.Printf("Отзыв получен! Имя: %s, Почта: %s, Сообщение: %s\n", input.Name, input.Email, input.Message)

	// Успешный ответ
	c.JSON(http.StatusOK, gin.H{"message": "Feedback submitted successfully"})
}
