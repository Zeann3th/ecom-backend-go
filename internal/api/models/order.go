package models

import "time"

type Order struct {
	UserId    int       `json:"userId"`
	ProductId int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	Total     float64   `json:"total"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderPayload struct {
	UserId    int `json:"userId"`
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}
