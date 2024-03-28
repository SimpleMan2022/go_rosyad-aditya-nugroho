package entities

import "github.com/google/uuid"

type Food struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Price       int       `json:"price"`
	Description string    `json:"description"`
}

type FoodRequest struct {
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

type FoodResponse struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}
