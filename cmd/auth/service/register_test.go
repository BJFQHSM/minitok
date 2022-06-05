package service

import (
	"context"
	"fmt"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mysql"
	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
	"log"
	"os"
	"regexp"
	"testing"
	"time"
)

func TestRegister(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	mysql.InitMysql()
	s := &registerServiceImpl{
		Req: &auth.RegisterRequest{
			Username: "testUser",
			Password: "testPassword",
		},
		Ctx: context.Background(),
	}
	err = s.doRegister()
	if err != nil {
		log.Fatal(err)
	}
}

func TestLogin(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	mysql.InitMysql()
	s := &loginServiceImpl{
		Req: &auth.LoginRequest{
			Username: "testUser",
			Password: "testPassword",
		},
		Ctx: context.Background(),
	}
	err = s.doLogin()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", s.user)
}

func TestTokenGenerate(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	mysql.InitMysql()
	s := &loginServiceImpl{
		Req: &auth.LoginRequest{
			Username: "testUser",
			Password: "testPassword",
		},
		Ctx: context.Background(),
	}
	err = s.doLogin()
	if err != nil {
		log.Fatal(err)
	}
	token, err := dal.JwtGenerateToken(s.user.UserId, time.Hour)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(token)

	user, err := dal.JwtParseUser(token)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v\n", user)
}

func TestReg(t *testing.T) {
	match, err := regexp.Match("^[\u4E00-\u9FA5A-Za-z\\d]{6,10}$", []byte("usernametoolong"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(match)
}
