package service

import (
	"context"

	authpb "github.com/Administration-LGL/miam-apiserver/api/auth/v1"
	userpb "github.com/Administration-LGL/miam-apiserver/api/user/v1"
	"github.com/Administration-LGL/miam-apiserver/internal/biz"
	"github.com/Administration-LGL/miam-apiserver/internal/pkg/util"

	"github.com/go-kratos/kratos/v2/log"
)

// GreeterService is a greeter service.
type UserService struct {
	uc *biz.UserUsecase
	userpb.UnimplementedUserServer
	authpb.UnimplementedAuthServer
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (us *UserService) Create(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserReply, error) {
	// 参数合法性校验

	// 转换
	user := util.CreateUserReq2BizUser(req)
	if err := us.uc.Create(ctx, user); err != nil {
		return &userpb.CreateUserReply{}, err
	} else {
		return &userpb.CreateUserReply{
			Id: user.ID,
		}, nil
	}
}

func (us *UserService) Get(ctx context.Context, req *userpb.GetUserRequest) (*userpb.UserReply, error) {
	user := util.GetUserReq2BizUser(req)
	if user, err := us.uc.Get(ctx, user); err != nil {
		return &userpb.UserReply{}, err
	} else {
		return util.BizUser2UserReply(user), nil
	}
}

func (us *UserService) Login(ctx context.Context, req *authpb.LoginRequest) (*authpb.LoginReply, error) {
	log.Debug(req)
	if token, err := us.uc.Login(ctx, req.Username, req.Password); err != nil {
		return &authpb.LoginReply{}, err
	} else {
		return &authpb.LoginReply{Token: token}, nil
	}
}
