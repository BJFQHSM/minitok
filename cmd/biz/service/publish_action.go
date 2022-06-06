package service

import (
	"context"
<<<<<<< HEAD
=======
	"errors"
	"io/ioutil"
	"os"
	"time"

>>>>>>> cloud_deploy
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

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
<<<<<<< HEAD

	userId int64
	video  dal.Video
	ossUrl string
=======
>>>>>>> cloud_deploy
}

func (s *publishActionServiceImpl) DoService() *biz.PublishActionResponse {
	var err error
	for i := 0; i < 1; i++ {

		if err = s.validateParams(); err != nil {
			break
		}

<<<<<<< HEAD
		if err = s.publishToOss(); err != nil {
			break
		}

		if err = s.publishToDatabase(); err != nil {
			// todo rollback oss action
			break
		}

=======
		if err = s.doPublish(); err != nil {
			break
		}
>>>>>>> cloud_deploy
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *publishActionServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("request could not be nil")
	}
	return nil
}

func (s *publishActionServiceImpl) doPublish() error {
	req := s.Req
	var err error
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	filename := util.GenerateRandomStr(20) + ".mp4"
	if err = ioutil.WriteFile(path+filename, req.Data, 0666); err != nil {
		return err
	}
	videoId := req.UserIdFromToken<<31 + int64(util.GenerateRandomInt32())
	url := "http:127.0.0.1:8080/static/" + filename
	video := &dal.Video{
		VideoId:       videoId,
		UserId:        req.UserIdFromToken,
		PlayUrl:       url,
		Favorites:     []int64{},
		FavoriteCount: 0,
		Comments:      []*dal.Comment{},
		CommentCount:  0,
		PublishDate:   time.Now(),
		Title:         req.Title,
	}
	return dal.PublishVideo(s.Ctx, s.Req.UserIdFromToken, video)
}

func (s *publishActionServiceImpl) publishToOss() error {
	url, err := dal.PublishToOss(s.Req.Data)
	if err != nil {
		return err
	}
	s.ossUrl = url
	return nil
}

func (s *publishActionServiceImpl) publishToDatabase() error {
	s.video = dal.Video{
		VideoId:       generateVideoId(s.userId),
		UserId:        s.userId,
		PlayUrl:       s.ossUrl,
		FavoriteCount: 0,
		Favorites:     []int64{},
		CommentCount:  0,
		Comments:      []dal.Comment{},
		Title:         s.Req.Title,
	}
	err := dal.PublishVideo(s.Ctx, s.video)
	if err != nil {
		return err
	}
	return nil
}

func generateVideoId(userId int64) int64 {
	low32bitUserId := int32(userId)
	return int64(low32bitUserId)<<31 + int64(util.GenerateRandomInt32())
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
