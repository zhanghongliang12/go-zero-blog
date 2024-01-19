package svc

import (
	"go-zero-blog/user/api/internal/config"
	"go-zero-blog/user/rpc/greetclient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc greetclient.Greet
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
