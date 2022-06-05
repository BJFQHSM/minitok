package service

import (
	"context"
	"errors"
	"log"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
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

	userId int64
}

func (s *queryUserInfoServiceImpl) DoService() *biz.QueryUserInfoResponse {

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}
		if s.Resp.User, err = QueryUserInfoByUID(s.Ctx, s.Req.UserId, s.userId); err != nil {
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
		return errors.New("illegal params: user_id could not be negative number")
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
		log.Printf("%+v", err)
		return nil, err
	}

	isFollow, err := dal.QueryIsFollow(ctx, uid, tokenUserId)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	respUser := biz.User{
		Id:            user.UserId,
		Name:          user.Username,
		FollowCount:   user.FollowCount,
		FollowerCount: user.FollowerCount,
		IsFollow:      isFollow,
	}
	return &respUser, nil
}
