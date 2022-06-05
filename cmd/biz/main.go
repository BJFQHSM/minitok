package main

import (
	"fmt"
	"log"
	"net"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/rpc"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"google.golang.org/grpc"
)

func main() {
	// todo constants and others
	dal.Init()
	rpc.Init()

	lis, err := net.Listen("tcp", fmt.Sprintf(":8889"))
	if err != nil {
		panic(err)
	}
	srv := grpc.NewServer()
	biz.RegisterBizServiceServer(srv, &BizServerImpl{})
	if err != nil {
		log.Panic(err)
	}
	srv.Serve(lis)
}
