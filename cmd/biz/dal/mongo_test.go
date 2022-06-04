package dal

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"os"
	"testing"
)

func TestQuery(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	os.Setenv("env", "dev")
	InitMongoDB()
	//QueryVideoById(context.Background(), 1)
	filter := bson.M{"video_id": 1}
	update := bson.M{
		"$addToSet": bson.M{"favorites": 1},
	}
	updateResult, err := MongoCli.Database("tiktok").Collection("video").UpdateOne(context.TODO(), filter, update)
	if err != nil || updateResult.MatchedCount == 0 {
		log.Printf("error to modify %v\n", err)
		return
	}
	videos, err := QueryVideosByUserId(context.Background(), 1)
	if err != nil {
		log.Printf("error to query %v\n", err)
		return
	}
	for _, video := range videos {
		fmt.Printf("%+v\n", *video)
	}
}
