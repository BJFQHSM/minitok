package dal

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	UserId        int64   `bson:"user_id"`
	Username      string  `bson:"username"`
	FollowCount   int64   `bson:"follow_count"`   // 关注数
	FollowerCount int64   `bson:"follower_count"` // 粉丝数
	Follows       []int64 `bson:"follows"`        // 关注列表
	Followers     []int64 `bson:"followers"`      // 粉丝列表
	PublishList   []int64 `bson:"publish_list"`   // 发布视频列表
	FavoriteList  []int64 `bson:"favorite_list"`  // 点赞列表
}

func ChangeFollowRelation(ctx context.Context, followee int64, follower int64) error {
	userColl := MongoCli.Database("tiktok").Collection("user")

	// 定义事务
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// 重复关注校验
		err := userColl.FindOne(sessCtx, bson.D{{"user_id", followee}, {"followers", bson.D{{"$all", bson.A{follower}}}}}).Err()
		if err == nil {
			return nil, errors.New("follow again")
		}
		if err != mongo.ErrNoDocuments {
			log.Printf("%v\n", err)
			return nil, err
		}

		filter := bson.D{{"user_id", followee}}
		update := bson.D{
			{"$inc", bson.D{{"follower_count", 1}}},
			{"$addToSet", bson.D{{"followers", follower}}},
		}
		if updateResult, err := userColl.UpdateOne(sessCtx, filter, update); err != nil {
			return nil, err
		} else if updateResult.MatchedCount == 0 {
			return nil, errors.New("no followee was found")
		}
		filter = bson.D{{"user_id", follower}}
		update = bson.D{
			{"$inc", bson.D{{"follow_count", 1}}},
			{"$addToSet", bson.D{{"follows", followee}}},
		}
		if updateResult, err := userColl.UpdateOne(sessCtx, filter, update); err != nil {
			return nil, err
		} else if updateResult.MatchedCount == 0 {
			return nil, errors.New("no follower was found")
		}
		return nil, nil
	}

	// 开启会话
	session, err := MongoCli.StartSession()
	if err != nil {
		log.Printf("ERROR: fail to start mongo session. %v\n", err)
		return err
	}
	defer session.EndSession(ctx)

	// 执行事务
	_, err = session.WithTransaction(ctx, callback)
	if err != nil {
		log.Printf("ERROR: fail to ChangeFollowRelation. %v\n", err)
		return err
	}
	return nil
}
