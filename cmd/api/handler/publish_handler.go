package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func PublishAction(c *gin.Context) {
	var req biz.PublishActionRequest
	err := c.ShouldBindQuery(&req)

	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.PublishAction(context.Background(), &req)
	if err != nil {
		// todo
	}

	log.Printf("Resp: %+v\n", resp)
	c.JSON(http.StatusOK, resp)
}

func QueryPublishList(c *gin.Context) {
	var req biz.QueryPublishListRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		log.Printf("ERROR: parse from http reqbody %v\n", err)
		// todo
	}
	resp, err := rpc.BizClient.QueryPublishList(context.Background(), &req)
	if err != nil {
		log.Printf("ERROR:  %v\n", err)
		// todo
	}

	log.Printf("Resp: %+v\n", resp)
	c.JSON(http.StatusOK, resp)
}
