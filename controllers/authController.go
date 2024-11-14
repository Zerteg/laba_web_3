package controllers

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"laba_web_3/middlewares"
	"laba_web_3/models"
	"laba_web_3/utils"
	"net/http"
)

var db *gorm.DB

func RegisterUser(c *gin.Context) {
	var input models.User
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Проверяем, существует ли уже пользователь с таким именем
	existingUser, _ := models.GetUserByUsername(middlewares.DB, input.Username)
	if existingUser != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username already exists"})
		return
	}

	// Создаем нового пользователя
	password_hash, err := utils.HashPassword(input.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: password_hash, // Хешированный пароль
	}

	err = models.CreateUser(middlewares.DB, &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User registered successfully"})
}

// Функция для входа пользователя
func LoginUser(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Проверка существования пользователя
	user, err := models.GetUserByEmail(middlewares.DB, input.Email)

	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user"})
		return
	}

	// Сравнение пароля
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Создание и отправка JWT-токена
	token, err := utils.GenerateJWT(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})

}

func GetUserProfile(c *gin.Context) {
	userID := c.GetString("userID") // Получаем ID пользователя из контекста, который был сохранен в AuthMiddleware

	// Получаем профиль пользователя из базы данных
	var user models.User
	user, err := models.GetUserProfile(db, userID) // вызываем функцию для получения профиля по ID
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch user profile"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"profile": user})
}
