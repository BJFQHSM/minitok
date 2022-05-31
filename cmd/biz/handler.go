package main

import (
	"context"
	"log"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type BizServerImpl struct {
	biz.UnimplementedBizServiceServer
}

func (s *BizServerImpl) Feed(ctx context.Context, req *biz.FeedRequest) (*biz.FeedResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryUserInfo(ctx context.Context, req *biz.QueryUserInfoRequest) (*biz.QueryUserInfoResponse, error) {
	ser := service.NewQueryUserInfoService(req, ctx)
	ser.DoService()
	return ser.GetResponse(), nil
}

func (s *BizServerImpl) PublishAction(ctx context.Context, req *biz.PublishActionRequest) (*biz.PublishActionResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryPublishList(ctx context.Context, req *biz.QueryPublishListRequest) (*biz.QueryPublishListResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionRequest) (*biz.FavoriteActionResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryFavoriteList(ctx context.Context, req *biz.QueryFavoriteListRequest) (*biz.QueryFavoriteListResponse, error) {
	return nil, nil
}
func (s *BizServerImpl) CommentAction(ctx context.Context, req *biz.CommentActionRequest) (*biz.CommentActionResponse, error) {
	return nil, nil
}

func (s *BizServerImpl) QueryCommentList(ctx context.Context, req *biz.QueryCommentListRequest) (*biz.QueryCommentListResponse, error) {
	return nil, nil
}
func (s *BizServerImpl) RelationAction(ctx context.Context, req *biz.RelationActionRequest) (*biz.RelationActionResponse, error) {
	resp, err := service.NewRelationActionService(ctx).RelationAction(req)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	return resp, nil
}

func (s *BizServerImpl) QueryFollowList(ctx context.Context, req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	resp, err := service.NewQueryFollowListService(ctx).QueryFollowList(req)
	if err != nil {
		log.Printf("%+v", err)
		return nil, err
	}
	return resp, nil
}

func (s *BizServerImpl) QueryFollowerList(ctx context.Context, req *biz.QueryFollowerListRequest) (*biz.QueryFollowerListResponse, error) {
	return nil, nil
}
