package service

import (
	"context"
	"errors"
	"log"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type FollowListService interface {
	DoService() *biz.QueryFollowListResponse
}

func NewFollowListService(ctx context.Context, r *biz.QueryFollowListRequest) FollowListService {
	return &followListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryFollowListResponse{}}
}

type followListServiceImpl struct {
	Req  *biz.QueryFollowListRequest
	Resp *biz.QueryFollowListResponse
	Ctx  context.Context

	userId int64
}

func (s *followListServiceImpl) DoService() *biz.QueryFollowListResponse {
	var err error
	for i := 0; i < 1; i++ {

		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.queryFollowList(); err != nil {
			break
		}

	}
	s.buildResponse(err)
	return s.Resp
}

func (s *followListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("illegal params: user_id cannot lower than 0")
	}
	return nil
}

func (s *followListServiceImpl) queryFollowList() error {

	users, err := dal.QueryFollowsByUserId(s.Ctx, s.Req.UserId)

	if err != nil {
		return nil
	}

	userList, err := DalUserToBizUser(s.Ctx, users, s.userId)
	if err != nil {
		return err
	}
	s.Resp.UserList = userList
	return nil
}

func (s *followListServiceImpl) buildResponse(err error) {
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

func DalUserToBizUser(ctx context.Context, users []*dal.User, tokenId int64) ([]*biz.User, error) {
	var result []*biz.User
	for _, user := range users {

		isFollow, err := dal.QueryIsFollow(ctx, user.UserId, tokenId)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}

		bizU := biz.User{
			Id:            user.UserId,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      isFollow,
		}

		result = append(result, &bizU)
	}
	return result, nil
}
