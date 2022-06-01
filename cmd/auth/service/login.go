package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type LoginService interface {
	DoService() *auth.LoginResponse
}


func NewLoginService(ctx context.Context, r *auth.LoginRequest) LoginService {
	return &loginServiceImpl{Req: r, Ctx: ctx, Resp: &auth.LoginResponse{}}
}

type loginServiceImpl struct {
	Req *auth.LoginRequest
	Resp *auth.LoginResponse
	Ctx context.Context
}

func (s *loginServiceImpl) DoService() *auth.LoginResponse {
	// mock
	msg := "success"
	s.Resp = &auth.LoginResponse{
		UserId:     1,
		Token:      "fsjflsjdf",
		StatusCode: 0,
		StatusMsg:  &msg,
	}

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams() ; err != nil {
			break
		}

		// todo
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *loginServiceImpl) validateParams() error {
	return nil
}


func (s *loginServiceImpl) buildResponse(err error) {
	if err != nil {
		errMsg := err.Error()
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 500
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 200
	}
}
