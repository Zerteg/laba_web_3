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
		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization token not provided"})
			c.Abort()
			return
		}

		// Проверка токена
		userID, err := utils.ValidateJWT(tokenString)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or expired token"})
			c.Abort()
			return
		}

		// Передача userID для использования в обработчике
		c.Set("userID", userID)
		c.Next()
	}
}
