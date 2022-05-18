package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal/db"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"log"
)

type QueryPublishListService struct {
	req biz.QueryPublishListRequest
	resp biz.QueryPublishListResponse
	ctx context.Context
}

func (s *QueryPublishListService) QueryPublishList() {
	videos, err := queryPublishListByUserId(s.ctx, s.req.UserId)
	if err != nil {
		log.Printf("ERROR: fail to query videolist by id %v\n", err)
		return
	}
	var videoList []*biz.Video
	for _, video := range videos {
		videoList = append(videoList, transDoToDto(video))
	}
}

func transDoToDto(video *db.Video) *biz.Video {
	return nil
}

func queryPublishListByUserId(ctx context.Context, userId int64) ([]*db.Video, error) {
	return db.QueryVideosByUserId(ctx, userId)
}
