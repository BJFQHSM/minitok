package service

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
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

	userId int64
	video  dal.Video
	ossUrl string
}

func (s *publishActionServiceImpl) DoService() *biz.PublishActionResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.authenticate(); err != nil {
			break
		}

		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.publishToOss(); err != nil {
			break
		}

		if err = s.publishToDatabase(); err != nil {
			// todo rollback oss action
			break
		}

	}
	s.buildResponse(err)
	return s.Resp
}

func (s *publishActionServiceImpl) authenticate() error {
	authReq := &auth.AuthenticateRequest{
		Token: s.Req.Token,
	}
	resp, err := rpc.AuthClient.Authenticate(s.Ctx, authReq)
	if err != nil {
		// todo
	}
	s.userId = resp.UserId
	return nil
}

func (s *publishActionServiceImpl) validateParams() error {
	return nil
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
