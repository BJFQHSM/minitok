package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"io/ioutil"
	"time"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type PublishActionService interface {
	DoService() *biz.PublishActionResponse
}

func NewPublishActionService(ctx context.Context, r *biz.PublishActionRequest) PublishActionService {
	return &publishActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.PublishActionResponse{}}
}

type publishActionServiceImpl struct {
	Req  *biz.PublishActionRequest
	Resp *biz.PublishActionResponse
	Ctx  context.Context
}

func (s *publishActionServiceImpl) DoService() *biz.PublishActionResponse {
	var err error
	for i := 0; i < 1; i++ {

		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doPublish(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *publishActionServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("request could not be nil")
	}
	if req.UserId < 0 {
		return errors.New("illegal params: user_id could not be negative")
	}
	return nil
}

func (s *publishActionServiceImpl) doPublish() error {
	req := s.Req
	var err error
	filename := util.GenerateRandomStr(20) + ".mp4"
	if err = ioutil.WriteFile(filename, req.Data, 0666); err != nil {
		return err
	}
	videoId := req.UserId << 31 + int64(util.GenerateRandomInt32())
	url := "http:127.0.0.1:8080/video/" + filename
	video := &dal.Video{
		VideoId: videoId,
		UserId: req.UserId,
		PlayUrl: url,
		Favorites: []int64{},
		FavoriteCount: 0,
		Comments: []*dal.Comment{},
		CommentCount: 0,
		PublishDate: time.Now(),
		Title: req.Title,
	}
	return dal.PublishVideo(s.Ctx, s.Req.UserId, video)
}

func (s *publishActionServiceImpl) buildResponse(err error) {
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
