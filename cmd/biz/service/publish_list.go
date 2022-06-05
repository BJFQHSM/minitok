package service

import (
	"context"
	"errors"
	"log"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryPublishListService interface {
	DoService() *biz.QueryPublishListResponse
}

func NewQueryPublishListService(r *biz.QueryPublishListRequest, ctx context.Context) QueryPublishListService {
	return &queryPublishListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryPublishListResponse{}}
}

type queryPublishListServiceImpl struct {
	Req  *biz.QueryPublishListRequest
	Resp *biz.QueryPublishListResponse
	Ctx  context.Context

	userId int64
}

func (s *queryPublishListServiceImpl) DoService() *biz.QueryPublishListResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.queryPublishListByUID(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *queryPublishListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("illegal params: user_id could not be negative")
	}
	return nil
}

func (s *queryPublishListServiceImpl) queryPublishListByUID() error {
	uid := s.Req.UserId
	videos, err := dal.QueryVideosByUserId(s.Ctx, uid)
	if err != nil {
		return err
	}
	videoList := s.Resp.GetVideoList()
	for _, video := range videos {
		v := transDoToDto(video, s.userId)
		videoList = append(videoList, v)
	}
	s.Resp.VideoList = videoList
	return nil
}

// todo extract to be a public method in other pkg
func transDoToDto(video *dal.Video, tokenId int64) *biz.Video {
	author, err := QueryUserInfoByUID(context.TODO(), video.UserId, tokenId)
	if err != nil {
		log.Println(err)
		return nil
	}

	ret := biz.Video{
		Id:            video.VideoId,
		Author:        author,
		PlayUrl:       video.PlayUrl,
		CoverUrl:      video.CoverUrl,
		FavoriteCount: video.FavoriteCount,
		CommentCount:  video.CommentCount,
		IsFavorite:    len(video.Favorites) != 0,
	}
	return &ret
}

func (s *queryPublishListServiceImpl) buildResponse(err error) {
	if err != nil {
		errMsg := err.Error()
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 1
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 0
	}
}
