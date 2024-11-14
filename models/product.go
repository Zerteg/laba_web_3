package models

type Product struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}
