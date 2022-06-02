package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

func TestQuery(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../../")
	dal.InitMongoDB()
	//QueryVideoById(context.Background(), 1)
	videos, err := dal.QueryVideosByUserId(context.Background(), 1)
	if err != nil {
		log.Printf("error to query %v\n", err)
		return
	}
	for _, video := range videos {
		fmt.Printf("%+v\n", *video)
	}
}
