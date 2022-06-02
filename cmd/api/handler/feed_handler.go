package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

// Feed godoc
// @Summary      get feed
// @Description  get feed
// @Tags         feed
// @Accept       json
// @Produce      json
// @Param        latest_time body int false "the latest time to get"
// @Success      200 {object} biz.FeedResponse
// @Failure      500 {object} biz.FeedResponse
// @Router       /feed [get]
func Feed(c *gin.Context) {
	var req biz.FeedRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
		log.Fatal("Fail to get feed, an error has happened:%v!",err)
	}
	resp, err := rpc.BizClient.Feed(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("Fail to get feed, an error has happened:%v!",err)
	}

	c.JSON(http.StatusOK, &resp)
}
