package models

type Review struct {
	ID          uint   `json:"id" gorm:"primary_key"`
	ProductID   uint   `json:"product_id" `
	Rating      uint   `json:"rating"`
	Author      string `json:"author"`
	Description string `json:"description"`
	Date        string `json:"date"`
}
