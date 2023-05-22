package mocks

import (
	"api-gateway/pkg/auth/pb"
	"context"

	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type MockAuthServiceClient struct {
	mock.Mock
}

func (m *MockAuthServiceClient) Register(ctx context.Context, in *pb.RegisterRequest, opts ...grpc.CallOption) (*pb.RegisterResponse, error) {
	args := m.Called(ctx, in)
	res := args.Get(0)
	err := args.Error(1)
	if res == nil {
		return nil, err
	}
	return res.(*pb.RegisterResponse), err
}

func (m *MockAuthServiceClient) Validate(ctx context.Context, in *pb.ValidateRequest, opts ...grpc.CallOption) (*pb.ValidateResponse, error) {
	args := m.Called(ctx, in)
	res := args.Get(0)
	err := args.Error(1)
	if res == nil {
		return nil, err
	}
	return res.(*pb.ValidateResponse), err
}

func (m *MockAuthServiceClient) Login(ctx context.Context, in *pb.LoginRequest, opts ...grpc.CallOption) (*pb.LoginResponse, error) {
	args := m.Called(ctx, in)
	res := args.Get(0)
	err := args.Error(1)
	if res == nil {
		return nil, err
	}
	return res.(*pb.LoginResponse), err
}
