package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryCommentListService interface {
	DoService() *biz.QueryCommentListResponse
}


func NewQueryCommentListService(ctx context.Context, r *biz.QueryCommentListRequest) QueryCommentListService {
	return &queryCommentListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryCommentListResponse{}}
}

type queryCommentListServiceImpl struct {
	Req *biz.QueryCommentListRequest
	Resp *biz.QueryCommentListResponse
	Ctx context.Context
}

func (s *queryCommentListServiceImpl) DoService() *biz.QueryCommentListResponse {
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

func (s *queryCommentListServiceImpl) validateParams() error {
	return nil
}


func (s *queryCommentListServiceImpl) buildResponse(err error) {
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