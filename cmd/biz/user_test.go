package main

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
	"github.com/bytedance2022/minimal_tiktok/cmd/biz/service"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/biz"
)

func TestQueryUserInfo(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../")
	dal.InitMongoDB()
	req := &biz.QueryUserInfoRequest{
		UserId: 1,
		Token:  "1",
	}
	ser := service.NewQueryUserInfoService(req, nil)
	ser.DoService()
	fmt.Printf("%+v", ser.GetResponse())
}
