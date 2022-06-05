package main

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/service"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type AuthServiceImpl struct {
	auth.UnimplementedAuthServiceServer
}

func (s *AuthServiceImpl) Register(ctx context.Context, req *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	util.LogInfof("Register request: %+v\n", req)
	resp := service.NewRegisterService(ctx, req).DoService()
	util.LogInfof("Register response: %+v\n", resp)
	return resp, nil
}

func (s *AuthServiceImpl) Login(ctx context.Context, req *auth.LoginRequest) (*auth.LoginResponse, error) {
	util.LogInfof("Login request: %+v\n", req)
	resp := service.NewLoginService(ctx, req).DoService()
	util.LogInfof("Login response: %+v\n", resp)
	return resp, nil
}

func (s *AuthServiceImpl) Authenticate(ctx context.Context, req *auth.AuthenticateRequest) (*auth.AuthenticateResponse, error) {
	userId, err := dal.JwtParseUser(req.Token)
	resp := &auth.AuthenticateResponse{}
	if err != nil {
		resp.IsAuthed = false
	} else {
		resp.UserId = userId
		resp.IsAuthed = true
	}
	return resp, nil
}
