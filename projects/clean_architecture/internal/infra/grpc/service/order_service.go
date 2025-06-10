package service

import (
	"context"
	"fmt"
	"strconv"

	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/dto"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/infra/grpc/pb"
	"github.com/rafaelcamelo31/graduate-go-course/4-module/clean_architecture/internal/usecase"
	"google.golang.org/protobuf/types/known/emptypb"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

func NewOrderService(createOrderUseCase usecase.CreateOrderUseCase, listOrdersUseCase usecase.ListOrdersUseCase) *OrderService {
	return &OrderService{
		CreateOrderUseCase: createOrderUseCase,
		ListOrdersUseCase:  listOrdersUseCase,
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

func (s *OrderService) GetOrders(context.Context, *emptypb.Empty) (*pb.Orders, error) {
	output, err := s.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	orders := &pb.Orders{}
	for _, v := range output.Orders {
		orders.Orders = append(orders.Orders,
			&pb.Order{
				Id:         fmt.Sprint(v.ID),
				Price:      float32(v.Price),
				Tax:        float32(v.Tax),
				FinalPrice: float32(v.FinalPrice),
			},
		)
	}

	return orders, nil
}
