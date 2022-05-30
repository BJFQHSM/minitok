package handler

import (
	"context"
	"log"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req auth.LoginRequest
	err := c.ShouldBindQuery(&req)
	log.Printf("reqeust : %+v\n", req)

	if err != nil {
		log.Println(err)
		// todo
	}
	resp, err := rpc.AuthClient.Login(context.Background(), &req)
	if err != nil {
		log.Println(err)
	}
	c.JSON(http.StatusOK, resp)
	log.Println(c.Errors)
}

func Register(c *gin.Context) {
	var req auth.RegisterRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		// todo
	}
	resp, err := rpc.AuthClient.Register(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

func QueryInfo(c *gin.Context) {
	var req biz.QueryInfoRequest
	err := c.ShouldBindQuery(&req)

	log.Printf("reqeust : %+v\n", req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryInfo(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
