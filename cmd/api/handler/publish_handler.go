package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// PublishAction godoc
// @Summary      publish video
// @Description  publish video
// @Tags         publish
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Success      200 {object} biz.PublishActionResponse
// @Failure      500 {object} biz.PublishActionResponse
// @Router       /publish/action [post]
func PublishAction(c *gin.Context) {
	var req biz.PublishActionRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.PublishAction(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// QueryPublishList godoc
// @Summary      get publish list
// @Description  get publish video by userId
// @Tags         publish
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Param        data body string true "video data"
// @Success      200 {object} biz.QueryPublishListResponse
// @Failure      500 {object} biz.QueryPublishListResponse
// @Router       /publish/list [get]
func QueryPublishList(c *gin.Context) {
	var req biz.QueryPublishListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryPublishList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
