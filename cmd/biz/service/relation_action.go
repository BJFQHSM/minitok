package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type RelationActionService interface {
	DoService() *biz.RelationActionResponse
}


func NewRelationActionService(ctx context.Context, r *biz.RelationActionRequest) RelationActionService {
	return &relationActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.RelationActionResponse{}}
}

type relationActionServiceImpl struct {
	Req *biz.RelationActionRequest
	Resp *biz.RelationActionResponse
	Ctx context.Context
}

func (s *relationActionServiceImpl) DoService() *biz.RelationActionResponse {
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

func (s *relationActionServiceImpl) validateParams() error {
	return nil
}


func (s *relationActionServiceImpl) buildResponse(err error) {
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


