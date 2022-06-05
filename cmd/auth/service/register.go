package service

import (
	"context"
	"errors"
	"github.com/bytedance2022/minimal_tiktok/pkg/util"
	"regexp"

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

		// todo
	}
	s.buildResponse(err)
	return s.Resp
}

func (s *registerServiceImpl) validateParams() error {
	req := s.Req
	if len(req.Username) < 5 && len(req.Password) < 8 {
		return errors.New("request params length illegal")
	}
	if match, err := regexp.Match("^[\\u4E00-\\u9FA5A-Za-z\\d]+$", []byte(req.Username)); err != nil {
		return errors.New("fail to validate username")
	} else if !match {
		return errors.New("username can only contains nums, English letters and Chinese")
	}
	if match, err := regexp.Match("^[A-Za-z\\d]+$", []byte(req.Password)); err != nil {
		return errors.New("fail to validate password")
	} else if !match {
		return errors.New("password can only contains nums, English letters and underline character")
	}

	return nil
}

func (s *registerServiceImpl) doRegister() {

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
	}
}
