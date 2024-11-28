package service

import (
	"context"
	user "github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user"
)

type GetUserService struct {
	ctx context.Context
} // NewGetUserService new GetUserService
func NewGetUserService(ctx context.Context) *GetUserService {
	return &GetUserService{ctx: ctx}
}

// Run create note info
func (s *GetUserService) Run(req *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// Finish your business logic.

	return
}
