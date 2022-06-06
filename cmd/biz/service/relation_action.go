package service

import (
	"context"
	"errors"
	"log"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"

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
}

func (s *relationActionServiceImpl) DoService() *biz.RelationActionResponse {
	var err error
	for i := 0; i < 1; i++ {

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

func (s *relationActionServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.ToUserId < 0 {
		return errors.New("illegal params: to_user_id could not lower than 0")
	}
	return nil
}

func (s *relationActionServiceImpl) doFollowAction() error {
	if s.Req.ActionType == 1 {
		err := dal.FollowRelation(s.Ctx, s.Req.ToUserId, s.Req.UserIdFromToken)
		if err != nil {
			log.Printf("%+v", err)
			return err
		}
	} else {
		err := dal.UnFollowRelation(s.Ctx, s.Req.ToUserId, s.Req.UserIdFromToken)
		if err != nil {
			log.Printf("%+v", err)
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
