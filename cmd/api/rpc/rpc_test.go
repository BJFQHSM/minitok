package rpc

import (
	"context"
	"fmt"
	"go.etcd.io/etcd/client/v3"
	"log"
	"testing"
)

func TestRpc(t *testing.T) {
	cli, err := clientv3.NewFromURL("http://localhost:2379")
	get, err := cli.Get(context.Background(), "kitex/registry-etcd/tiktok_biz/127.0.0.1:8889", clientv3.WithIgnoreLease())
	if err != nil {
		log.Println(err)
	}
	fmt.Printf("%+v\n", get)
}
