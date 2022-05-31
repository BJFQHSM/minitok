package service

import (
	"context"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

type RelationActionService struct {
	ctx context.Context
}

func NewRelationActionService(ctx context.Context) *RelationActionService {
	return &RelationActionService{
		ctx: ctx,
	}
}

func (s *RelationActionService) RelationAction(req *biz.RelationActionRequest) (*biz.RelationActionResponse, error) {
	var resp biz.RelationActionResponse
	err := dal.ChangeFollowRelation(s.ctx, req.ToUserId, req.UserId)

	if err != nil {
		resp.StatusCode = 1
		errMsg := err.Error()
		resp.StatusMsg = &errMsg
		return &resp, err
	}
	resp.StatusCode = 0
	msg := "SUCCESS"
	resp.StatusMsg = &msg
	return &resp, nil
}
