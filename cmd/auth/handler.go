package main

import (
	"context"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type AuthServiceImpl struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	msg := "success"
	return &auth.RegisterResponse{
		UserId:     1,
		Token:      "fsjfs",
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	msg := "success"
	return &auth.LoginResponse{
		UserId:     2,
		Token:      "2",
		StatusCode: 0,
		StatusMsg:  &msg,
	}, nil
}
