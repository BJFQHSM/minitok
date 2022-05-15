package handler

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Login godoc
// @Summary      login
// @Description  login
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        username body string true "username"
// @Param        password body string true "password"
// @Success      200 {object} user.LoginResponse
// @Failure      500 {object} user.LoginResponse
// @Router       /user/login [post]
func Login(c *gin.Context) {
	var req user.LoginRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.UserClient.Login(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// Register godoc
// @Summary      register
// @Description  register
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        username body string true "username"
// @Param        password body string true "password"
// @Success      200 {object} user.RegisterResponse
// @Failure      500 {object} user.RegisterResponse
// @Router       /user/register [post]
func Register(c *gin.Context) {
	var req user.RegisterRequest
	err := c.ShouldBind(&req)
	if err != nil {
		// todo
	}
	resp, err := rpc.UserClient.Register(context.Background(), &req)
	if err != nil {
		// todo
	}
	c.JSON(http.StatusOK, resp)
}

// QueryInfo godoc
// @Summary      get user info
// @Description  get user info
// @Tags         user
// @Accept       json
// @Produce      json
// @Param        user_id body int true "user_id"
// @Param        token body string true "token"
// @Success      200 {object} biz.QueryInfoResponse
// @Failure      500 {object} biz.QueryInfoResponse
// @Router       /user [get]
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
