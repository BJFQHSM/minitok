package dal

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Video struct {
	VideoId       int64     `bson:"video_id"`
	UserId        int64     `bson:"user_id"`
	PlayUrl       string    `bson:"play_url"`
	CoverUrl      string    `bson:"cover_url"`
	FavoriteCount int64     `bson:"favorite_count"`
	Favorites     []int64   `bson:"favorites"`
	CommentCount  int64     `bson:"comment_count"`
	Comments      []Comment `bson:"comments, inline"`
	PublishDate   time.Time `bson:"publish_date"`
}

type Comment struct {
	CommentId  int64     `bson:"comment_id"`
	UserId     int64     `bson:"user_id"`
	Content    string    `bson:"content"`
	CreateDate time.Time `bson:"create_date"`
}

func QueryVideoByVideoId(ctx context.Context, videoId int64) (*Video, error) {
	videoColl := MongoCli.Database("tiktok").Collection("video")
	var result bson.D
	opts := options.FindOne().SetProjection(bson.D{{"_id", 0}})
	err := videoColl.FindOne(ctx, bson.D{{"video_id", videoId}}, opts).Decode(&result)
	if err != nil {
		log.Printf("Erorr to queryVideoById %v\n", err)
		return nil, err
	}
	marshal, err := bson.Marshal(result)
	if err != nil {
		log.Printf("error to marshal from result %v\n", err)
		return nil, err
	}
	var video Video
	err = bson.Unmarshal(marshal, &video)
	if err != nil {
		log.Printf("error to unmarshal from result %v\n", err)
		return nil, err
	}
	return &video, nil
}
