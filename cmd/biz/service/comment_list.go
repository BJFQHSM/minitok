package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type QueryCommentListService interface {
	DoService() *biz.QueryCommentListResponse
}

func NewQueryCommentListService(ctx context.Context, r *biz.QueryCommentListRequest) QueryCommentListService {
	return &queryCommentListServiceImpl{Req: r, Ctx: ctx, Resp: &biz.QueryCommentListResponse{}}
}

type queryCommentListServiceImpl struct {
	Req  *biz.QueryCommentListRequest
	Resp *biz.QueryCommentListResponse
	Ctx  context.Context
}

func (s *queryCommentListServiceImpl) DoService() *biz.QueryCommentListResponse {
	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.queryCommentList(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *queryCommentListServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("request could not be nil")
	}
	if req.VideoId < 0 {
		return errors.New("illegal params: video_id could not be negative")
	}
	return nil
}

func (s *queryCommentListServiceImpl) queryCommentList() error {
	videoId := s.Req.VideoId
	comments, err := dal.QueryCommentLists(s.Ctx, videoId)
	if err != nil {
		return err
	}
	for _, comment := range comments {
		s.Resp.CommentList = append(s.Resp.CommentList, transDalCommentToDTO(comment))
	}
	return nil
}

func transDalCommentToDTO(comment *dal.Comment) *biz.Comment {
	return &biz.Comment{
		Id: comment.CommentId,
		User: &biz.User{
			Id: comment.UserId,
		},
		Content:    comment.Content,
		CreateDate: comment.CreateDate.String(),
	}
}

func (s *queryCommentListServiceImpl) buildResponse(err error) {
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
