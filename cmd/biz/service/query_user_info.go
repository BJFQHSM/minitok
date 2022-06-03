package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"strconv"
)

type QueryUserInfoService interface {
	DoService() *biz.QueryUserInfoResponse
}

func NewQueryUserInfoService(ctx context.Context, r *biz.QueryUserInfoRequest) QueryUserInfoService {
	return &queryUserInfoServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryUserInfoResponse{}}
}

type queryUserInfoServiceImpl struct {
	Req  *biz.QueryUserInfoRequest
	Resp *biz.QueryUserInfoResponse
	Ctx  context.Context
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
		if err = s.validateParams(); err != nil {
			break
		}
		var tokenUserId int64
		tokenUserId, err = strconv.ParseInt(s.Req.Token, 10, 64)
		if err != nil {
			break
		}
		if s.Resp.User, err = QueryUserInfoByUID(s.Ctx, s.Req.UserId, tokenUserId); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *queryUserInfoServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("params: userId could not be negative number")
	}
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

func QueryUserInfoByUID(ctx context.Context, uid int64, tokenUserId int64) (*biz.User, error) {
	user, err := dal.QueryUserById(ctx, uid)
	if err != nil {
		return nil, err
	}

	isFollow, err := dal.QueryIsFollow(ctx, uid, tokenUserId)
	if err != nil {
		return nil, err
	}
	respUser := biz.User{
		Id:            int64(user.UserId),
		Name:          user.Username,
		FollowCount:   int64(user.FollowCount),
		FollowerCount: int64(user.FollowerCount),
		IsFollow:      isFollow,
	}
	return &respUser, nil
}
