package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
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
	// mock
	//s.Resp = &biz.QueryUserInfoResponse{
	//	User: &biz.User{
	//		Id:            1,
	//		Name:          "dfs",
	//		FollowerCount: 10,
	//		FollowCount:   20,
	//	},
	//	StatusCode: 0,
	//}

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.authenticate(); err != nil {
			break
		}
		if s.Resp.User, err = QueryUserInfoByUID(s.Ctx, s.Req.UserId, s.userId); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *queryUserInfoServiceImpl) authenticate() error {
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
