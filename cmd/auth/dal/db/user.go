package db

import (
	"context"
	"log"
)

type User struct {
	UserId        int    `json:"user_id"`
	Username      string `json:"user_name"`
	Password      string `json:"password"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
}

func QueryUserByUID(ctx context.Context, userId int64) (*User, error) {
	var res User
	if err := MysqlDB.WithContext(ctx).Where("user_id = ?", userId).Find(&res).Error; err != nil {
		log.Printf("Erorr to queryUserInfo %v\n", err)
		return nil, err
	}
	return &res, nil
}
