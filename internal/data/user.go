package data

import (
	"context"

	"niki-api/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *userRepo) SignIn(ctx context.Context, u *biz.User) (*biz.User, error) {
	err := r.data.db.Model(&biz.User{}).Where("username = ? AND password = ?", u.Username, u.Password).First(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	// TODO: 密码加密
	u.Password = u.Password + ":hash:123"

	// TODO: 生成 uid
	u.Uid = "123456789"

	// TODO: 默认头像
	u.Avatar = "https://avatars.githubusercontent.com/u/56598311?s=60&v=4"

	err := r.data.db.Model(&biz.User{}).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
