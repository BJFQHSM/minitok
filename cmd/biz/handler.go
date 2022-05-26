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
	return &biz.FeedResponse{
		Video: []*biz.Video{
			{
				Id:            1,
				Author:        &biz.User{Id: 1},
				PlayUrl:       "fsfs",
				CommentCount:  10,
				FavoriteCount: 10,
			},
		},
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryInfo(ctx context.Context, req *biz.QueryInfoRequest) (*biz.QueryInfoResponse, error) {
	return &biz.QueryInfoResponse{
		User: &biz.User{
			Id:            1,
			Name:          "dfs",
			FollowerCount: 10,
			FollowCount:   20,
		},
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) PublishAction(ctx context.Context, req *biz.PublishActionRequest) (*biz.PublishActionResponse, error) {
	return &biz.PublishActionResponse{
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryPublishList(ctx context.Context, req *biz.QueryPublishListRequest) (*biz.QueryPublishListResponse, error) {
	serv := service.NewQueryPublishListService(req, ctx)
	serv.DoService()
	return serv.GetResponse(), nil
}

func (s *BizServerImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionRequest) (*biz.FavoriteActionResponse, error) {
	return &biz.FavoriteActionResponse{
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryFavoriteList(ctx context.Context, req *biz.QueryFavoriteListRequest) (*biz.QueryFavoriteListResponse, error) {
	return &biz.QueryFavoriteListResponse{
		StatusCode: 0,
	}, nil
}
func (s *BizServerImpl) CommentAction(ctx context.Context, req *biz.CommentActionRequest) (*biz.CommentActionResponse, error) {
	return &biz.CommentActionResponse{
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryCommentList(ctx context.Context, req *biz.QueryCommentListRequest) (*biz.QueryCommentListResponse, error) {
	return &biz.QueryCommentListResponse{
		StatusCode: 0,
	}, nil
}
func (s *BizServerImpl) RelationAction(ctx context.Context, req *biz.RelationActionRequest) (*biz.RelationActionResponse, error) {
	return &biz.RelationActionResponse{
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryFollowList(ctx context.Context, req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	return &biz.QueryFollowListResponse{
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryFollowerList(ctx context.Context, req *biz.QueryFollowerListRequest) (*biz.QueryFollowerListResponse, error) {
	return &biz.QueryFollowerListResponse{
		StatusCode: 0,
	}, nil
}
