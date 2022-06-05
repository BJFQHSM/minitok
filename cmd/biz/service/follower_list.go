package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type FollowerListService interface {
	DoService() *biz.QueryFollowerListResponse
}

func NewFollowerListService(ctx context.Context, r *biz.QueryFollowerListRequest) FollowerListService {
	return &followerListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryFollowerListResponse{}}
}

type followerListServiceImpl struct {
	Req  *biz.QueryFollowerListRequest
	Resp *biz.QueryFollowerListResponse
	Ctx  context.Context
}

func (s *followerListServiceImpl) DoService() *biz.QueryFollowerListResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		// todo
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *followerListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	return nil
}

func (s *followerListServiceImpl) buildResponse(err error) {
	if err != nil {
		errMsg := err.Error()
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 1
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 0
	}
}
