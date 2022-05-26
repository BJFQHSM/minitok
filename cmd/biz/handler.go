package main

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/service"
	"log"
	"time"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type BizServerImpl struct {
	biz.UnimplementedBizServiceServer
}

func (s *BizServerImpl) Feed(ctx context.Context, req *biz.FeedRequest) (*biz.FeedResponse, error) {
	resp := new(biz.FeedResponse)
	videos, nextTime, err := service.NewFeedService(ctx).Feed(req)

	if err != nil {
		resp.Video = []*biz.Video{}
		resp.NextTime = time.Now().Unix()
		resp.StatusCode = 405
		msg := "Fail to get videos!"
		resp.StatusMsg = &msg
		return resp, err
	}
	resp.Video = videos
	resp.NextTime = nextTime
	resp.StatusCode = 200
	msg := "Success！"
	resp.StatusMsg = &msg
	return resp, nil
}

func (s *BizServerImpl) QueryInfo(ctx context.Context, req *biz.QueryInfoRequest) (*biz.QueryInfoResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) PublishAction(ctx context.Context, req *biz.PublishActionRequest) (*biz.PublishActionResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryPublishList(ctx context.Context, req *biz.QueryPublishListRequest) (*biz.QueryPublishListResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionRequest) (*biz.FavoriteActionResponse, error) {
	resp, err := service.NewFavoriteActionService(ctx).FavoriteAction(req)
	if err!=nil{
		return nil,errors.New("FavoriteAction failed")
	}
	return resp, nil
}

func (s *BizServerImpl) QueryFavoriteList(ctx context.Context, req *biz.QueryFavoriteListRequest) (*biz.QueryFavoriteListResponse, error) {
	resp,err:=service.NewFavoriteListService(ctx).FavoriteList(req)
	if err!=nil{
		log.Println(err)
		return nil,err
	}
	return resp, nil
}

func (s *BizServerImpl) CommentAction(ctx context.Context, req *biz.CommentActionRequest) (*biz.CommentActionResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryCommentList(ctx context.Context, req *biz.QueryCommentListRequest) (*biz.QueryCommentListResponse, error) {
	return nil, nil
}
func (s *BizServerImpl) RelationAction(ctx context.Context, req *biz.RelationActionRequest) (*biz.RelationActionResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryFollowList(ctx context.Context, req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryFollowerList(ctx context.Context, req *biz.QueryFollowerListRequest) (*biz.QueryFollowerListResponse, error) {
	return nil, nil
}
