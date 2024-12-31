package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"micro.mall/mall/order/api/internal/config"
	user "micro.mall/mall/user/rpc/userclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
