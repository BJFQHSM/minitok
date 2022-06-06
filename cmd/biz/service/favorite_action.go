package service

import (
	"context"
	"errors"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type FavoriteActionService interface {
	DoService() *biz.FavoriteActionResponse
}

func NewFavoriteActionService(ctx context.Context, r *biz.FavoriteActionRequest) FavoriteActionService {
	return &favoriteActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.FavoriteActionResponse{}}
}

type favoriteActionServiceImpl struct {
	Req  *biz.FavoriteActionRequest
	Resp *biz.FavoriteActionResponse
	Ctx  context.Context
}

func (s *favoriteActionServiceImpl) DoService() *biz.FavoriteActionResponse {
	var err error
	for i := 0; i < 1; i++ {

		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doFavoriteAction(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *favoriteActionServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.VideoId < 0 {
		return errors.New("illegal video id")
	} else if req.ActionType != 1 && req.ActionType != 2 {
		return errors.New("illegal favorite action type")
	}
	return nil
}

func (s *favoriteActionServiceImpl) doFavoriteAction() error {
	return dal.FavoriteAction(s.Ctx, s.Req.UserIdFromToken, s.Req.VideoId, s.Req.ActionType)
}

func (s *favoriteActionServiceImpl) buildResponse(err error) {
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
