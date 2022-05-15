package main

import (
	"context"
	"grpc_test/grpc_gen"
)


type UserServiceImpl struct {
	grpc_gen.UnimplementedUserServiceServer
}

func (s *UserServiceImpl) Register(ctx context.Context, req *grpc_gen.RegisterRequest) (*grpc_gen.RegisterResponse, error) {
	return nil, nil
}


func (s *UserServiceImpl) Login(ctx context.Context, req *grpc_gen.LoginRequest) (*grpc_gen.LoginResponse, error) {
	return &grpc_gen.LoginResponse{
		UserId: 11111,
	}, nil
}

func (s *UserServiceImpl) QueryInfo(ctx context.Context, req *grpc_gen.QueryInfoRequest) (*grpc_gen.QueryInfoResponse, error) {
	return nil, nil
}