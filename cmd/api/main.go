package main

import (
	"context"
	"github.com/bytedance2022/minimal_tiktok/cmd/api/rpc"
	"github.com/bytedance2022/minimal_tiktok/kitex_gen/user"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	rpc.InitUserRpc()
	r := gin.New()

	// todo routines
	v1 := r.Group("/v1")
	user1 := v1.Group("/user")
	user1.GET("/login", login)

	if err := http.ListenAndServe(":8080", r); err != nil {
		klog.Fatal(err)
	}
}

// test
func login(c * gin.Context) {
	response, err := rpc.UserClient.Login(context.Background(), &user.LoginRequest{})
	if err != nil {
		return
	}
	c.JSON(http.StatusOK, response)
}

