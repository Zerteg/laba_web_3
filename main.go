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
	DB.AutoMigrate(&models.User{})
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

	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.GET("/profile", middlewares.AuthMiddleware(), controllers.GetUserProfile)

	router.Run(":8080")
}
