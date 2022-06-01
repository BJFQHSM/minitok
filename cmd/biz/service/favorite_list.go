package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryFavoriteListService interface {
	DoService() *biz.QueryFavoriteListResponse
}


func NewQueryFavoriteListService(ctx context.Context, r *biz.QueryFavoriteListRequest) QueryFavoriteListService {
	return &queryFavoriteListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryFavoriteListResponse{}}
}

type queryFavoriteListServiceImpl struct {
	Req *biz.QueryFavoriteListRequest
	Resp *biz.QueryFavoriteListResponse
	Ctx context.Context
}

func (s *queryFavoriteListServiceImpl) DoService() *biz.QueryFavoriteListResponse {
	// mock
	s.Resp = &biz.QueryFavoriteListResponse{
		VideoList: []*biz.Video{
			{Id: 2, Author: &biz.User{Id: 1}, CommentCount: 1, FavoriteCount: 2, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/6a87950831471cb0691b0f3dc2ae4428/628f35a7/video/tos/cn/tos-cn-ve-15c001-alinc2/cc660d533592437cb3377017d949ee13/?a=6383&ch=26&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=660&bt=660&cs=0&ds=3&ft=OyFYlOZZI0J.125TmVQbfzo57usylqG7Uag&mime_type=video_mp4&qs=0&rc=OWU7aGRlOjdkaDw1NmZnNEBpM3hweTU6ZmZvPDMzNGkzM0AxMTMuXzZiXzYxXzZjXzAvYSNkbWxzcjRvMGVgLS1kLS9zcw%3D%3D&l=202205261502500102020551523700C98B"},
			{Id: 3, Author: &biz.User{Id: 1}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
			{Id: 3, Author: &biz.User{Id: 1}, CommentCount: 2, FavoriteCount: 3, CoverUrl: "https://wallpapercave.com/wp/wp8233069.png", PlayUrl: "https://v26-web.douyinvod.com/8f4a23e58e5e011ff35f6f9a67d73dd8/628f3470/video/tos/cn/tos-cn-ve-15-alinc2/a71d679eb2d84b0cb5dbab33f68d94e1/?a=6383&ch=224&cr=0&dr=0&lr=all&cd=0%7C0%7C0%7C0&cv=1&br=1349&bt=1349&cs=0&ds=6&ft=5q_lc5mmnPD12Nuw3q.-UxHoFuYKc3wv25Na&mime_type=video_mp4&qs=0&rc=Omg8ZzkzN2Y7PDNoaWZnNUBpM3dkd2Q6Zmc3ODMzNGkzM0AzYjZjLWBgXi4xYWIzL2IzYSNuZi1fcjQwbHFgLS1kLTBzcw%3D%3D&l=202205261502460102081020853A00A5CB"},
		},
		StatusCode: 0,
	}

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams() ; err != nil {
			break
		}

		// todo biz
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *queryFavoriteListServiceImpl) validateParams() error {
	return nil
}


func (s *queryFavoriteListServiceImpl) buildResponse(err error) {
	if err != nil {
		errMsg := err.Error()
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 500
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 200
	}
}
