package middlewares

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"laba_web_3/utils"
	"net/http"
)

// AuthMiddleware защищает маршруты, требующие аутентификации
var DB *gorm.DB

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			// Если токен не предоставлен, отправляем ошибку
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен аутентификации не предоставлен"})
			c.Abort()
			return
		}

		// Проверка токена JWT
		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			// Если токен недействителен или истек, отправляем ошибку
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный или истекший токен"})
			c.Abort()
			return
		}

		// Передача userID для использования в других обработчиках
		c.Set("userID", userID)
		c.Next()
	}
}
