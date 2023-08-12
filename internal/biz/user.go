package biz

import (
	"context"
	"gorm.io/gorm"
	v1 "niki-api/gen/api/helloworld/v1"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(v1.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User .
type User struct {
	gorm.Model
	Uid      string `gorm:"column:uid;primary_key" json:"uid"`
	Username string `gorm:"column:username;unique" json:"username"`
	Email    string `gorm:"column:email;unique" json:"email"`
	Password string `gorm:"column:password" json:"password"`
	Bio      string `gorm:"column:bio" json:"bio"`
	Avatar   string `gorm:"column:image" json:"avatar"`
}

// UserRepo .
type UserRepo interface {
	SignIn(context.Context, *User) (*User, error)
	Save(context.Context, *User) (*User, error)
	//Update(context.Context, *User) (*User, error)
	//FindByID(context.Context, int64) (*User, error)
	//ListByHello(context.Context, string) ([]*User, error)
	//ListAll(context.Context) ([]*User, error)
}

type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *UserUsecase) SignIn(ctx context.Context, u *User) (*User, error) {
	return uc.repo.SignIn(ctx, u)
}

func (uc *UserUsecase) CreateUser(ctx context.Context, u *User) (*User, error) {
	return uc.repo.Save(ctx, u)
}
