package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
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

	userID int64
}

func (s *followerListServiceImpl) DoService() *biz.QueryFollowerListResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.queryFollowerList(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *followerListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("illegal params: user_id cannot lower than 0")
	}
	return nil
}

func (s *followerListServiceImpl) queryFollowerList() error {

	users, err := dal.QueryFollowersByUserId(s.Ctx, s.Req.UserId)

	if err != nil {
		return nil
	}

	userList, err := DalUserToBizUser(s.Ctx, users, s.userID)
	if err != nil {
		return err
	}
	s.Resp.UserList = userList
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
