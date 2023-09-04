package models

type Product struct {
	ID          uint    `json:"id" gorm:"primary_key"`
	Brand       string  `json:"brand"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Category    string  `json:"category"`
	Price       float32 `json:"price"`
}
