package biz

import (
	"context"
	"miam-apiserver/middleware/auth/jwt"
	"miam-apiserver/pkg/util/idutil"

	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

// Secret represents a secret restful resource.
// It is also used as gorm model.
type Secret struct {
	// May add TypeMeta in the future.
	// TypeMeta `json:",inline"`

	// Standard object's metadata.
	ObjectMeta `       json:"metadata,omitempty"`
	Username   string `json:"username"           gorm:"column:username"  validate:"omitempty"`
	//nolint: tagliatelle
	SecretID  string `json:"secretID"           gorm:"column:secretID"  validate:"omitempty"`
	SecretKey string `json:"secretKey"          gorm:"column:secretKey" validate:"omitempty"`

	// Required: true
	Expires     int64  `json:"expires"     gorm:"column:expires"     validate:"omitempty"`
	Description string `json:"description" gorm:"column:description" validate:"description"`
}

// SecretList is the whole list of all secrets which have been stored in stroage.
type SecretList struct {
	// May add TypeMeta in the future.
	// TypeMeta `json:",inline"`

	// Standard list metadata.
	ListMeta `json:",inline"`

	// List of secrets
	Items []*Secret `json:"items"`
}

// TableName maps to mysql table name.
func (s *Secret) TableName() string {
	return "secret"
}

// AfterCreate run after create database record.
func (s *Secret) AfterCreate(tx *gorm.DB) error {
	s.InstanceID = idutil.GetInstanceID(s.ID, "secret-")

	return tx.Save(s).Error
}

type SecretRepo interface {
	Create(ctx context.Context, secret Secret, opts CreateOptions) error
	Update(ctx context.Context, secret Secret, opts UpdateOptions) error
	Delete(ctx context.Context, username, secretID string, opts DeleteOptions) error
	DeleteCollection(ctx context.Context, username string, secretIDs []string, opts DeleteOptions) error
	Get(ctx context.Context, username, secretID string, opts GetOptions) (Secret, error)
	List(ctx context.Context, username string, opts ListOptions) (*SecretList, error)
}

// PolicyUsecase is a Policy usecase.
type SecretUsecase struct {
	repo SecretRepo
	log  *log.Helper
}

func NewSecretUsecase(repo SecretRepo, logger log.Logger) *SecretUsecase {
	return &SecretUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *SecretUsecase) List(ctx context.Context, opts ListOptions) (*SecretList, error) {
	username, err := jwt.GetUsername(ctx)
	if err != nil {
		return &SecretList{}, err
	}
	return uc.repo.List(ctx, username, opts)
}

func (uc *SecretUsecase) Create(ctx context.Context, secret Secret) error {
	username, err := jwt.GetUsername(ctx)
	if err != nil {
		return err
	}
	secret.Username = username
	secret.SecretKey = idutil.NewSecretKey()
	secret.SecretID = idutil.NewSecretID()
	return uc.repo.Create(ctx, secret, CreateOptions{})
}
