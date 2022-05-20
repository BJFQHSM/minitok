package db

import (
	"context"
	"log"
)

type User struct {
	UserId            int    `json:"user_id"`
	Username          string `json:"user_name"`
	EncryptedPassword string `json:"password"`
	FollowCount       int    `json:"follow_count"`
	FollowerCount     int    `json:"follower_count"`
}

type Follow struct {
	UserId       int `json:"user_id"`
	FollowUserId int `json:"follow_user_id "`
}

func QueryUserByUID(ctx context.Context, userId int64) (*User, error) {
	var res User
	log.Print(userId)
	if err := MysqlDB.Table("user_info").WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		log.Printf("Erorr to queryUserInfo %v\n", err)
		return nil, err
	}
	return &res, nil
}

func QueryFollowUserByUID(ctx context.Context, tokenUserId int64, userId int64) (bool, error) {
	var follow Follow
	if err := MysqlDB.Table("user_follow").WithContext(ctx).Where(&Follow{int(tokenUserId), int(userId)}).Find(&follow).Error; err != nil {
		log.Printf("Erorr to queryUserInfo %v\n", err)
		return false, err
	}
	if (Follow{}) == follow {
		return false, nil
	}
	return true, nil
}
