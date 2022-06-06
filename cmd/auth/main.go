package main

import (
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// todo constants and others
	dal.Init()
	lis, err := net.Listen("tcp", fmt.Sprintf(":8890"))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	auth.RegisterAuthServiceServer(srv, &AuthServiceImpl{})
	if err != nil {
		log.Panic(err)
	}
	srv.Serve(lis)
}
