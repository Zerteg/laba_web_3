package routes

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/gorm"
	"laba_web_3/controllers"
	_ "laba_web_3/docs"
	"laba_web_3/middlewares"
)

var DB *gorm.DB

func RegisterAuthRoutes(router *gin.Engine) {

	productController := &controllers.ProductController{
		DB: DB,
	}

	feedbackController := &controllers.FeedbackController{
		DB: DB,
	}

	// Создаем подгруппу маршрутов для аутентификации
	authRoutes := router.Group("/auth")

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Маршрут для регистрации
	authRoutes.POST("/register", controllers.RegisterUser)

	// Маршрут для входа
	authRoutes.POST("/login", controllers.LoginUser)

	// Добавляем маршрут для отправки отзыва
	authRoutes.POST("/feedback", feedbackController.FeedbackHandler)

	// Группа маршрутов для защищенных ресурсов
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middlewares.AuthMiddleware()) // Применяем middleware

	// Пример защищенного маршрута - просмотр карточек товаров
	protectedRoutes.GET("/products", productController.GetProducts)

}
