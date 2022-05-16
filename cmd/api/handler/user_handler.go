package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary      login
// @Description  login
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        username body string true "username"
// @Param        password body string true "password"
// @Success      200 {object} auth.LoginResponse
// @Failure      500 {object} auth.LoginResponse
// @Router       /auth/login [post]
func Login(c *gin.Context) {
	var req auth.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.AuthClient.Login(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// Register godoc
// @Summary      register
// @Description  register
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        username body string true "username"
// @Param        password body string true "password"
// @Success      200 {object} auth.RegisterResponse
// @Failure      500 {object} auth.RegisterResponse
// @Router       /auth/register [post]
func Register(c *gin.Context) {
	var req auth.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.AuthClient.Register(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// QueryInfo godoc
// @Summary      get auth info
// @Description  get auth info
// @Tags         auth
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Success      200 {object} biz.QueryInfoResponse
// @Failure      500 {object} biz.QueryInfoResponse
// @Router       /auth [get]
func QueryInfo(c *gin.Context) {
	var req biz.QueryInfoRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.BizClient.QueryInfo(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}
