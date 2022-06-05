package handler

import (
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func RelationAction(c *gin.Context) {
	var req biz.RelationActionRequest
	resp := &biz.RelationActionResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)
	if err != nil {
		msg := "invalid request"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("RelationAction response: %+v\n", &req)
		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if err != nil || authResp == nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp.StatusMsg = &msg
		} else {
			req.UserId = authResp.UserId
			resp, err = rpc.BizClient.RelationAction(c, &req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		util.LogInfof("RelationAction response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}

}

func QueryFollowList(c *gin.Context) {
	var req biz.QueryFollowListRequest
	resp := &biz.QueryFollowListResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("QueryFollowList response: %+v\n", &req)
		resp, err = rpc.BizClient.QueryFollowList(c, &req)
		if err != nil || resp == nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		util.LogInfof("QueryFollowList response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}
}

func QueryFollowerList(c *gin.Context) {
	var req biz.QueryFollowerListRequest
	resp := &biz.QueryFollowerListResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("QueryFollowerList response: %+v\n", &req)
		resp, err = rpc.BizClient.QueryFollowerList(c, &req)
		if err != nil || resp == nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		util.LogInfof("QueryFollowerList response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}
}
