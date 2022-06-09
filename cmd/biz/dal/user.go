package dal

import (
	"context"
	"errors"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// 关注操作
func FollowRelation(ctx context.Context, followee int64, follower int64) error {
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
			// {"$inc", bson.D{{"follower_count", 1}}},
			{"$addToSet", bson.D{{"followers", follower}}},
		}
		if updateResult, err := userColl.UpdateOne(sessCtx, filter, update); err != nil {
			return nil, err
		} else if updateResult.MatchedCount == 0 {
			return nil, errors.New("no followee was found")
		}
		filter = bson.D{{"user_id", follower}}
		update = bson.D{
			// {"$inc", bson.D{{"follow_count", 1}}},
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
		log.Printf("ERROR: fail to FollowRelation. %v\n", err)
		return err
	}
	return nil
}

// 赞操作
func FavoriteAction(ctx context.Context, user_id int64, video_id int64, actionType int32) error {
	userColl := MongoCli.Database("tiktok").Collection("user")
	videoColl := MongoCli.Database("tiktok").Collection("video")
	// 定义事务
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		//根据id获取对应的用户和视频数据
		filter := bson.D{{"user_id", user_id}}
		var user User
		err := userColl.FindOne(ctx, filter).Decode(&user)
		if err != nil {
			log.Println(err)
			return nil, errors.New("user_id not exist")
		}
		filter = bson.D{{"video_id", video_id}}
		var video Video
		err = videoColl.FindOne(ctx, filter).Decode(&video)
		if err != nil {
			log.Println(err)
			return nil, errors.New("video_id not exist")
		}

		if actionType == 1 { //点赞
			err = Favorite(ctx, user_id, video_id)
			if err != nil {
				log.Println(err)
				return nil, errors.New("点赞报错")
			}
		} else { //取消点赞
			err = CancelFavorite(ctx, user_id, video_id)
			if err != nil {
				log.Println(err)
				return nil, errors.New("取消点赞报错")
			}
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
		log.Printf("ERROR: fail to favorite. %v\n", err)
		return err
	}
	return nil
}

// 取消关注
func UnFollowRelation(ctx context.Context, followee int64, follower int64) error {
	userColl := MongoCli.Database("tiktok").Collection("user")

	// 定义事务
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		// 是否关注
		err := userColl.FindOne(sessCtx, bson.D{{"user_id", followee}, {"followers", bson.D{{"$all", bson.A{follower}}}}}).Err()
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, errors.New("the user is not followed")
			}
			log.Printf("%v\n", err)
			return nil, err
		}

		filter := bson.D{{"user_id", followee}}
		update := bson.D{
			// {"$inc", bson.D{{"follower_count", -1}}},
			{"$pull", bson.D{{"followers", follower}}},
		}
		if updateResult, err := userColl.UpdateOne(sessCtx, filter, update); err != nil {
			return nil, err
		} else if updateResult.MatchedCount == 0 {
			return nil, errors.New("no followee was found")
		}
		filter = bson.D{{"user_id", follower}}
		update = bson.D{
			// {"$inc", bson.D{{"follow_count", -1}}},
			{"$pull", bson.D{{"follows", followee}}},
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
		log.Printf("ERROR: fail to UnFollowRelation. %v\n", err)
		return err
	}
	return nil
}

//获取用户点赞列表
func GetFavoriteList(ctx context.Context, user_id int64) ([]*Video, error) {
	user, err := QueryUserById(ctx, user_id)
	if err != nil {
		log.Println(err)
		return []*Video{}, err
	}
	res := []*Video{}
	for i := 0; i < len(user.FavoriteList); i++ {
		item, err := QueryVideoByVideoId(ctx, user.FavoriteList[i])
		if err != nil {
			log.Printf("%+v\n", err)
			return nil, err
		}
		res = append(res, item)
	}
	return res, nil
}

// 凭用户Id查询用户信息
func QueryUserById(ctx context.Context, id int64) (*User, error) {
	coll := MongoCli.Database("tiktok").Collection("user")
	var user *User
	err := coll.FindOne(ctx, bson.D{{"user_id", id}}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no such user")
		}
		log.Printf("%+v", err)
		return nil, err
	}
	user.FollowCount = int64(len(user.Follows))
	user.FollowerCount = int64(len(user.Followers))
	return user, nil
}

