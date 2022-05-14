package do

import (
	"context"
	"github.com/Bytedance2022/minimal_tiktok/db/mongodb"
	"log"
	"testing"
)

func TestInsertUser(t *testing.T) {
	mongodb.InitMongoDB()
	cli := mongodb.MongoCli
	coll := cli.Database("tiktok").Collection("user")
	user := User{
		Id:   1,
		Name: "testuser",
	}
	_, err := coll.InsertOne(context.TODO(), &user)
	if err != nil {
		log.Fatal(err)
	}
}
