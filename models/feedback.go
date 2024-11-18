package models

// Feedback - структура для описания отзыва
// @Description Структура отзыва, включает имя, email и сообщение пользователя

type Feedback struct {
	ID      uint   `json:"id" gorm:"primaryKey"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}
