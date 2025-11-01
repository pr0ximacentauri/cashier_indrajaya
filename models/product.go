package models

import "time"

type Product struct {
	ID       string  `json:"id,omitempty"`
	Name     string  `json:"name"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
	Stock    int     `json:"stock"`
	CreatedAt time.Time  `json:"createdAt"`
}
