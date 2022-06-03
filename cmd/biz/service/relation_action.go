package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type RelationActionService interface {
	DoService() *biz.RelationActionResponse
}

func NewRelationActionService(ctx context.Context, r *biz.RelationActionRequest) RelationActionService {
	return &relationActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.RelationActionResponse{}}
}

type relationActionServiceImpl struct {
	Req  *biz.RelationActionRequest
	Resp *biz.RelationActionResponse
	Ctx  context.Context

	userId int64
}

func (s *relationActionServiceImpl) DoService() *biz.RelationActionResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.authenticate(); err != nil {
			break
		}

		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doFollowAction(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *relationActionServiceImpl) authenticate() error {
	authReq := &auth.AuthenticateRequest{
		Token: s.Req.Token,
	}
	resp, err := rpc.AuthClient.Authenticate(s.Ctx, authReq)
	if err != nil {
		// todo
	}
	s.userId = resp.UserId
	return nil
}

func (s *relationActionServiceImpl) validateParams() error {
	return nil
}

func (s *relationActionServiceImpl) doFollowAction() error {
	if s.Req.ActionType == 1 {
		err := dal.FollowRelation(s.Ctx, s.Req.ToUserId, s.userId)
		if err != nil {
			return err
		}
	} else {
		err := dal.UnFollowRelation(s.Ctx, s.Req.ToUserId, s.userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *relationActionServiceImpl) buildResponse(err error) {
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
