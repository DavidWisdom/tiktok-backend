package service

import (
	"context"
	"testing"
	user "github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user"
)

func TestGetUser_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetUserService(ctx)
	// init req and assert value

	req := &user.DouyinUserRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
