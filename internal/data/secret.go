package data

import (
	"context"
	"github.com/Administration-LGL/miam-apiserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type secretRepo struct {
	data *Data
	log  *log.Helper
}

// Create implements biz.SecretRepo.
func (repo *secretRepo) Create(ctx context.Context, secret biz.Secret, opts biz.CreateOptions) error {
	return repo.data.db.Create(&secret).Error
}

// Delete implements biz.SecretRepo.
func (*secretRepo) Delete(ctx context.Context, username string, secretID string, opts biz.DeleteOptions) error {
	panic("unimplemented")
}

// DeleteCollection implements biz.SecretRepo.
func (*secretRepo) DeleteCollection(ctx context.Context, username string, secretIDs []string, opts biz.DeleteOptions) error {
	panic("unimplemented")
}

// Get implements biz.SecretRepo.
func (*secretRepo) Get(ctx context.Context, username string, secretID string, opts biz.GetOptions) (biz.Secret, error) {
	panic("unimplemented")
}

// List implements biz.SecretRepo.
func (repo *secretRepo) List(ctx context.Context, username string, opts biz.ListOptions) (*biz.SecretList, error) {
	tx := repo.data.db
	if username != "" {
		tx = tx.Where("username=?", username)
	}
	//
	var secrets []*biz.Secret
	var count int64
	err := tx.Find(&secrets).Count(&count).Error
	return &biz.SecretList{
		Items:    secrets,
		ListMeta: biz.ListMeta{count},
	}, err
}

// Update implements biz.SecretRepo.
func (*secretRepo) Update(ctx context.Context, secret biz.Secret, opts biz.UpdateOptions) error {
	panic("unimplemented")
}

// NewGreeterRepo .
func NewSecretRepo(data *Data, logger log.Logger) biz.SecretRepo {
	return &secretRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
