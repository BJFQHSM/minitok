package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryUserInfoService interface {
	DoService() *biz.QueryUserInfoResponse
}


func NewQueryUserInfoService(ctx context.Context, r *biz.QueryUserInfoRequest) QueryUserInfoService {
	return &queryUserInfoServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryUserInfoResponse{}}
}

type queryUserInfoServiceImpl struct {
	Req *biz.QueryUserInfoRequest
	Resp *biz.QueryUserInfoResponse
	Ctx context.Context
}

func (s *queryUserInfoServiceImpl) DoService() *biz.QueryUserInfoResponse {
	// mock
	s.Resp = &biz.QueryUserInfoResponse{
		User: &biz.User{
			Id:            1,
			Name:          "dfs",
			FollowerCount: 10,
			FollowCount:   20,
		},
		StatusCode: 0,
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

func (s *queryUserInfoServiceImpl) validateParams() error {
	return nil
}


func (s *queryUserInfoServiceImpl) buildResponse(err error) {
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
