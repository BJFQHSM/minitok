package service

import (
	"context"
	"log"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/db"
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

func (s *QueryFollowListService) QueryFollowList(req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	var resp biz.QueryFollowListResponse
	users, err := dal.QueryFollowsByUserId(s.ctx, req.UserId)

	if err != nil {
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		return &resp, err
	}
	userList, err := DalUserToBizUser(s.ctx, users)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	resp.UserList = userList
	resp.StatusCode = 0
	msg := "SUCCESS"
	resp.StatusMsg = &msg
	return &resp, nil
}
func DalUserToBizUser(ctx context.Context, users []*dal.User) ([]*biz.User, error) {
	var result []*biz.User
	for _, user := range users {
		//假设登录id为1
		row, err := db.QueryFollowUserByUID(ctx, 1, user.UserId)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
		f := false
		if len(row) != 0 {
			f = true
		}
		bizU := biz.User{
			Id:            user.UserId,
			Name:          user.Username,
			FollowCount:   user.FollowCount,
			FollowerCount: user.FollowerCount,
			IsFollow:      &f,
		}

		result = append(result, &bizU)
	}
	return result, nil
}
