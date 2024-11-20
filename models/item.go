package models

import "database/sql"

type Item struct {
	ID          int     `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	ImageURL    string  `json:"image_url"`
	ProductID   int     `json:"product_id"`
}

// Получить все товары
func GetItems(db *sql.DB) ([]Item, error) {
	rows, err := db.Query("SELECT id, title, description, price, image_url, product_id FROM items")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.ImageURL, &item.ProductID); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Получить товары по категории
func GetItemsByCategory(db *sql.DB, productID int) ([]Item, error) {
	rows, err := db.Query("SELECT id, title, description, price, image_url, product_id FROM items WHERE product_id = $1", productID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var items []Item
	for rows.Next() {
		var item Item
		if err := rows.Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.ImageURL, &item.ProductID); err != nil {
			return nil, err
		}
		items = append(items, item)
	}
	return items, nil
}

// Получить товар по ID
func GetItemByID(db *sql.DB, itemID int) (*Item, error) {
	var item Item
	err := db.QueryRow("SELECT id, title, description, price, image_url, product_id FROM items WHERE id = $1", itemID).Scan(&item.ID, &item.Title, &item.Description, &item.Price, &item.ImageURL, &item.ProductID)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	return &item, nil
}
