package models

import "time"

type Order struct {
	UserId    int       `json:"userId"`
	ProductId int       `json:"productId"`
	Quantity  int       `json:"quantity"`
	CreatedAt time.Time `json:"createdAt"`
}

type OrderCreationPayload struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
}

type OrderUpdatePayload struct {
	Quantity int `json:"quantity"`
}
