package logic

import (
	"context"
	"go-zero-blog/demo/rpc/demo/demo"

	"go-zero-blog/demo/api/demo/internal/svc"
	"go-zero-blog/demo/api/demo/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DemoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDemoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DemoLogic {
	return &DemoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DemoLogic) Demo(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	l.svcCtx.DemoRpc.Ping(l.ctx, &demo.Request{
		Ping: req.Name,
	})
	return resp, nil
}
