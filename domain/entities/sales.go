package entities

import "time"

type Sale struct {
	ID          int       `json:"id" db:"id"`
	ProductID   int       `json:"product_id" db:"product_id"`
	ProductName string    `json:"product_name" db:"product_name"`
	Quantity    int       `json:"quantity" db:"quantity"`
	Price       float64   `json:"price" db:"price"`
	Total       float64   `json:"total" db:"total"`
	SaleDate    time.Time `json:"sale_date" db:"sale_date"`
}