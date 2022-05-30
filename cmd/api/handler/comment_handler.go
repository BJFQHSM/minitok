package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func CommentAction(c *gin.Context) {
	var req biz.CommentActionRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)

	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.CommentAction(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

func QueryCommentList(c *gin.Context) {
	var req biz.QueryCommentListRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)

	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryCommentList(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
