package main

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
)

type BizServerImpl struct {
	biz.UnimplementedBizServiceServer
}

func (s *BizServerImpl) Feed(ctx context.Context, req *biz.FeedRequest) (*biz.FeedResponse, error) {
	util.LogInfof("Feed request: %+v\n", req)
	resp := service.NewFeedService(ctx, req).DoService()
	util.LogInfof("Feed response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) QueryUserInfo(ctx context.Context, req *biz.QueryUserInfoRequest) (*biz.QueryUserInfoResponse, error) {
	util.LogInfof("QueryUserInfo request: %+v\n", req)
	resp := service.NewQueryUserInfoService(ctx, req).DoService()
	util.LogInfof("QueryUserInfo response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) PublishAction(ctx context.Context, req *biz.PublishActionRequest) (*biz.PublishActionResponse, error) {
	//util.LogInfof("PublishAction request: %+v\n", req)
	resp := service.NewPublishActionService(ctx, req).DoService()
	util.LogInfof("PublishAction response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) QueryPublishList(ctx context.Context, req *biz.QueryPublishListRequest) (*biz.QueryPublishListResponse, error) {
	util.LogInfof("QueryPublishList request: %+v\n", req)
	resp := service.NewQueryPublishListService(req, ctx).DoService()
	util.LogInfof("QueryPublishList response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionRequest) (*biz.FavoriteActionResponse, error) {
	util.LogInfof("FavoriteAction request: %+v\n", req)
	resp := service.NewFavoriteActionService(ctx, req).DoService()
	util.LogInfof("FavoriteAction response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) QueryFavoriteList(ctx context.Context, req *biz.QueryFavoriteListRequest) (*biz.QueryFavoriteListResponse, error) {
	util.LogInfof("QueryFavoriteList request: %+v\n", req)
	resp := service.NewQueryFavoriteListService(ctx, req).DoService()
	util.LogInfof("QueryFavoriteList response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) CommentAction(ctx context.Context, req *biz.CommentActionRequest) (*biz.CommentActionResponse, error) {
	util.LogInfof("CommentAction request: %+v\n", req)
	resp := service.NewCommentActionService(ctx, req).DoService()
	util.LogInfof("CommentAction response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) QueryCommentList(ctx context.Context, req *biz.QueryCommentListRequest) (*biz.QueryCommentListResponse, error) {
	util.LogInfof("QueryCommentList request: %+v\n", req)
	resp := service.NewQueryCommentListService(ctx, req).DoService()
	util.LogInfof("QueryCommentList response: %+v\n", resp)
	return resp, nil
}
func (s *BizServerImpl) RelationAction(ctx context.Context, req *biz.RelationActionRequest) (*biz.RelationActionResponse, error) {
	util.LogInfof("RelationAction request: %+v\n", req)
	resp := service.NewRelationActionService(ctx, req).DoService()
	util.LogInfof("RelationAction response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) QueryFollowList(ctx context.Context, req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	util.LogInfof("QueryFollowList request: %+v\n", req)
	resp := service.NewFollowListService(ctx, req).DoService()
	util.LogInfof("QueryFollowList response: %+v\n", resp)
	return resp, nil
}

func (s *BizServerImpl) QueryFollowerList(ctx context.Context, req *biz.QueryFollowerListRequest) (*biz.QueryFollowerListResponse, error) {
	util.LogInfof("QueryFollowerList request: %+v\n", req)
	resp := service.NewFollowerListService(ctx, req).DoService()
	util.LogInfof("QueryFollowerList response: %+v\n", resp)
	return resp, nil
}
