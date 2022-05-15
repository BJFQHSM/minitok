package main

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/kitex_gen/biz"
)

// BizServiceImpl implements the last service interface defined in the IDL.
type BizServiceImpl struct{}

// Feed implements the BizServiceImpl interface.
func (s *BizServiceImpl) Feed(ctx context.Context, req *biz.FeedRequest) (resp *biz.FeedResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishAction implements the BizServiceImpl interface.
func (s *BizServiceImpl) PublishAction(ctx context.Context, req *biz.PublishActionRequest) (resp *biz.PublishActionResponse, err error) {
	// TODO: Your code here...
	return
}

// PublishList implements the BizServiceImpl interface.
func (s *BizServiceImpl) PublishList(ctx context.Context, req *biz.PublishListRequest) (resp *biz.PublishListResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteAction implements the BizServiceImpl interface.
func (s *BizServiceImpl) FavoriteAction(ctx context.Context, req *biz.FavoriteActionRequest) (resp *biz.FavoriteActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FavoriteList implements the BizServiceImpl interface.
func (s *BizServiceImpl) FavoriteList(ctx context.Context, req *biz.FavoriteListRequest) (resp *biz.FavoriteListResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentAction implements the BizServiceImpl interface.
func (s *BizServiceImpl) CommentAction(ctx context.Context, req *biz.CommentActionRequest) (resp *biz.CommentActionResponse, err error) {
	// TODO: Your code here...
	return
}

// CommentList implements the BizServiceImpl interface.
func (s *BizServiceImpl) CommentList(ctx context.Context, req *biz.CommentListRequest) (resp *biz.CommentListResponse, err error) {
	// TODO: Your code here...
	return
}

// RelationAction implements the BizServiceImpl interface.
func (s *BizServiceImpl) RelationAction(ctx context.Context, req *biz.RelationActionRequest) (resp *biz.RelationActionResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowList implements the BizServiceImpl interface.
func (s *BizServiceImpl) FollowList(ctx context.Context, req *biz.FollowListRequest) (resp *biz.FollowListResponse, err error) {
	// TODO: Your code here...
	return
}

// FollowerList implements the BizServiceImpl interface.
func (s *BizServiceImpl) FollowerList(ctx context.Context, req *biz.FollowerListRequest) (resp *biz.FollowerListResponse, err error) {
	// TODO: Your code here...
	return
}
