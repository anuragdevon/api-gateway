package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"api-gateway/pkg/order/pb"
)

type OrderServiceClient struct {
	mock.Mock
}

func (m *OrderServiceClient) CreateOrder(ctx context.Context, in *pb.CreateOrderRequest, opts ...grpc.CallOption) (*pb.CreateOrderResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.CreateOrderResponse), args.Error(1)
}
