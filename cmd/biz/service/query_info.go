package service

import (
	"context"
	"errors"
	"strconv"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryUserInfoService interface {
	DoService()
	GetResponse() *biz.QueryUserInfoResponse
}

func NewQueryUserInfoService(r *biz.QueryUserInfoRequest, ctx context.Context) QueryUserInfoService {
	return &queryUserInfoServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryUserInfoResponse{}}
}

type queryUserInfoServiceImpl struct {
	Req  *biz.QueryUserInfoRequest
	Resp *biz.QueryUserInfoResponse
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
		return errors.New("params: userId could not be negative number")
	}
	return nil
}

func (s *queryUserInfoServiceImpl) queryUserInfoByUID() error {
	tokenUserId, err := strconv.ParseInt(s.Req.Token, 10, 64)
	if err != nil {
		return err
	}
	uid := s.Req.UserId
	user, err := dal.QueryUserById(s.Ctx, uid)
	if err != nil {
		return err
	}

	isFollow, err := dal.QueryIsFollow(s.Ctx, uid, int64(tokenUserId))
	if err != nil {
		return err
	}
	respUser := biz.User{
		Id:            int64(user.UserId),
		Name:          user.Username,
		FollowCount:   int64(user.FollowCount),
		FollowerCount: int64(user.FollowerCount),
		IsFollow:      isFollow,
	}
	s.Resp.User = &respUser
	return nil
}

func (s *queryUserInfoServiceImpl) GetResponse() *biz.QueryUserInfoResponse {
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
