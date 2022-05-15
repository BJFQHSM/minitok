package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// CommentAction godoc
// @Summary      comment a video
// @Description  put comment
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Param        video_id body int true "video_id"
// @Param        action_type body int true "1 - comment 2 - delete"
// @Param        comment_text body string false "needed when action_type=1"
// @Param        comment_id body int false "needed when action_type=2"
// @Success      200 {object} biz.CommentActionResponse
// @Failure      500 {object} biz.CommentActionResponse
// @Router       /comment/action [post]
func CommentAction(c *gin.Context) {
	var req biz.CommentActionRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.CommentAction(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// QueryCommentList godoc
// @Summary      get comment list
// @Description  get comment list by video id
// @Tags         comment
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Param        video_id body int true "video_id"
// @Success      200 {object} biz.QueryCommentListResponse
// @Failure      500 {object} biz.QueryCommentListResponse
// @Router       /comment/list [get]
func QueryCommentList(c *gin.Context) {
	var req biz.QueryCommentListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryCommentList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
