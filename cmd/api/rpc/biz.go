package rpc

import (
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"os"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var BizClient biz.BizServiceClient

func initBiz() {
	util.LogInfo("BizClient initiation starting...")
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	//resolver, err := etcd.NewEtcdResolver([]string{"etcd:2379"})
	//if err != nil {
	//	log.Panic(err)
	//}
	//resolve, err := resolver.Resolve(context.Background(), "tiktok_biz")
	//if err != nil {
	//	return
	//}
	//if len(resolve.Instances) == 0 {
	//	return
	//
	//}
	//addr := resolve.Instances[0].Address().String()
	//conn, err := grpc.Dial("127.0.0.1:8888", opts...)
	var conn *grpc.ClientConn
	var err error
	isDev := os.Getenv("env") == "dev"
	if isDev {
		conn, err = grpc.Dial("127.0.0.1:8889", opts...)
	} else {
		conn, err = grpc.Dial("biz:8889", opts...)
	}
	if BizClient = biz.NewBizServiceClient(conn); err != nil || BizClient == nil {
		util.LogFatalf("BizClient dial error : %+v\n", err)
	}

	util.LogInfo("BizClient initiate success!")
}
