package service

import (
	"context"
	"errors"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/db"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type QueryUserInfoService interface {
	DoService()
	GetResponse() *auth.QueryUserInfoResponse
}

func NewQueryUserInfoService(r *auth.QueryUserInfoRequest, ctx context.Context) QueryUserInfoService {
	return &queryUserInfoServiceImpl{Req: r, Ctx: ctx, Resp: &auth.QueryUserInfoResponse{}}
}

type queryUserInfoServiceImpl struct {
	Req  *auth.QueryUserInfoRequest
	Resp *auth.QueryUserInfoResponse
	Ctx  context.Context
}

func (s *queryUserInfoServiceImpl) DoService() {
	var err error

	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.queryUserInfoByUID(); err != nil {
			break
		}
	}
	s.buildResponse(err)
}

func (s *queryUserInfoServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("params: userId could not be negative")
	}
	return nil
}

func (s *queryUserInfoServiceImpl) queryUserInfoByUID() error {
	uid := s.Req.UserId
	user, err := db.QueryUserByUID(s.Ctx, uid)
	if err != nil {
		return err
	}
	// 没有做isFollow
	s.Resp.User.Id = int64(user.UserId)
	s.Resp.User.Name = user.Username
	s.Resp.User.FollowCount = int64(user.FollowCount)
	s.Resp.User.FollowerCount = int64(user.FollowerCount)
	s.Resp.User.IsFollow = true

	return nil
}

func (s *queryUserInfoServiceImpl) GetResponse() *auth.QueryUserInfoResponse {
	return s.Resp
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
