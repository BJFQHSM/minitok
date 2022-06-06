package handler

import (
	"context"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	var req biz.CommentActionRequest
	resp := &biz.CommentActionResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("CommentAction response: %+v\n", &req)
		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if authResp == nil || err != nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp.StatusMsg = &msg
		} else {
			req.UserIdFromToken = authResp.UserId
			resp, err = rpc.BizClient.CommentAction(c, &req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		util.LogInfof("CommentAction response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}

}

func QueryCommentList(c *gin.Context) {
	var req biz.QueryCommentListRequest
	resp := &biz.QueryCommentListResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("QueryCommentList response: %+v\n", &req)
		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if authResp == nil || err != nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp.StatusMsg = &msg
		} else {
			req.UserIdFromToken = authResp.UserId
			resp, err = rpc.BizClient.QueryCommentList(context.Background(), &req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		util.LogInfof("QueryCommentList response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}
}
