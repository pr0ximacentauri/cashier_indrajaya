package models

import "time"

type TransactionItem struct {
	ProductID   string  `json:"productId"`
	ProductName string  `json:"productName"`
	Quantity    int     `json:"quantity"`
	Price       float64 `json:"price"`
	Subtotal    float64 `json:"subtotal"`
}

type Transaction struct {
	ID        string             `json:"id,omitempty"`
	Items     []TransactionItem  `json:"items"`
	Total     float64            `json:"total"`
	CreatedAt time.Time          `json:"createdAt"`
}
