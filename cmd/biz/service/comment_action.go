package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type CommentActionService interface {
	DoService() *biz.CommentActionResponse
}

func NewCommentActionService(ctx context.Context, r *biz.CommentActionRequest) CommentActionService {
	return &commentActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.CommentActionResponse{}}
}

type commentActionServiceImpl struct {
	Req  *biz.CommentActionRequest
	Resp *biz.CommentActionResponse
	Ctx  context.Context
}

func (s *commentActionServiceImpl) DoService() *biz.CommentActionResponse {
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

func (s *commentActionServiceImpl) validateParams() error {
	return nil
}

func (s *commentActionServiceImpl) buildResponse(err error) {
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
