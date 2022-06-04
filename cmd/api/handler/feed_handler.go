package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func Feed(c *gin.Context) {
	var req biz.FeedRequest
	err := c.ShouldBindQuery(&req)

	if err != nil {
		// todo
		log.Printf("Fail to get feed, an error has happened:%v!", err)
	}
	resp, err := rpc.BizClient.Feed(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("Fail to get feed, an error has happened:%v!", err)
	}

	log.Printf("Resp: %+v\n", resp)
	c.JSON(http.StatusOK, &resp)
}
