package service

import (
	"context"
	"errors"
	"strconv"
	"strings"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/db"
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
		return errors.New("params: userId could not be negative")
	}
	return nil
}

func (s *queryUserInfoServiceImpl) queryUserInfoByUID() error {
	tokenUserId, err := strconv.ParseInt(strings.Split(s.Req.Token, ".")[0], 10, 64)
	if err != nil {
		return err
	}
	uid := s.Req.UserId
	user, err := db.QueryUserByUID(s.Ctx, uid)

	if err != nil {
		return err
	}
	isFollow, err := db.QueryFollowUserByUID(s.Ctx, int64(tokenUserId), uid)
	if err != nil {
		return err
	}
	respUser := biz.User{
		Id:            int64(user.UserId),
		Name:          user.Username,
		FollowCount:   int64(user.FollowCount),
		FollowerCount: int64(user.FollowerCount),
		IsFollow:      &isFollow,
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
		s.Resp.StatusCode = 500
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 200
	}
}
