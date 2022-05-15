package rpc

import (
	"github.com/bytedance2022/minimal_tiktok/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/retry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
)

var UserClient userservice.Client

func InitUserRpc() {
	r, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		panic(err)
	}

	c, err := userservice.NewClient(
		"tiktok_user",
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(3*time.Second),              // rpc timeout
		client.WithConnectTimeout(50*time.Millisecond),    // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
	)
	if err != nil {
		panic(err)
	}
	UserClient = c
}