// 查询当前用户是否关注该用户
func QueryIsFollow(ctx context.Context, followeeId int64, followerId int64) (bool, error) {
	coll := MongoCli.Database("tiktok").Collection("user")
	var follower *User
	err := coll.FindOne(ctx, bson.D{{"user_id", followerId}}).Decode(&follower)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return false, errors.New("no user was found")
		}
		log.Printf("%+v", err)
		return false, err
	}
	isFollow := false
	follows := follower.Follows
	for _, userId := range follows {
		if userId == followeeId {
			isFollow = true
			return isFollow, nil
		}
	}
	return isFollow, nil
}

// 查询关注
func QueryFollowsByUserId(ctx context.Context, userId int64) ([]*User, error) {
	userColl := MongoCli.Database("tiktok").Collection("user")
	filter := bson.D{{"user_id", userId}}
	pro := bson.D{{"_id", 0}}
	opts := options.FindOne().SetProjection(pro)

	var user User
	err := userColl.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no such user")
		}
		log.Printf("%+v", err)
		return nil, err
	}
	if len(user.Follows) == 0 {
		return []*User{}, nil
	}
	var re []*User
	for _, followeeId := range user.Follows {
		followee, err := QueryUserById(ctx, followeeId)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
		re = append(re, followee)
	}
	return re, nil
}

// 查询粉丝
func QueryFollowersByUserId(ctx context.Context, userId int64) ([]*User, error) {
	userColl := MongoCli.Database("tiktok").Collection("user")
	filter := bson.D{{"user_id", userId}}
	pro := bson.D{{"_id", 0}}
	opts := options.FindOne().SetProjection(pro)

	var user User
	err := userColl.FindOne(ctx, filter, opts).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("no such user")
		}
		log.Printf("%+v", err)
		return nil, err
	}
	if len(user.Followers) == 0 {
		return []*User{}, nil
	}
	var re []*User
	for _, followerId := range user.Followers {
		follower, err := QueryUserById(ctx, followerId)
		if err != nil {
			log.Printf("%+v", err)
			return nil, err
		}
		re = append(re, follower)
	}
	return re, nil
}

//点赞
func Favorite(ctx context.Context, userId, videoId int64) error {
	collection := MongoCli.Database("tiktok").Collection("user")
	query := bson.M{"user_id": userId}
	update := bson.M{"$push": bson.M{"favorite_list": videoId}}
	_, err := collection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	collection = MongoCli.Database("tiktok").Collection("video")
	//更新点赞列表
	query = bson.M{"video_id": videoId}
	update = bson.M{"$push": bson.M{"favorites": userId}}
	_, err = collection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}
	return nil
}

//取消点赞
func CancelFavorite(ctx context.Context, userId, videoId int64) error {
	collection := MongoCli.Database("tiktok").Collection("user")
	query := bson.M{"user_id": userId}
	update := bson.M{"$pull": bson.M{"favorite_list": videoId}}
	_, err := collection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}

	collection = MongoCli.Database("tiktok").Collection("video")
	//更新点赞列表
	query = bson.M{"video_id": videoId}
	update = bson.M{"$pull": bson.M{"favorites": userId}}
	_, err = collection.UpdateOne(ctx, query, update)
	if err != nil {
		return err
	}
	return nil
}

func PublishVideo(ctx context.Context, userId int64, video *Video) error {
	userColl := MongoCli.Database("tiktok").Collection("user")
	videoColl := MongoCli.Database("tiktok").Collection("video")

	// 定义事务
	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		filter := bson.D{{"user_id", userId}}
		update := bson.D{
			{"$addToSet", bson.M{"publish_list": video.VideoId}},
		}
		if updateResult, err := userColl.UpdateOne(sessCtx, filter, update); err != nil {
			return nil, err
		} else if updateResult.MatchedCount == 0 {
			return nil, errors.New("no user was found")
		}

		if _, err := videoColl.InsertOne(sessCtx, video); err != nil {
			return nil, err
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
		log.Printf("ERROR: fail to UnFollowRelation. %v\n", err)
		return err
	}
	return nil
}
