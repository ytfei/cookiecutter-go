package router

import (
	"github.com/tal-tech/go-zero/rest"
	"net/http"

	"{{cookiecutter.app_relative_path}}/{{cookiecutter.app_name}}/internal/service/admin"
)

// admin API server routes:
func RegisterAdminRoutes(engine *rest.Server, svc *admin.Service) {
	// register:
	engine.AddRoutes(
		[]rest.Route{
			// rpc api test:
			{
				Method: http.MethodGet,
				Path:   "/admin/api/rpc/test", // query block
				Handler: HandlerWrap(func(r *http.Request) (resp interface{}, err error) {
					return svc.AdminCall(r)
				}),
			},
		},
	)
}
