package handler

import (
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	var req biz.FavoriteActionRequest
	resp := &biz.FavoriteActionResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("FavoriteAction response: %+v\n", &req)

		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if authResp == nil || err != nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp.StatusMsg = &msg
		} else {
			req.UserId = authResp.UserId
			resp, err = rpc.BizClient.FavoriteAction(c, &req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		util.LogInfof("FavoriteAction response: %+v\n", resp)
		c.JSON(http.StatusOK, resp)
	}
}

func QueryFavoriteList(c *gin.Context) {
	var req biz.QueryFavoriteListRequest
	resp := &biz.QueryFavoriteListResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("QueryFavoriteList response: %+v\n", &req)
		resp, err = rpc.BizClient.QueryFavoriteList(c, &req)
		if err != nil || resp == nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			util.LogInfof("QueryFavoriteList response: %+v\n", resp)
			c.JSON(http.StatusOK, resp)
		}
	}
}
