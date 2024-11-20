package middlewares

import (
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// Секретный ключ для подписи токена
var jwtKey = []byte("Authorization")
var DB *gorm.DB

// AuthMiddleware — middleware для проверки токена
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Получаем токен из заголовка Authorization
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен не предоставлен"})
			c.Abort()
			return
		}

		// Проверка, что заголовок начинается с "Bearer"
		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Неверный формат заголовка"})
			c.Abort()
			return
		}

		tokenString := parts[1]
		fmt.Println("Received token:", tokenString) // Логирование токена для отладки

		// Парсинг и проверка токена
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// Проверка алгоритма токена
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtKey, nil
		})

		if err != nil {
			fmt.Println("Error parsing token:", err) // Логирование ошибки парсинга токена
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный или истекший токен"})
			c.Abort()
			return
		}

		// Проверка истечения срока действия токена
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			exp, ok := claims["exp"].(float64)
			if !ok {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Ошибка проверки срока действия токена"})
				c.Abort()
				return
			}

			if time.Unix(int64(exp), 0).Before(time.Now()) {
				fmt.Println("Token has expired") // Логирование истечения токена
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Токен истек"})
				c.Abort()
				return
			}

			c.Next() // Токен валиден, продолжаем выполнение запроса
		} else {
			fmt.Println("Invalid token claims or token not valid") // Логирование проблемы с claims
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Недействительный токен"})
			c.Abort()
		}
	}
}
