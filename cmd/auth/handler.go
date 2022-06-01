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
	resp := service.NewRegisterService(ctx, req).DoService()
	return resp, nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	resp := service.NewLoginService(ctx, req).DoService()
	return resp, nil
}
