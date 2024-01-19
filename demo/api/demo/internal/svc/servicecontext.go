package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"go-zero-blog/demo/api/demo/internal/config"
	"go-zero-blog/demo/rpc/demo/democlient"
)

type ServiceContext struct {
	Config  config.Config
	DemoRpc democlient.Demo
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		DemoRpc: democlient.NewDemo(zrpc.MustNewClient(c.DemoRpcConf)),
	}
}
