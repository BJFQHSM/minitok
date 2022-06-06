package service

import (
	"context"
	"errors"
	"regexp"
	"time"

	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal"
	"github.com/bytedance2022/minimal_tiktok/cmd/auth/dal/mysql"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"

	"github.com/bytedance2022/minimal_tiktok/grpc_gen/auth"
)

type LoginService interface {
	DoService() *auth.LoginResponse
}

func NewLoginService(ctx context.Context, r *auth.LoginRequest) LoginService {
	return &loginServiceImpl{Req: r, Ctx: ctx, Resp: &auth.LoginResponse{}}
}

type loginServiceImpl struct {
	Req  *auth.LoginRequest
	Resp *auth.LoginResponse
	Ctx  context.Context

	user  *mysql.User
	token string
}

func (s *loginServiceImpl) DoService() *auth.LoginResponse {
	// mock
	msg := "success"
	s.Resp = &auth.LoginResponse{
		UserId:     2,
		Token:      "fsjflsjdf",
		StatusCode: 0,
		StatusMsg:  &msg,
	}

	var err error
	for i := 0; i < 1; i++ {
		if err = s.validateParams(); err != nil {
			break
		}

		if err = s.doLogin(); err != nil {
			break
		}

		if err = s.generateToken(); err != nil {
			break
		}
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *loginServiceImpl) validateParams() error {
	req := s.Req
	if req == nil {
		return errors.New("params: request could not be nil")
	}
	if match, err := regexp.Match("^[\u4E00-\u9FA5A-Za-z\\d]{1,20}$", []byte(req.Username)); err != nil {
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

func (s *loginServiceImpl) doLogin() error {
	user, err := mysql.QueryUserByUsername(s.Ctx, s.Req.Username)
	if err != nil {
		return err
	}
	info, err := mysql.QueryEncryptInfoByUserId(user.UserId)
	if err != nil {
		return err
	}
	if util.MD5Encrypt(s.Req.Password, info.Salt) != user.EncryptPwd {
		return errors.New("please enter the correct password")
	}

	s.user = user
	return nil
}

func (s *loginServiceImpl) generateToken() error {
	token, err := dal.JwtGenerateToken(s.user, 24*time.Hour)
	if err != nil {
		return err
	}
	s.token = token
	return nil
}

func (s *loginServiceImpl) buildResponse(err error) {
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
