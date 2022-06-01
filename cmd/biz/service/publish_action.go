package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type PublishActionService interface {
	DoService() *biz.PublishActionResponse
}


func NewPublishActionService(ctx context.Context, r *biz.PublishActionRequest) PublishActionService {
	return &publishActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.PublishActionResponse{}}
}

type publishActionServiceImpl struct {
	Req *biz.PublishActionRequest
	Resp *biz.PublishActionResponse
	Ctx context.Context
}

func (s *publishActionServiceImpl) DoService() *biz.PublishActionResponse {
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

func (s *publishActionServiceImpl) validateParams() error {
	return nil
}


func (s *publishActionServiceImpl) buildResponse(err error) {
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

