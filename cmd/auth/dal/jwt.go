package dal

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mysql"

	"github.com/dgrijalva/jwt-go"
	//"github.com/sirupsen/logrus"
)

var AppSecret = ""     //viper.GetString会设置这个值(32byte长度)
var AppIss = "minitok" //这个值会被viper.GetString重写

type userStdClaims struct {
	jwt.StandardClaims
	*mysql.User
}

//实现 `type Claims interface` 的 `Valid() error` 方法,自定义校验内容
func (c userStdClaims) Valid() (err error) {
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}
	if !c.VerifyIssuer(AppIss, true) {
		return errors.New("token's issuer is wrong")
	}
	if c.User.UserId < 1 {
		return errors.New("invalid user in jwt")
	}
	return
}

func JwtGenerateToken(m *mysql.User, d time.Duration) (string, error) {
	m.EncryptPwd = ""
	expireTime := time.Now().Add(d)
	stdClaims := jwt.StandardClaims{
		ExpiresAt: expireTime.Unix(),
		IssuedAt:  time.Now().Unix(),
		Id:        fmt.Sprintf("%d", m.Id),
		Issuer:    AppIss,
	}

	uClaims := userStdClaims{
		StandardClaims: stdClaims,
		User:           m,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, uClaims)
	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(AppSecret))
	if err != nil {
		log.Fatal("config is wrong, can not generate jwt")
	}
	return tokenString, err
}

//JwtParseUser 解析payload的内容,得到用户信息
//gin-middleware 会使用这个方法
func JwtParseUser(tokenString string) (*mysql.User, error) {
	if tokenString == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := userStdClaims{}
	_, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(AppSecret), nil
	})
	if err != nil {
		return nil, err
	}
	return claims.User, err
}
