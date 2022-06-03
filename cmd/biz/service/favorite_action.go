package service

import (
	"context"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type FavoriteActionService interface {
	DoService() *biz.FavoriteActionResponse
}

func NewFavoriteActionService(ctx context.Context, r *biz.FavoriteActionRequest) FavoriteActionService {
	return &favoriteActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.FavoriteActionResponse{}}
}

type favoriteActionServiceImpl struct {
	Req  *biz.FavoriteActionRequest
	Resp *biz.FavoriteActionResponse
	Ctx  context.Context

	userId int64
}

func (s *favoriteActionServiceImpl) DoService() *biz.FavoriteActionResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.authenticate(); err != nil {
			break
		}

		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doFavoriteAction(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *favoriteActionServiceImpl) authenticate() error {
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

func (s *favoriteActionServiceImpl) validateParams() error {
	return nil
}

func (s *favoriteActionServiceImpl) doFavoriteAction() error {
	//获取用户的ID
	user,err := dal.QueryUserByToken(s.Ctx, s.Req.Token)
	if err!=nil{
		log.Printf("获取不到用户信息：%v", err)
		return err
	}
	//userId := s.Req.UserId
	//点赞或取消点赞
	err = dal.FavoriteAction(s.Ctx, user.UserId, s.Req.VideoId, s.Req.ActionType)
	return err
}

func (s *favoriteActionServiceImpl) buildResponse(err error) {
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
