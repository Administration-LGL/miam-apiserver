package biz

import (
	"context"
	"errors"
	"github.com/Administration-LGL/miam-apiserver/internal/conf"
	"github.com/Administration-LGL/miam-apiserver/pkg/encrypt"
	"github.com/Administration-LGL/miam-apiserver/pkg/util/idutil"
	"time"

	mjwt "github.com/Administration-LGL/miam-apiserver/middleware/auth/jwt"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/golang-jwt/jwt/v4"
	"gorm.io/gorm"
)

type User struct {
	ObjectMeta

	Status int `json:"status" gorm:"column:status" validate:"omitempty"`

	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname" validate:"required,min=1,max=30"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password" validate:"required"`

	// Required: true
	Email string `json:"email" gorm:"column:email" validate:"required,email,min=1,max=100"`

	Phone string `json:"phone" gorm:"column:phone" validate:"omitempty"`

	IsAdmin int `json:"isAdmin,omitempty" gorm:"column:isAdmin" validate:"omitempty"`

	TotalPolicy int64 `json:"totalPolicy" gorm:"-" validate:"omitempty"`

	LoginedAt time.Time `json:"loginedAt,omitempty" gorm:"column:loginedAt"`
}

// TableName maps to mysql table name.
func (u *User) TableName() string {
	return "user"
}

// Compare with the plain text password. Returns true if it's the same as the encrypted one (in the `User` struct).
func (u *User) Compare(pwd string) error {
	// if err := auth.Compare(u.Password, pwd); err != nil {
	// 	return fmt.Errorf("failed to compile password: %w", err)
	// }
	// 比较密码
	return encrypt.ComparePwd(u.Password, pwd)
}

// UserList is the whole list of all users which have been stored in stroage.
type UserList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	ListMeta `json:",inline"`

	Items []*User `json:"items"`
}

// PolicyRepo is a Greater repo.
type UserRepo interface {
	Create(ctx context.Context, user *User, opts CreateOptions) error
	Update(ctx context.Context, user *User, opts UpdateOptions) error
	Delete(ctx context.Context, username string, opts DeleteOptions) error
	DeleteCollection(ctx context.Context, usernames []string, opts DeleteOptions) error
	Get(ctx context.Context, user *User, opts GetOptions) (*User, error)
	List(ctx context.Context, opts ListOptions) (*UserList, error)
}

// PolicyUsecase is a Policy usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
	jwt  *conf.Server_JWT
}

func NewUserUsecase(repo UserRepo, logger log.Logger, conf *conf.Server) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger), jwt: conf.GetJwt()}
}

func (uuc *UserUsecase) Create(ctx context.Context, u *User) error {
	var err error
	// 密码加密
	u.Password, err = encrypt.EncryptPwd(u.Password)
	if err != nil {
		return err
	}
	// 登录时间更新
	u.LoginedAt = time.Now()

	u.Status = 1
	return uuc.repo.Create(ctx, u, CreateOptions{})
}

func (uuc *UserUsecase) Login(ctx context.Context, username, password string) (string, error) {
	// 通过用户名读取用户
	user, err := uuc.repo.Get(ctx, &User{ObjectMeta: ObjectMeta{Name: username}}, GetOptions{})
	if err != nil || user == nil {
		return "", errors.New(err.Error())
	}
	if err := user.Compare(password); err != nil {
		return "", err
	}
	// 密码校验通过，签发token
	claims := mjwt.AuthClaims{}
	// 写入token过期时间
	expire := time.Now().Add(time.Duration(uuc.jwt.Timeout) * time.Minute)
	claims.ExpiresAt = &jwt.NumericDate{expire}
	// 设置用户名
	claims.Username = user.Name
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err := token.SignedString([]byte(uuc.jwt.Secret))
	if err != nil {
		return "", err
	}
	if serverContext, ok := transport.FromServerContext(ctx); ok {
		serverContext.ReplyHeader().Add("cookie", tokenStr)
	}
	// 重定向至主页
	return tokenStr, nil
}

func (uuc *UserUsecase) Get(ctx context.Context, u *User) (*User, error) {
	return uuc.repo.Get(ctx, u, GetOptions{})
}

// AfterCreate run after create database record.
func (u *User) AfterCreate(tx *gorm.DB) error {
	u.InstanceID = idutil.GetInstanceID(u.ID, "user-")

	return tx.Save(u).Error
}
