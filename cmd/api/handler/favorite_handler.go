package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func FavoriteAction(c *gin.Context) {
	var req biz.FavoriteActionRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.FavoriteAction(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

func QueryFavoriteList(c *gin.Context) {
	var req biz.QueryFavoriteListRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryFavoriteList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
