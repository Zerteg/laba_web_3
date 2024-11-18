package models

// Product - структура товара
// @Description Product object
// @Param product body models.Product true "Product data"
// @Success 200 {object} models.Product
// @Failure 400 {object} models.Error
type Product struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Title    string `json:"title"`
	ImageURL string `json:"image_url"`
}
