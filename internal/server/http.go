package server

import (
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/swagger-api/openapiv2"
	"niki-api/gen/api/conf"
	helloworld "niki-api/gen/api/helloworld/v1"
	user "niki-api/gen/api/user/v1"
	"niki-api/internal/service"
	"niki-api/middleware"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, gs *service.GreeterService, us *service.UserService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			validate.Validator(),
			selector.Server(middleware.JWTMiddleware()).Path().Build(),
		),
		http.Filter(handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"Content-Type", "X-Requested-With", "Authorization", "User-Agent"}),
		)),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	srv := http.NewServer(opts...)
	openAPIhandler := openapiv2.NewHandler()
	srv.HandlePrefix("/q/", openAPIhandler)

	helloworld.RegisterGreeterHTTPServer(srv, gs)
	user.RegisterUserHTTPServer(srv, us)
	return srv
}
