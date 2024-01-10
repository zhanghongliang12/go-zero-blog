// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	user "go-zero-blog/user/api/internal/handler/user"
	"go-zero-blog/user/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user/register",
				Handler: user.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/user/hello_world",
				Handler: user.Hello_worldHandler(serverCtx),
			},
		},
		rest.WithPrefix("/user/v1"),
	)
}
