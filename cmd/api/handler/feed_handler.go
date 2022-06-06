package handler

import (
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	var req biz.FeedRequest
	resp := &biz.FeedResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("Feed request: %+v\n", &req)
		authResp, err := rpc.AuthClient.Authenticate(c, &auth.AuthenticateRequest{Token: req.Token})
		if err != nil || authResp == nil {
			c.JSON(http.StatusInternalServerError, resp)
			return
		}
		if !authResp.IsAuthed {
			msg := "token invalid"
			resp = &biz.FeedResponse{
				StatusCode: 1,
				StatusMsg:  &msg,
			}
		} else {
			req.UserIdFromToken = authResp.UserId
			resp, err = rpc.BizClient.Feed(c, &req)
			if err != nil || resp == nil {
				c.JSON(http.StatusInternalServerError, resp)
				return
			}
		}
		util.LogInfof("Feed response: %+v\n", resp)
		c.JSON(http.StatusOK, &resp)
	}
}
