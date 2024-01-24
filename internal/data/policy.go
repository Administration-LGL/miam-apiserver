package data

import (
	"context"
	"errors"
	"github.com/Administration-LGL/miam-apiserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type policyRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewPolicyRepo(data *Data, logger log.Logger) biz.PolicyRepo {
	return &policyRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (repo *policyRepo) Create(ctx context.Context, policy *biz.Policy, opts biz.CreateOptions) error {
	// 实现具体的 Create 逻辑
	return repo.data.db.Create(policy).Error
}

func (repo *policyRepo) Update(ctx context.Context, policy *biz.Policy, opts biz.UpdateOptions) error {
	// 实现具体的 Update 逻辑
	return errors.New("Not implemented")
}

func (repo *policyRepo) Delete(ctx context.Context, username string, name string, opts biz.DeleteOptions) error {
	// 实现具体的 Delete 逻辑
	return errors.New("Not implemented")
}

func (repo *policyRepo) DeleteCollection(ctx context.Context, username string, names []string, opts biz.DeleteOptions) error {
	// 实现具体的 DeleteCollection 逻辑
	return errors.New("Not implemented")
}

func (repo *policyRepo) Get(ctx context.Context, username string, name string, opts biz.GetOptions) (*biz.Policy, error) {
	// 实现具体的 Get 逻辑
	return nil, errors.New("Not implemented")
}

func (repo *policyRepo) List(ctx context.Context, username string, opts biz.ListOptions) (*biz.PolicyList, error) {
	// 实现具体的 List 逻辑
	return nil, errors.New("Not implemented")
}
