package main

import (
	"fmt"
	"testing"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

func TestQueryUserInfo(t *testing.T) {
	req := &biz.QueryUserInfoRequest{
		UserId: 0,
		Token:  "1.12345",
	}
	ser := service.NewQueryUserInfoService(req, nil)
	ser.DoService()
	fmt.Printf("%+v", ser.GetResponse())
}
