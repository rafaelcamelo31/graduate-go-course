package dto

import "github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/entity"

type OrderInputDTO struct {
	ID    string  `json:"id"`
	Price float64 `json:"price"`
	Tax   float64 `json:"tax"`
}

type OrderOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

type OrdersInputDTO struct {
}

type OrdersOutputDTO struct {
	Orders []entity.Order `json:"orders"`
}
