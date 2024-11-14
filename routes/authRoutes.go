package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"laba_web_3/controllers"
	"laba_web_3/middlewares"
)

var DB *gorm.DB

func RegisterAuthRoutes(router *gin.Engine) {

	productController := &controllers.Controller{
		DB: DB,
	}

	// Создаем подгруппу маршрутов для аутентификации
	authRoutes := router.Group("/auth")

	// Маршрут для регистрации
	authRoutes.POST("/register", controllers.RegisterUser)

	// Маршрут для входа
	authRoutes.POST("/login", controllers.LoginUser)

	// Группа маршрутов для защищенных ресурсов
	protectedRoutes := router.Group("/protected")
	protectedRoutes.Use(middlewares.AuthMiddleware()) // Применяем middleware

	// Пример защищенного маршрута - просмотр карточек товаров
	protectedRoutes.GET("/products", productController.GetProducts)

}
