package main

import (
	"context"
	"fmt"
	"github.com/Bytedance2022/minimal_tiktok/db/mongodb"
	"github.com/Bytedance2022/minimal_tiktok/db/mysql"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"testing"
)

func TestInitMysql(t *testing.T) {
	mysql.InitMysql()
}

func TestInitMongo(t *testing.T) {
	mongodb.InitMongoDB()

	var result bson.D
	err := mongodb.MongoCli.Database("myFirstDatabase").Collection("mycollection").FindOne(context.TODO(), bson.D{}).Decode(&result)
	if err != nil {
		log.Fatal("error: ", err)
	}
	fmt.Printf("%+v", result)
}
