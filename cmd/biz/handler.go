package main

import (
	"context"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type BizServerImpl struct {
	biz.UnimplementedBizServiceServer
}

func (s *BizServerImpl) Feed(ctx context.Context, req *biz.FeedRequest) (*biz.FeedResponse, error) {
	resp := service.NewFeedService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) QueryUserInfo(ctx context.Context, req *biz.QueryUserInfoRequest) (*biz.QueryUserInfoResponse, error) {
	resp := service.NewQueryUserInfoService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) PublishAction(ctx context.Context, req *biz.PublishActionRequest) (*biz.PublishActionResponse, error) {
	resp := service.NewPublishActionService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) QueryPublishList(ctx context.Context, req *biz.QueryPublishListRequest) (*biz.QueryPublishListResponse, error) {
	resp := service.NewQueryPublishListService(req, ctx).DoService()
	return resp, nil
}

func (s *BizServerImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionRequest) (*biz.FavoriteActionResponse, error) {
	resp := service.NewFavoriteActionService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) QueryFavoriteList(ctx context.Context, req *biz.QueryFavoriteListRequest) (*biz.QueryFavoriteListResponse, error) {
	resp := service.NewQueryFavoriteListService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) CommentAction(ctx context.Context, req *biz.CommentActionRequest) (*biz.CommentActionResponse, error) {
	resp := service.NewCommentActionService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) QueryCommentList(ctx context.Context, req *biz.QueryCommentListRequest) (*biz.QueryCommentListResponse, error) {
	resp := service.NewQueryCommentListService(ctx, req).DoService()
	return resp, nil
}
func (s *BizServerImpl) RelationAction(ctx context.Context, req *biz.RelationActionRequest) (*biz.RelationActionResponse, error) {
	resp := service.NewRelationActionService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) QueryFollowList(ctx context.Context, req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	resp := service.NewFollowListService(ctx, req).DoService()
	return resp, nil
}

func (s *BizServerImpl) QueryFollowerList(ctx context.Context, req *biz.QueryFollowerListRequest) (*biz.QueryFollowerListResponse, error) {
	resp := service.NewFollowerListService(ctx, req).DoService()
	return resp, nil
}
