package main

import (
	"context"
	"fmt"
	"github.com/RaymondCode/simple-demo/db/mongodb"
	"github.com/RaymondCode/simple-demo/db/mysql"
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
