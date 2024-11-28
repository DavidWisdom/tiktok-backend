package user

import (
	"context"
	user "github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/pkg/klog"
)

func Register(ctx context.Context, req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (resp *user.DouyinUserRegisterResponse, err error) {
	resp, err = defaultClient.Register(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Register call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func Login(ctx context.Context, req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (resp *user.DouyinUserLoginResponse, err error) {
	resp, err = defaultClient.Login(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "Login call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}

func GetUser(ctx context.Context, req *user.DouyinUserRequest, callOptions ...callopt.Option) (resp *user.DouyinUserResponse, err error) {
	resp, err = defaultClient.GetUser(ctx, req, callOptions...)
	if err != nil {
		klog.CtxErrorf(ctx, "GetUser call failed,err =%+v", err)
		return nil, err
	}
	return resp, nil
}
