package user

import (
	"context"
	user "github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user"

	"github.com/DavidWisdom/tiktok-backend/rpc_gen/kitex_gen/user/userservice"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
)

type RPCClient interface {
	KitexClient() userservice.Client
	Service() string
	Register(ctx context.Context, Req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error)
	Login(ctx context.Context, Req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *user.DouyinUserLoginResponse, err error)
	GetUser(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error)
}

func NewRPCClient(dstService string, opts ...client.Option) (RPCClient, error) {
	kitexClient, err := userservice.NewClient(dstService, opts...)
	if err != nil {
		return nil, err
	}
	cli := &clientImpl{
		service:     dstService,
		kitexClient: kitexClient,
	}

	return cli, nil
}

type clientImpl struct {
	service     string
	kitexClient userservice.Client
}

func (c *clientImpl) Service() string {
	return c.service
}

func (c *clientImpl) KitexClient() userservice.Client {
	return c.kitexClient
}

func (c *clientImpl) Register(ctx context.Context, Req *user.DouyinUserRegisterRequest, callOptions ...callopt.Option) (r *user.DouyinUserRegisterResponse, err error) {
	return c.kitexClient.Register(ctx, Req, callOptions...)
}

func (c *clientImpl) Login(ctx context.Context, Req *user.DouyinUserLoginRequest, callOptions ...callopt.Option) (r *user.DouyinUserLoginResponse, err error) {
	return c.kitexClient.Login(ctx, Req, callOptions...)
}

func (c *clientImpl) GetUser(ctx context.Context, Req *user.DouyinUserRequest, callOptions ...callopt.Option) (r *user.DouyinUserResponse, err error) {
	return c.kitexClient.GetUser(ctx, Req, callOptions...)
}