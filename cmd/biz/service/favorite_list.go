package service

import (
	"context"
	"errors"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryFavoriteListService interface {
	DoService() *biz.QueryFavoriteListResponse
}

func NewQueryFavoriteListService(ctx context.Context, r *biz.QueryFavoriteListRequest) QueryFavoriteListService {
	return &queryFavoriteListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryFavoriteListResponse{}}
}

type queryFavoriteListServiceImpl struct {
	Req  *biz.QueryFavoriteListRequest
	Resp *biz.QueryFavoriteListResponse
	Ctx  context.Context
}

func (s *queryFavoriteListServiceImpl) DoService() *biz.QueryFavoriteListResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.queryFavoriteList(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *queryFavoriteListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("illegal params: user_id")
	}
	return nil
}

func (s *queryFavoriteListServiceImpl) queryFavoriteList() error {
	list, err := dal.GetFavoriteList(s.Ctx, s.Req.UserId)

	if err != nil {
		return err
	}

	for i := 0; i < len(list); i++ {
		s.Resp.VideoList = append(s.Resp.VideoList, MongoVdoToBizVdo(s.Ctx, list[i], s.Req.UserIdFromToken))
	}
	return nil
}

func (s *queryFavoriteListServiceImpl) buildResponse(err error) {
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
