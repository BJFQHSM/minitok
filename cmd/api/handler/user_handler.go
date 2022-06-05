package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"net/http"

	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var req auth.LoginRequest
	resp := &auth.LoginResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("Login request: %+v\n", &req)
		resp, err = rpc.AuthClient.Login(context.Background(), &req)
		if err != nil || resp == nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			util.LogInfof("Login response: %+v\n", resp)
			c.JSON(http.StatusOK, resp)
		}
	}
}

func Register(c *gin.Context) {
	var req auth.RegisterRequest
	resp := &auth.RegisterResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("Register request: %+v\n", &req)
		resp, err = rpc.AuthClient.Register(c, &req)
		if err != nil || resp == nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			util.LogInfof("Register response: %+v\n", resp)
			c.JSON(http.StatusOK, resp)
		}
	}
}

func QueryUserInfo(c *gin.Context) {
	var req biz.QueryUserInfoRequest
	resp := &biz.QueryUserInfoResponse{StatusCode: 1}
	err := c.ShouldBindQuery(&req)

	if err != nil {
		msg := "invalid request params"
		resp.StatusMsg = &msg
		c.JSON(http.StatusBadRequest, resp)
	} else {
		util.LogInfof("QueryUserInfo request: %+v\n", &req)
		resp, err = rpc.BizClient.QueryUserInfo(context.Background(), &req)
		if err != nil || resp == nil {
			c.JSON(http.StatusInternalServerError, resp)
		} else {
			util.LogInfof("QueryUserInfo response: %+v\n", resp)
			c.JSON(http.StatusOK, resp)
		}
	}
}
