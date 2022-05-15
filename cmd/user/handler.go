package main

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/user"
)

type UserServiceImpl struct {
	user.UnimplementedUserServiceServer
}

func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (*user.RegisterResponse, error) {
	return nil, nil
}

func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (*user.LoginResponse, error) {
	return &user.LoginResponse{
		UserId: 11111,
	}, nil
}
