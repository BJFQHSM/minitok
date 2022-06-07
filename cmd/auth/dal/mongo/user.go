package mongo

import (
	"context"

	"github.com/bytedance2022/minimal_tiktok/pkg/util"
)

type User struct {
	UserId       int64   `bson:"user_id"`
	Username     string  `bson:"username"`
	Follows      []int64 `bson:"follows"`       // 关注列表
	Followers    []int64 `bson:"followers"`     // 粉丝列表
	PublishList  []int64 `bson:"publish_list"`  // 发布视频列表
	FavoriteList []int64 `bson:"favorite_list"` // 点赞列表
}

func InsertUser(ctx context.Context, user *User) error {
	util.LogInfof("start to insert into mongo %+v\n", user)
	coll := Cli.Database("tiktok").Collection("user")
	_, err := coll.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	util.LogInfof("success to insert into mongo %+v\n", user)
	return nil
}
