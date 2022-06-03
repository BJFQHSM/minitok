package dal

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
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
	Title         string    `bson:"title"`
}

type Comment struct {
	CommentId  int64     `bson:"comment_id"`
	UserId     int64     `bson:"user_id"`
	Content    string    `bson:"content"`
	CreateDate time.Time `bson:"create_date"`
}

func QueryVideosByTime(t time.Time) ([]*Video, error) {
	// 指定获取要操作的数据集
	collection := MongoCli.Database("tiktok").Collection("video")
	findOptions := options.Find()
	findOptions.SetLimit(30) //设置一次获取的最大视频数
	sort := bson.D{{"publish_date", 1}}
	findOptions.SetSort(sort)
	results := []*Video{}
	cur, err := collection.Find(context.TODO(), bson.M{"publish_date": bson.M{"$gte": t}}, findOptions)

	// 完成后关闭游标
	defer cur.Close(context.TODO())

	if err != nil {
		log.Fatal(err)
	}
	for cur.Next(context.TODO()) {
		var elem Video
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, &elem)
	}
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	return results, nil
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

// QueryVideosByUserId2 deprecated
//func QueryVideosByUserId2(ctx context.Context, userId int64) ([]*Video, error) {
//	videoColl := MongoCli.Database("tiktok").Collection("video")
//	cursor, err := videoColl.Find(ctx, bson.D{{"user_id", userId}})
//	if err != nil {
//		return []*Video{}, err
//	}
//	var results []bson.D
//	if err = cursor.All(ctx, &results); err != nil {
//		return []*Video{}, err
//	}
//	var videos []*Video
//	for _, result := range results {
//		marshal, err := bson.Marshal(result)
//		if err != nil {
//			log.Printf("error to marshal from result %v\n", err)
//			return []*Video{}, err
//		}
//		var video Video
//		err = bson.Unmarshal(marshal, &video)
//		if err != nil {
//			log.Printf("error to unmarshal from result %v\n", err)
//			return []*Video{}, err
//		}
//		videos = append(videos, &video)
//	}
//	return videos, nil
//}

// QueryVideosByUserId
// favorite_list(favorites filed) only show login userId
// if not favor favorites is empty
func QueryVideosByUserId(ctx context.Context, userId int64) ([]*Video, error) {
	videoColl := MongoCli.Database("tiktok").Collection("video")

	matchStage := bson.D{{"$match", bson.D{{"user_id", userId}}}}

	projectStage := bson.D{
		{
			"$project",
			bson.M{
				"_id":            0,
				"user_id":        1,
				"video_id":       1,
				"play_url":       1,
				"cover_url":      1,
				"favorite_count": 1,
				"comment_count":  1,
				"comments":       1,
				"publish_date":   1,
				// only return element which equals user_id
				"favorites": bson.D{{"$filter", bson.D{{"input", "$favorites"}, {"as", "f"}, {"cond", bson.M{"$eq": bson.A{"$$f", userId}}}}}},
			},
		},
	}

	cursor, err := videoColl.Aggregate(ctx, mongo.Pipeline{matchStage, projectStage})
	defer cursor.Close(ctx)
	if err != nil {
		return []*Video{}, err
	}
	var results []bson.M
	if err = cursor.All(ctx, &results); err != nil {
		return []*Video{}, err
	}
	var videos []*Video
	for _, result := range results {
		//fmt.Println(result)
		marshal, err := bson.Marshal(result)
		if err != nil {
			log.Printf("error to marshal from result %v\n", err)
			return []*Video{}, err
		}
		var video Video
		err = bson.Unmarshal(marshal, &video)
		if err != nil {
			log.Printf("error to unmarshal from result %v\n", err)
			return []*Video{}, err
		}
		videos = append(videos, &video)
	}
	return videos, nil
}
