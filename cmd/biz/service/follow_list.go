package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type FollowListService interface {
	DoService() *biz.QueryFollowListResponse
}


func NewFollowListService(ctx context.Context, r *biz.QueryFollowListRequest) FollowListService {
	return &followListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryFollowListResponse{}}
}

type followListServiceImpl struct {
	Req *biz.QueryFollowListRequest
	Resp *biz.QueryFollowListResponse
	Ctx context.Context
}

func (s *followListServiceImpl) DoService() *biz.QueryFollowListResponse {
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

func (s *followListServiceImpl) validateParams() error {
	return nil
}


func (s *followListServiceImpl) buildResponse(err error) {
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