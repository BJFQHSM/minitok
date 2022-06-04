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

	if err != nil {
		// todo
	}
	resp, err := rpc.AuthClient.Register(context.Background(), &req)
	if err != nil {
		// todo
	}

	log.Printf("Resp: %+v\n", resp)
	c.JSON(http.StatusOK, resp)
}

func QueryInfo(c *gin.Context) {
	var req biz.QueryUserInfoRequest
	err := c.ShouldBindQuery(&req)

	if err != nil {
		// todo
		log.Printf("ERROR: parse from http reqbody %v\n", err)
	}
	resp, err := rpc.BizClient.QueryUserInfo(context.Background(), &req)
	if err != nil {
		// todo
		log.Printf("ERROR:  %v\n", err)
	}

	log.Printf("Resp: %+v\n", resp)
	c.JSON(http.StatusOK, resp)
}
