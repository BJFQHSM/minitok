package main

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/registry"
	etcd "github.com/kitex-contrib/registry-etcd"
	"google.golang.org/grpc"
	"grpc_test/grpc_gen"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:8888"))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	grpc_gen.RegisterUserServiceServer(srv, &UserServiceImpl{})
	r, err := etcd.NewEtcdRegistry([]string{"127.0.0.1:2379"})
	err = r.Register(&registry.Info{
		ServiceName: "test_user",
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