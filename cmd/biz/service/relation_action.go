package service

import (
	"context"
	"log"

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

func (s *RelationActionService) RelationAction(req *biz.RelationActionRequest) *biz.RelationActionResponse {
	resp := &biz.RelationActionResponse{}
	if req.ActionType == 1 {
		err := dal.FollowRelation(s.ctx, req.ToUserId, req.UserId)
		if err != nil {
			log.Printf("关注运行到这了---%+v", err)
			resp.StatusCode = 1
			errMsg := err.Error()
			resp.StatusMsg = &errMsg
			log.Printf("resp---%+v", resp)
			return resp
		}
		resp.StatusCode = 0
		msg := "Follow this user successfully"
		resp.StatusMsg = &msg
		return resp
	} else {
		err := dal.UnFollowRelation(s.ctx, req.ToUserId, req.UserId)

		if err != nil {
			log.Printf("运行到这了---%+v", err)
			resp.StatusCode = 1
			errMsg := err.Error()
			resp.StatusMsg = &errMsg
			return resp
		}
		resp.StatusCode = 0
		msg := "Unfollow this user successfully"
		resp.StatusMsg = &msg
		return resp
	}

}
