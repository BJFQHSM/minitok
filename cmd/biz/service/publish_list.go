package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal/db"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryPublishListService interface {
	DoService()
	GetResponse() *biz.QueryPublishListResponse
}


func NewQueryPublishListService(r *biz.QueryPublishListRequest, ctx context.Context) QueryPublishListService {
	return &queryPublishListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryPublishListResponse{}}
}

type queryPublishListServiceImpl struct {
	Req *biz.QueryPublishListRequest
	Resp *biz.QueryPublishListResponse
	Ctx context.Context
}

func (s *queryPublishListServiceImpl) DoService() {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams() ; err != nil {
			break
		}

		if err = s.queryPublishListByUID(); err != nil {
			break
		}
	}
	s.buildResponse(err)
}

func (s *queryPublishListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("params: userId could not be negative")
	}
	return nil
}

func (s *queryPublishListServiceImpl) queryPublishListByUID() error {
	uid := s.Req.UserId
	videos, err := db.QueryVideosByUserId(s.Ctx, uid)
	if err != nil {
		return err
	}
	videoList := s.Resp.GetVideoList()
	for _, video := range videos {
		v := transDoToDto(video)
		videoList = append(videoList, v)
	}
	s.Resp.VideoList = videoList
	return nil
}

func (s *queryPublishListServiceImpl) GetResponse() *biz.QueryPublishListResponse {
	return s.Resp
}

// todo extract to be a public method in other pkg
func transDoToDto(video *db.Video) *biz.Video {
	isFavorite := len(video.Favorites) != 0
	ret := biz.Video{
		Id: video.VideoId,
		Author: &biz.User{
			Id: video.UserId,
			// todo other info ?
		},
		PlayUrl: video.PlayUrl,
		CoverUrl: video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount: video.CommentCount,
		IsFavorite: &isFavorite,
	}
	return &ret
}

func (s *queryPublishListServiceImpl) buildResponse(err error) {
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

