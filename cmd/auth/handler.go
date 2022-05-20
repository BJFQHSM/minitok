package main

import (
	"context"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type AuthServiceImpl struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	return nil, nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	return &auth.LoginResponse{
		UserId: 11111,
	}, nil
}

func (s *AuthServiceImpl) QueryUserInfo(ctx context.Context, req *auth.QueryUserInfoRequest) (*auth.QueryUserInfoResponse, error) {
	ser := service.NewQueryUserInfoService(req, ctx)
	ser.DoService()
	return ser.GetResponse(), nil
}
