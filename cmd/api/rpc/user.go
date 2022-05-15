package rpc

import (
	"context"
	etcd "github.com/kitex-contrib/registry-etcd"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"grpc_test/grpc_gen"
	"log"
)

var UserClient grpc_gen.UserServiceClient

func InitUser() {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	resolver, err := etcd.NewEtcdResolver([]string{"127.0.0.1:2379"})
	if err != nil {
		log.Panic(err)
	}
	resolve, err := resolver.Resolve(context.Background(), "test_user")
	if err != nil {
		return
	}
	if len(resolve.Instances) == 0 {
		return

	}
	addr := resolve.Instances[0].Address().String()
	//conn, err := grpc.Dial("127.0.0.1:8888", opts...)
	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
	}
	UserClient = grpc_gen.NewUserServiceClient(conn)
}