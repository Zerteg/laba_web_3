package routes

import (
	"github.com/gin-gonic/gin"
	"laba_web_3/controllers"
	"laba_web_3/middlewares"
)

func RegisterAuthRoutes(router *gin.Engine) {
	// Создаем подгруппу маршрутов для аутентификации
	authRoutes := router.Group("/auth")

	// Маршрут для регистрации
	authRoutes.POST("/register", controllers.RegisterUser)

	// Маршрут для входа
	authRoutes.POST("/login", controllers.LoginUser)

	// Пример защищенного маршрута, доступного только авторизованным пользователям
	authRoutes.GET("/profile", middlewares.AuthMiddleware(), controllers.GetUserProfile)
}
