package dal

import (
	"context"
	"log"
	"os"
	"testing"
	"time"
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
	users, err := QueryFollowsByUserId(context.Background(), 1252117233)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}

	for key, user := range users {
		log.Printf("%d 个user为:%+v", key, user)
	}
}
func TestPublishVideo(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Printf("ERROR: fail to get current dir %v\n", err)
		return
	}
	os.Setenv("WORK_DIR", pwd+"/../../../")
	InitMongoDB()
	video := &Video{
		VideoId:       3597527201278699950,
		UserId:        1675229147,
		PlayUrl:       "http:127.0.0.1:8080/static/QzSmKHaDFmYZruEImvRj.mp4",
		FavoriteCount: 0,
		CommentCount:  0,
		PublishDate:   time.Now(),
		Title:         "12345",
	}
	err = PublishVideo(context.TODO(), 1675229147, video)
	if err != nil {
		log.Printf("%v\n", err)
		return
	}
}
