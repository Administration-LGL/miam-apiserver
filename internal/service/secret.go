package service

import (
	"context"

	secretv1 "github.com/Administration-LGL/miam-apiserver/api/secret/v1"
	"github.com/Administration-LGL/miam-apiserver/internal/biz"
	"github.com/Administration-LGL/miam-apiserver/internal/pkg/util"
)

// GreeterService is a greeter service.
type SecretService struct {
	uc *biz.SecretUsecase
	secretv1.UnimplementedSecretServer
}

// NewGreeterService new a greeter service.
func NewSecretService(uc *biz.SecretUsecase) *SecretService {
	return &SecretService{uc: uc}
}

func (srv *SecretService) CreateSecret(ctx context.Context, req *secretv1.CreateSecretRequest) (*secretv1.CreateSecretReply, error) {
	return &secretv1.CreateSecretReply{}, srv.uc.Create(ctx, biz.Secret{Expires: req.Expires})
}
func (srv *SecretService) ListSecret(ctx context.Context, req *secretv1.ListSecretRequest) (*secretv1.ListSecretReply, error) {
	list, err := srv.uc.List(ctx, req.Username, biz.ListOptions{})
	if err != nil {
		return &secretv1.ListSecretReply{}, err
	}
	return util.SecretList2SecretListReply(list), nil
}
