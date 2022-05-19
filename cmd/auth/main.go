package main

import (
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// todo constants and others
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8888"))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	auth.RegisterAuthServiceServer(srv, &AuthServiceImpl{})
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	err = r.Register(&registry.Info{
		ServiceName: "tiktok_auth",
		Addr: &net.TCPAddr{
			IP:   []byte{127, 0, 0, 1},
			Port: 8888,
		},
	})
	if err != nil {
		log.Panic(err)
	}
	srv.Serve(lis)
}
