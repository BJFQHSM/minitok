package service

import (
	"context"

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
	var resp *biz.FavoriteActionResponse

	//点赞或取消点赞
	err:=dal.FavoriteAction(s.ctx,req.UserId,req.VideoId)

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
