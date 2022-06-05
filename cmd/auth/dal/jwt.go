package dal

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
	//"github.com/sirupsen/logrus"
)

var AppSecret = ""     //viper.GetString会设置这个值(32byte长度)
var AppIss = "minitok" //这个值会被viper.GetString重写

type userStdClaims struct {
	jwt.StandardClaims
	userId int64
}

// Valid 自定义校验内容
func (c userStdClaims) Valid() (err error) {
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}
	if !c.VerifyIssuer(AppIss, true) {
		return errors.New("token's issuer is wrong")
	}
	if c.userId < 0 {
		return errors.New("invalid user in jwt")
	}
	return
}

func JwtGenerateToken(userId int64, d time.Duration) (string, error) {
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", userId),
		Issuer:    AppIss,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		userId:         userId,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		log.Println("config is wrong, can not generate jwt")
	}
	return tokenString, err
}

//JwtParseUser 解析payload的内容,得到userID
func JwtParseUser(tokenString string) (int64, error) {
	if tokenString == "" {
		return -1, errors.New("no token is found in Authorization Bearer")
	}
	claims := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return -1, nil
	})
	if err != nil {
		return -1, err
	}
	return claims.userId, err
}
