package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type RegisterService interface {
	DoService() *auth.RegisterResponse
}


func NewRegisterService(ctx context.Context, r *auth.RegisterRequest) RegisterService {
	return &registerServiceImpl{Req: r, Ctx: ctx, Resp: &auth.RegisterResponse{}}
}

type registerServiceImpl struct {
	Req *auth.RegisterRequest
	Resp *auth.RegisterResponse
	Ctx context.Context
}

func (s *registerServiceImpl) DoService() *auth.RegisterResponse {
	// mock
	msg := "success"
	s.Resp = &auth.RegisterResponse{
		UserId:     1,
		Token:      "fsjfs",
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

func (s *registerServiceImpl) validateParams() error {
	return nil
}


func (s *registerServiceImpl) buildResponse(err error) {
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

