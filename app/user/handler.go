package main

import (
	"context"
	user "github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user"
	"github.com/DavidWisdom/tiktok-backend/app/user/biz/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, req *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = service.NewRegisterService(ctx).Run(req)

	return resp, err
}

// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, req *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	resp, err = service.NewLoginService(ctx).Run(req)

	return resp, err
}

// GetUser implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUser(ctx context.Context, req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	resp, err = service.NewGetUserService(ctx).Run(req)

	return resp, err
}
