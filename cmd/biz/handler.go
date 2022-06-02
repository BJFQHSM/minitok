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
	return &biz.FeedResponse{
		Video: []*biz.Video{
			{Id: 2, Author: &biz.User{Id: 2}, CommentCount: 1, FavoriteCount: 2, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/6a87950831471cb0691b0f3dc2ae4428/628f35a7/video/tos/cn/tos-cn-ve-15c001-alinc2/cc660d533592437cb3377017d949ee13/?a=6383&ch=26&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=660&bt=660&cs=0&ds=3&ft=OyFYlOZZI0J.125TmVQbfzo57usylqG7Uag&mime_type=video_mp4&qs=0&rc=OWU7aGRlOjdkaDw1NmZnNEBpM3hweTU6ZmZvPDMzNGkzM0AxMTMuXzZiXzYxXzZjXzAvYSNkbWxzcjRvMGVgLS1kLS9zcw%3D%3D&l=202205261502500102020551523700C98B"},
			{Id: 3, Author: &biz.User{Id: 3}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
			{Id: 3, Author: &biz.User{Id: 4}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
		},
		StatusCode: 0,
	}, nil
}

func (s *BizServerImpl) QueryUserInfo(ctx context.Context, req *biz.QueryUserInfoRequest) (*biz.QueryUserInfoResponse, error) {
	log.Printf("用户信息获取的参数为：%+v", req)
	ser := service.NewQueryUserInfoService(req, ctx)
	ser.DoService()
	return ser.GetResponse(), nil
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
		VideoList: []*biz.Video{
			{Id: 2, Author: &biz.User{Id: 1}, CommentCount: 1, FavoriteCount: 2, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/6a87950831471cb0691b0f3dc2ae4428/628f35a7/video/tos/cn/tos-cn-ve-15c001-alinc2/cc660d533592437cb3377017d949ee13/?a=6383&ch=26&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=660&bt=660&cs=0&ds=3&ft=OyFYlOZZI0J.125TmVQbfzo57usylqG7Uag&mime_type=video_mp4&qs=0&rc=OWU7aGRlOjdkaDw1NmZnNEBpM3hweTU6ZmZvPDMzNGkzM0AxMTMuXzZiXzYxXzZjXzAvYSNkbWxzcjRvMGVgLS1kLS9zcw%3D%3D&l=202205261502500102020551523700C98B"},
			{Id: 3, Author: &biz.User{Id: 1}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
			{Id: 3, Author: &biz.User{Id: 1}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
		},
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
	log.Printf("用户关注获取的参数为：%+v", req)
	resp := service.NewRelationActionService(ctx).RelationAction(req)
	return resp, nil
}

func (s *BizServerImpl) QueryFollowList(ctx context.Context, req *biz.QueryFollowListRequest) (*biz.QueryFollowListResponse, error) {
	log.Printf("用户关注列表获取的参数为：%+v", req)
	resp := service.NewQueryFollowListService(ctx).QueryFollowList(req)
	return resp, nil
}

func (s *BizServerImpl) QueryFollowerList(ctx context.Context, req *biz.QueryFollowerListRequest) (*biz.QueryFollowerListResponse, error) {
	return &biz.QueryFollowerListResponse{
		StatusCode: 0,
	}, nil
}
