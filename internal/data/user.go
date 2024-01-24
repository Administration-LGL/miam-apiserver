package data

import (
	"context"
	"github.com/Administration-LGL/miam-apiserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// Create implements biz.UserRepo.
func (ur *userRepo) Create(ctx context.Context, user *biz.User, opts biz.CreateOptions) error {
	return ur.data.db.Create(user).Error
}

// Delete implements biz.UserRepo.
func (ur *userRepo) Delete(ctx context.Context, username string, opts biz.DeleteOptions) error {
	panic("unimplemented")
}

// DeleteCollection implements biz.UserRepo.
func (ur *userRepo) DeleteCollection(ctx context.Context, usernames []string, opts biz.DeleteOptions) error {
	panic("unimplemented")
}

// Get implements biz.UserRepo.
func (ur *userRepo) Get(ctx context.Context, user *biz.User, opts biz.GetOptions) (*biz.User, error) {
	u := &biz.User{}
	tx := ur.data.db
	if user.ID != 0 {
		tx = tx.Where("id = ?", user.ID)
	}
	if user.Name != "" {
		tx = tx.Where("name = ?", user.Name)
	}
	if user.Nickname != "" {
		tx = tx.Where("nickname like %?%", user.Nickname)
	}

	err := tx.First(u).Error
	return u, err
}

// List implements biz.UserRepo.
func (ur *userRepo) List(ctx context.Context, opts biz.ListOptions) (*biz.UserList, error) {
	panic("unimplemented")
}

// Update implements biz.UserRepo.
func (ur *userRepo) Update(ctx context.Context, user *biz.User, opts biz.UpdateOptions) error {
	panic("unimplemented")
}

// NewGreeterRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
