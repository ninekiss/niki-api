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
	Uid      string `gorm:"<-:create;type: char(20);not null;uniqueIndex" json:"uid"`
	Username string `gorm:"type: char(20);not null;uniqueIndex" json:"username" form:"username"`
	Password string `gorm:"type: char(100);not null" json:"password" form:"password"`
	Nickname string `gorm:"type: char(20);comment: 昵称" json:"nickname" form:"nickname"`
	Email    string `gorm:"type: char(20);comment: 邮箱" json:"email" form:"email"`
	Phone    string `gorm:"type: char(20);comment: 手机" json:"phone" form:"phone"`
	Age      byte   `gorm:"type: int;comment: 年龄" json:"age" form:"age"`
	Gender   byte   `gorm:"type: int;default: 2;comment: 性别,0-男,1-女,2-未知;" json:"gender" form:"gender"`
	Address  string `gorm:"type: char(200);comment: 地址" json:"address" form:"address"`
	Avatar   string `gorm:"column:image" json:"avatar"`
	Status   byte   `gorm:"type: int;default: 1;comment: 状态,0-禁用,1-启用" json:"status" form:"status"`
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
