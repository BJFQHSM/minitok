package dal

import (
	"context"
)

type User struct {
	Id         int64  `gorm:"id"`
	UserId     int64  `gorm:"user_id"`
	Username   string `gorm:"username"`
	EncryptPwd string `gorm:"encrypt_pwd"`
}

func QueryUserByUsername(ctx context.Context, username string) (*User, error) {
	var user User
	err := MysqlDB.Select("username", username).Find(&user).Error
	if err != nil {
		return &user, err
	}
	return &user, nil
}

func Register(ctx context.Context, user User, info EncryptInfo) error {
	tx := MysqlDB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Create(info).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}
