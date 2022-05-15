package main

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/kitex_gen/user"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.RegisterRequest) (resp *user.RegisterResponse, err error) {
	// TODO: Your code here...
	return
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.LoginRequest) (resp *user.LoginResponse, err error) {
	// TODO: Your code here...
	resp = &user.LoginResponse{
		UserId: 11111,
	}
	return
}

// QueryInfo implements the UserServiceImpl interface.
func (s *UserServiceImpl) QueryInfo(ctx context.Context, req *user.QueryInfoRequest) (resp *user.QueryInfoResponse, err error) {
	// TODO: Your code here...
	return
}
