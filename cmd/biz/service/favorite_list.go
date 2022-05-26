package service

import (
	"context"
	"log"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

type FavoriteListService struct {
	ctx context.Context
}

func NewFavoriteListService(ctx context.Context)*FavoriteListService{
	return &FavoriteListService{
		ctx,
	}
}

func (s *FavoriteListService) FavoriteList(req *biz.QueryFavoriteListRequest)(*biz.QueryFavoriteListResponse,error){
	var resp *biz.QueryFavoriteListResponse
	list,err:=dal.GetFavoriteList(s.ctx, req.UserId)
	msg:="fail"
	if err!=nil{
		log.Println(err)
		resp.StatusCode=1
		resp.StatusMsg=&msg
		return resp,err
	}
	for i:=0;i<len(list);i++{
		resp.VideoList=append(resp.VideoList, MongoVdoToBizVdo(list[i]))
	}
	resp.StatusCode=0
	msg="success"
	resp.StatusMsg=&msg
	return resp, nil
}