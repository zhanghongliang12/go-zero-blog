package user

import (
	"context"

	"go-zero-blog/user/api/internal/svc"
	"go-zero-blog/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type Hello_worldLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewHello_worldLogic(ctx context.Context, svcCtx *svc.ServiceContext) *Hello_worldLogic {
	return &Hello_worldLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *Hello_worldLogic) Hello_world(req *types.TestReq) (resp *types.TestResp, err error) {
	// todo: add your logic here and delete this line
	resp = &types.TestResp{
		Info: "测试",
	}
	return resp, nil
}
