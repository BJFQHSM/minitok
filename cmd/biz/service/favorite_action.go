package service

import (
	"context"
	"fmt"
	"log"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

type FavoriteActionService struct {
	ctx context.Context
}

func NewFavoriteActionService(ctx context.Context) *FavoriteActionService {
	return &FavoriteActionService{
		ctx: ctx,
	}
}

func (s *FavoriteActionService) FavoriteAction (req *biz.FavoriteActionRequest) (*biz.FavoriteActionResponse,error){
	resp := &biz.FavoriteActionResponse{}

	//获取用户的ID
	user,err := dal.QueryUserByToken(s.ctx, req.Token)
	if err!=nil{
		log.Printf("获取不到用户信息：%v", err)
		return nil,err
	}
	userId := user.UserId
	//点赞或取消点赞
	err = dal.FavoriteAction(s.ctx,userId,req.VideoId)

	if err!=nil{
		resp.StatusCode=1
		msg:="Favoriteaction failed"
		resp.StatusMsg=&msg
		return resp,err
	}
	resp.StatusCode=0
	msg:="Favoriteaction success"
	resp.StatusMsg=&msg
	return resp,nil
}
