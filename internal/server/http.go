package server

import (
	"context"
	authv1 "miam-apiserver/api/auth/v1"
	policyv1 "miam-apiserver/api/policy/v1"
	userv1 "miam-apiserver/api/user/v1"
	secretv1 "miam-apiserver/api/secret/v1"

	"miam-apiserver/internal/conf"
	"miam-apiserver/internal/service"
	"miam-apiserver/middleware/auth/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	jwtv4 "github.com/golang-jwt/jwt/v4"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, policy *service.PolicyService, user *service.UserService,secret *service.SecretService, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			recovery.Recovery(),
			selector.Server(
				jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
					return []byte(c.Jwt.Secret), nil
				},
					jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
				)).Match(NewWhiteListMatcher()).Build(),
			logging.Server(logger),
		),
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
	policyv1.RegisterPolicyHTTPServer(srv, policy)
	userv1.RegisterUserHTTPServer(srv, user)
	authv1.RegisterAuthHTTPServer(srv, user)
	secretv1.RegisterSecretHTTPServer(srv, secret)
	return srv
}

func NewWhiteListMatcher() selector.MatchFunc {

	whiteList := make(map[string]struct{})
	// whiteList["/shop.interface.v1.ShopInterface/Login"] = struct{}{}
	whiteList["/auth.v1.Auth/Login"] = struct{}{}
	whiteList["/user.v1.User/Create"] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whiteList[operation]; ok {
			return false
		}
		return true
	}
}
