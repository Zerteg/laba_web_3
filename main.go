package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"laba_web_3/controllers"
	"laba_web_3/middlewares"
	"laba_web_3/models"
	"log"
	"time"
)

var DB *gorm.DB

func SetupDatabase() {
	var err error
	dsn := "user=postgres password=qwerty dbname=farengeit_shop port=5433 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Миграция
	DB.AutoMigrate(&models.User{}, &models.Product{})
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // Укажите origin, который будет использоваться
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	SetupDatabase()
	middlewares.DB = DB // сохраняем подключение к базе данных в мидлваре

	// Создаем объект контроллера с доступом к DB
	productController := &controllers.Controller{
		DB: DB,
	}

	// Публичные маршруты
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.POST("/feedback", controllers.FeedbackHandler)
	// Группа защищенных маршрутов
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middlewares.AuthMiddleware()) // добавляем AuthMiddleware

	// Защищенный маршрут для карточек товаров
	protectedRoutes.GET("/products", productController.GetProducts)

	// Обработчик для статичных файлов (изображений)
	protectedRoutes.Static("/media", "./media")

	router.Run(":8080")
}
