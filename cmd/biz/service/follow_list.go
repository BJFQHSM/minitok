package service

import (
	"context"
	"log"
	"strconv"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryFollowListService struct {
	ctx context.Context
}

func NewQueryFollowListService(ctx context.Context) *QueryFollowListService {
	return &QueryFollowListService{
		ctx: ctx,
	}
}

func (s *QueryFollowListService) QueryFollowList(req *biz.QueryFollowListRequest) *biz.QueryFollowListResponse {
	resp := &biz.QueryFollowListResponse{}
	users, err := dal.QueryFollowsByUserId(s.ctx, req.UserId)

	if err != nil {
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		return resp
	}
	tokenId, err := strconv.ParseInt(req.Token, 10, 64)
	if err != nil {
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		return resp
	}
	userList, err := DalUserToBizUser(s.ctx, users, tokenId)
	if err != nil {
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		return resp
	}
	resp.UserList = userList
	resp.StatusCode = 0
	msg := "SUCCESS"
	resp.StatusMsg = &msg
	return resp
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
