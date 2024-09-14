package models

import "time"

type Product struct {
	Id          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Price       float64   `json:"price"`
	Stock       int       `json:"stock"`
	SellerId    int       `json:"sellerId"`
	CreatedAt   time.Time `json:"createdAt"`
}

type ProductPayload struct {
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
	Stock       int     `json:"stock"`
	SellerId    int     `json:"sellerId"`
}
