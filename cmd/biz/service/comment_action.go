package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"time"
)

type CommentActionService interface {
	DoService() *biz.CommentActionResponse
}

func NewCommentActionService(ctx context.Context, r *biz.CommentActionRequest) CommentActionService {
	return &commentActionServiceImpl{Req: r, Ctx: ctx, Resp: &biz.CommentActionResponse{}}
}

type commentActionServiceImpl struct {
	Req  *biz.CommentActionRequest
	Resp *biz.CommentActionResponse
	Ctx  context.Context
}

func (s *commentActionServiceImpl) DoService() *biz.CommentActionResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doCommentAction(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *commentActionServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if req.VideoId < 0 {
		return errors.New("illegal video id")
	}
	if req.ActionType == 1 {
		if req.CommentText == nil || *req.CommentText == "" {
			return errors.New("comment cannot be blank")
		}
	} else if req.ActionType == 2 {
		if req.CommentId == nil || *req.CommentId <= 0 {
			return errors.New("comment_id illegal")
		}
	}
	return nil
}

func (s *commentActionServiceImpl) doCommentAction() error {
	if s.Req.ActionType == 1 {
		return s.publishComment()
	}
	return s.deleteComment()
}

func (s *commentActionServiceImpl) publishComment() error {
	comment := &dal.Comment{
		CommentId:  int64(util.GenerateRandomInt32()),
		UserId:     s.Req.UserIdFromToken,
		Content:    *s.Req.CommentText,
		CreateDate: time.Now(),
	}
	return dal.PublishCommentAction(s.Ctx, s.Req.VideoId, comment)
}

func (s *commentActionServiceImpl) deleteComment() error {
	return dal.DeleteCommentAction(s.Ctx, s.Req.VideoId, *s.Req.CommentId)
}

func (s *commentActionServiceImpl) buildResponse(err error) {
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
