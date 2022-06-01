package dal

import (
	"context"
	"log"
	"os"
	"testing"
)

func TestFollowRelation(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	InitMongoDB()
	err = FollowRelation(context.Background(), 1, 2)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}

func TestUnFollowRelation(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	InitMongoDB()
	err = UnFollowRelation(context.Background(), 1, 2)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}

func TestQueryFollowsByUserId(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	InitMongoDB()
	users, err := QueryFollowsByUserId(context.Background(), 2)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	for key, user := range users {
		log.Printf("%d 个user为:%+v", key, user)
	}
}
