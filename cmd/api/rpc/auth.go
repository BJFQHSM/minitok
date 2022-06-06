package rpc

import (
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

var AuthClient auth.AuthServiceClient

func initAuth() {
	util.LogInfo("AuthClient initiation starting...")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//resolver, err := etcd.NewEtcdResolver([]string{"etcd:2379"})
	//if err != nil {
	//	log.Panic(err)
	//}
	//resolve, err := resolver.Resolve(context.Background(), "tiktok_auth")
	//if err != nil {
	//	return
	//}
	//if len(resolve.Instances) == 0 {
	//	return
	//
	//}
	//addr := resolve.Instances[0].Address().String()
	isDev := os.Getenv("env") == "dev"
	var conn *grpc.ClientConn
	var err error
	if isDev {
		conn, err = grpc.Dial("127.0.0.1:8890", opts...)
	} else {
		conn, err = grpc.Dial("auth:8890", opts...)
	}
	if AuthClient = auth.NewAuthServiceClient(conn); err != nil || AuthClient == nil {
		util.LogFatalf("AuthClient dial error : %+v\n", err)
	}

	util.LogInfo("AuthClient initiate success!")
}
