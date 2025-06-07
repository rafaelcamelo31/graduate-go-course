package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/dto"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/infra/grpc/pb"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/usecase"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	int64Id, err := strconv.ParseInt(in.Id, 10, 64)
	dto := dto.OrderInputDTO{
		ID:    int64Id,
		Price: float64(in.Price),
		Tax:   float64(in.Tax),
	}
	output, err := s.CreateOrderUseCase.Execute(dto)
	if err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{
		Id:         fmt.Sprint(output.ID),
		Price:      float32(output.Price),
		Tax:        float32(output.Tax),
		FinalPrice: float32(output.FinalPrice),
	}, nil
}
