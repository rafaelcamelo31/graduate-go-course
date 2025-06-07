package usecase

import (
	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/dto"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/entity"
)

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrdersUseCase(
	OrderRepository entity.OrderRepositoryInterface,

) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (c *ListOrdersUseCase) Execute() (*dto.OrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return nil, err
	}
	output := &dto.OrdersOutputDTO{
		Orders: orders,
	}

	return output, nil
}
