package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mongo"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mysql"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"regexp"
	"time"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type RegisterService interface {
	DoService() *auth.RegisterResponse
}

func NewRegisterService(ctx context.Context, r *auth.RegisterRequest) RegisterService {
	return &registerServiceImpl{Req: r, Ctx: ctx, Resp: &auth.RegisterResponse{}}
}

type registerServiceImpl struct {
	Req  *auth.RegisterRequest
	Resp *auth.RegisterResponse
	Ctx  context.Context

	encryptSalt string
	encryptPwd  string
	user        *mysql.User
	token       string
}

func (s *registerServiceImpl) DoService() *auth.RegisterResponse {
	// mock
	msg := "success"
	s.Resp = &auth.RegisterResponse{
		UserId:     1,
		Token:      "fsjfs",
		StatusCode: 0,
		StatusMsg:  &msg,
	}

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doRegister(); err != nil {
			break
		}

		if err = s.generateToken(); err != nil {

		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *registerServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if match, err := regexp.Match("^[\u4E00-\u9FA5A-Za-z\\d]{1,10}$", []byte(req.Username)); err != nil {
		return errors.New("fail to validate username")
	} else if !match {
		return errors.New("username can only contains nums, English letters and Chinese")
	}
	if match, err := regexp.Match("^[A-Za-z\\d]{8,20}$", []byte(req.Password)); err != nil {
		return errors.New("fail to validate password")
	} else if !match {
		return errors.New("password can only contains nums, English letters and underline character")
	}
	return nil
}

func (s *registerServiceImpl) doRegister() error {
	userId := util.GenerateRandomInt32()
	s.encrypt()
	user := mysql.User{
		UserId:     int64(userId),
		Username:   s.Req.Username,
		EncryptPwd: s.encryptPwd,
	}
	info := mysql.EncryptInfo{
		UserId: int64(userId),
		Salt:   s.encryptSalt,
	}
	s.user = &user
	var err error
	// todo
	if err = mysql.InsertUser(s.Ctx, user, info); err != nil {
		return err
	}
	bizUser := &mongo.User{
		UserId:        int64(userId),
		Username:      user.Username,
		FollowCount:   0,
		Follows:       []int64{},
		FollowerCount: 0,
		Followers:     []int64{},
		PublishList:   []int64{},
		FavoriteList:  []int64{},
	}
	if err = mongo.InsertUser(s.Ctx, bizUser); err != nil {
		return err
	}
	return nil
}

func (s *registerServiceImpl) generateToken() error {
	token, _ := dal.JwtGenerateToken(s.user.UserId, 24*time.Hour)
	s.token = token
	return nil
}

func (s *registerServiceImpl) encrypt() {
	s.encryptSalt = util.GenerateRandomStr(10)
	s.encryptPwd = util.MD5Encrypt(s.Req.Password, s.encryptSalt)
}

func (s *registerServiceImpl) buildResponse(err error) {
	if err != nil {
		errMsg := err.Error()
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 1
	} else {
		errMsg := "SUCCESS"
		s.Resp.StatusMsg = &errMsg
		s.Resp.StatusCode = 0
		s.Resp.Token = s.token
		s.Resp.UserId = s.user.UserId
	}
}
