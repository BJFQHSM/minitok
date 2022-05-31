package main

import (
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal/db"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	// todo constants and others
	db.InitMongoDB()
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
