package mocks

import (
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"

	"api-gateway/pkg/inventory/pb"
)

type InventoryServiceClient struct {
	mock.Mock
}

func (m *InventoryServiceClient) CreateItem(ctx context.Context, in *pb.CreateItemRequest, opts ...grpc.CallOption) (*pb.CreateItemResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.CreateItemResponse), args.Error(1)
}

func (m *InventoryServiceClient) GetItem(ctx context.Context, in *pb.GetItemRequest, opts ...grpc.CallOption) (*pb.GetItemResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.GetItemResponse), args.Error(1)
}

func (m *InventoryServiceClient) GetAllItems(ctx context.Context, in *pb.GetAllItemsRequest, opts ...grpc.CallOption) (*pb.GetAllItemsResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.GetAllItemsResponse), args.Error(1)
}

func (m *InventoryServiceClient) UpdateItem(ctx context.Context, in *pb.UpdateItemRequest, opts ...grpc.CallOption) (*pb.UpdateItemResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.UpdateItemResponse), args.Error(1)
}

func (m *InventoryServiceClient) DeleteItem(ctx context.Context, in *pb.DeleteItemRequest, opts ...grpc.CallOption) (*pb.DeleteItemResponse, error) {
	args := m.Called(ctx, in)
	return args.Get(0).(*pb.DeleteItemResponse), args.Error(1)
}
