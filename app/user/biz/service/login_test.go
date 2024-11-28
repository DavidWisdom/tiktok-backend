package service

import (
	"context"
	"testing"
	user "github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.DouyinUserLoginRequest{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}
