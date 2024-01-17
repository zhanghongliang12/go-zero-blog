// Code generated by goctl. DO NOT EDIT.
// Source: greet.proto

package greetclient

import (
	"context"

	"go-zero-blog/user/rpc/greet"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	Request  = greet.Request
	Response = greet.Response

	Greet interface {
		Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
	}

	defaultGreet struct {
		cli zrpc.Client
	}
)

func NewGreet(cli zrpc.Client) Greet {
	return &defaultGreet{
		cli: cli,
	}
}

func (m *defaultGreet) Ping(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	client := greet.NewGreetClient(m.cli.Conn())
	return client.Ping(ctx, in, opts...)
}
