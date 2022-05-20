package main

import (
	"fmt"
	"testing"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

func TestQueryUserInfo(t *testing.T) {
	req := &auth.QueryUserInfoRequest{
		UserId: 2,
		Token:  "1.12345",
	}
	ser := service.NewQueryUserInfoService(req, nil)
	ser.DoService()
	fmt.Printf("%+v", ser.GetResponse())
}
