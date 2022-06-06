package service

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/bytedance2022/minimal_tiktok/cmd/biz/dal"
)

func TestDalUserToBizUser(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	dal.InitMongoDB()
	user := &dal.User{
		UserId:        1252117233,
		Username:      "user1",
		FollowCount:   1,
		FollowerCount: 1,
	}
	var users []*dal.User
	users = append(users, user)
	re, err := DalUserToBizUser(context.TODO(), users, 1362738155)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
	log.Printf("%+v", re[0])
}
