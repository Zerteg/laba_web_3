package models

import (
	"gorm.io/gorm"
	"time"
)

// Feedback - структура для хранения отзывов
// @Description Feedback object
type Feedback struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `json:"created_at"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"not null" json:"email"`
	Message   string    `gorm:"not null" json:"message"`
}

// CreateFeedback создает новый отзыв в базе данных
func CreateFeedback(db *gorm.DB, feedback *Feedback) error {
	result := db.Create(feedback)
	return result.Error
}
