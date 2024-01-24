package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Administration-LGL/miam-apiserver/pkg/util/idutil"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/ory/ladon"
	"gorm.io/gorm"
)

type Policy struct {
	// May add TypeMeta in the future.
	// TypeMeta `json:",inline"`

	// Standard object's metadata.
	ObjectMeta `json:"metadata,omitempty"`

	// The user of the policy.
	Username string `json:"username" gorm:"column:username" validate:"omitempty"`

	// AuthzPolicy policy, will not be stored in db.
	Policy AuthzPolicy `json:"policy,omitempty" gorm:"-" validate:"omitempty"`
	// Policy ladon.DefaultPolicy `json:"policy,omitempty" gorm:"-" validate:"omitempty"`

	// The ladon policy content, just a string format of ladon.DefaultPolicy. DO NOT modify directly.
	PolicyShadow string `json:"-" gorm:"column:policyShadow" validate:"omitempty"`
}

// AuthzPolicy defines iam policy type.
type AuthzPolicy struct {
	ladon.DefaultPolicy
}

// PolicyList is the whole list of all policies which have been stored in stroage.
type PolicyList struct {
	// Standard list metadata.
	ListMeta `json:",inline"`

	// List of policies.
	Items []*Policy `json:"items"`
}

// TableName maps to mysql table name.
func (p *Policy) TableName() string {
	return "policy"
}

// String returns the string format of Policy.
func (ap AuthzPolicy) String() string {
	data, _ := json.Marshal(ap)

	return string(data)
}

// BeforeCreate run before create database record.
func (p *Policy) BeforeCreate(tx *gorm.DB) error {
	if err := p.ObjectMeta.BeforeCreate(tx); err != nil {
		return fmt.Errorf("failed to run `BeforeCreate` hook: %w", err)
	}

	p.Policy.ID = p.Name
	p.PolicyShadow = p.Policy.String()

	return nil
}

// AfterCreate run after create database record.
func (p *Policy) AfterCreate(tx *gorm.DB) error {
	p.InstanceID = idutil.GetInstanceID(p.ID, "policy-")

	return tx.Save(p).Error
}

// BeforeUpdate run before update database record.
func (p *Policy) BeforeUpdate(tx *gorm.DB) error {
	if err := p.ObjectMeta.BeforeUpdate(tx); err != nil {
		return fmt.Errorf("failed to run `BeforeUpdate` hook: %w", err)
	}

	p.Policy.ID = p.Name
	p.PolicyShadow = p.Policy.String()

	return nil
}

// AfterFind run after find to unmarshal a policy string into ladon.DefaultPolicy struct.
func (p *Policy) AfterFind(tx *gorm.DB) error {
	if err := p.ObjectMeta.AfterFind(tx); err != nil {
		return fmt.Errorf("failed to run `AfterFind` hook: %w", err)
	}

	if err := json.Unmarshal([]byte(p.PolicyShadow), &p.Policy); err != nil {
		return fmt.Errorf("failed to unmarshal policyShadow: %w", err)
	}

	return nil
}

// PolicyRepo is a Greater repo.
type PolicyRepo interface {
	Create(ctx context.Context, policy *Policy, opts CreateOptions) error
	Update(ctx context.Context, policy *Policy, opts UpdateOptions) error
	Delete(ctx context.Context, username string, name string, opts DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, names []string, opts DeleteOptions) error
	Get(ctx context.Context, username string, name string, opts GetOptions) (*Policy, error)
	List(ctx context.Context, username string, opts ListOptions) (*PolicyList, error)
}

// PolicyUsecase is a Policy usecase.
type PolicyUsecase struct {
	repo PolicyRepo
	log  *log.Helper
}

func NewPolicyUsecase(repo PolicyRepo, logger log.Logger) *PolicyUsecase {
	return &PolicyUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (p *PolicyUsecase) Create(ctx context.Context, policy *Policy) error {
	return p.repo.Create(ctx, policy, CreateOptions{})
}
