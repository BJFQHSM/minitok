package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FavoriteAction godoc
// @Summary      like video action
// @Description
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Param        video_id body int true "video_id"
// @Param        action_type body int true "1 - like 2 - unlike"
// @Success      200 {object} biz.FavoriteActionResponse
// @Failure      500 {object} biz.FavoriteActionResponse
// @Router       /favorite/action [post]
func FavoriteAction(c *gin.Context) {
	var req biz.FavoriteActionRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.FavoriteAction(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// QueryFavoriteList godoc
// @Summary      get favorite list
// @Description  get favorite list by userId
// @Tags         favorite
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Success      200 {object} biz.QueryFavoriteListResponse
// @Failure      500 {object} biz.QueryFavoriteListResponse
// @Router       /favorite/list [get]
func QueryFavoriteList(c *gin.Context) {
	var req biz.QueryFavoriteListRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryFavoriteList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
