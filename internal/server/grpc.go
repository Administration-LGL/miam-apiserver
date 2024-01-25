package server

import (
	authv1 "github.com/Administration-LGL/miam-apiserver/api/auth/v1"
	policyv1 "github.com/Administration-LGL/miam-apiserver/api/policy/v1"
	secretv1 "github.com/Administration-LGL/miam-apiserver/api/secret/v1"
	userv1 "github.com/Administration-LGL/miam-apiserver/api/user/v1"
	"github.com/Administration-LGL/miam-apiserver/internal/conf"
	"github.com/Administration-LGL/miam-apiserver/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
)

// NewGRPCServer new a gRPC server.
func NewGRPCServer(c *conf.Server, policy *service.PolicyService, user *service.UserService, secret *service.SecretService, logger log.Logger) *grpc.Server {
	var opts = []grpc.ServerOption{
		grpc.Middleware(
			recovery.Recovery(),
			// selector.Server(
			// 	jwt.Server(func(token *jwtv4.Token) (interface{}, error) {
			// 		return []byte(c.Jwt.Secret), nil
			// 	},
			// 		jwt.WithSigningMethod(jwtv4.SigningMethodHS256),
			// 	)).Match(NewWhiteListMatcher()).Build(),
			logging.Server(logger),
		),
	}
	if c.Grpc.Network != "" {
		opts = append(opts, grpc.Network(c.Grpc.Network))
	}
	if c.Grpc.Addr != "" {
		opts = append(opts, grpc.Address(c.Grpc.Addr))
	}
	if c.Grpc.Timeout != nil {
		opts = append(opts, grpc.Timeout(c.Grpc.Timeout.AsDuration()))
	}
	srv := grpc.NewServer(opts...)
	policyv1.RegisterPolicyServer(srv, policy)
	userv1.RegisterUserServer(srv, user)
	authv1.RegisterAuthServer(srv, user)
	secretv1.RegisterSecretServer(srv, secret)
	return srv
}
