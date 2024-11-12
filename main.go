package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"laba_web_3/controllers"
	"laba_web_3/middlewares"
	"laba_web_3/models"
	"log"
	"net/http"
	"os"
	"time"
)

// Переменная для базы данных и хранилище сессий
var (
	DB    *gorm.DB
	store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
)

func SetupDatabase() {
	var err error
	dsn := "user=postgres password=qwerty dbname=farengeit_shop port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Миграция
	DB.AutoMigrate(&models.User{})
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	SetupDatabase()
	middlewares.DB = DB

	// Роуты для регистрации, входа и проверки профиля
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", loginUser)
	router.GET("/profile", AuthMiddleware(), controllers.GetUserProfile)
	router.POST("/logout", logoutUser)

	router.Run(":8080")
}

// Обработчик входа в систему
func loginUser(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")
	// Выполните логику проверки учетных данных пользователя (например, controllers.LoginUser).
	// Предположим, что проверка прошла успешно:
	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "Успешный вход в систему"})
}

// Обработчик выхода из системы
func logoutUser(c *gin.Context) {
	session, _ := store.Get(c.Request, "session-name")

	// Очистить сессию
	session.Values["authenticated"] = false
	session.Save(c.Request, c.Writer)

	c.JSON(http.StatusOK, gin.H{"message": "Вы успешно вышли из системы"})
}

// Middleware для проверки сессии
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, err := store.Get(c.Request, "session-name")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Ошибка сервера"})
			c.Abort()
			return
		}

		// Проверяем, активна ли сессия
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Необходима авторизация"})
			c.Abort()
			return
		}

		c.Next()
	}
}
