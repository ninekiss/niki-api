package data

import (
	"context"
	"errors"
	"gorm.io/gorm"
	"niki-api/utils"

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
	var user biz.User
	err := r.data.db.Model(&biz.User{}).Where("username = ?", u.Username).First(&user).Error
	if err != nil {
		return nil, err
	}

	// 密码解密
	err = utils.ParseRandPassword(user.Password, []byte(u.Password))
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	// 密码加密
	salt, _ := utils.RandSalt(utils.SaltRange, 6)
	hashed, _ := utils.CryptPassword([]byte(u.Password), salt)
	u.Password = utils.RandHashedPassword(hashed, salt)

	// 生成 uid
	var uid string
	var digits int = 9 // 8 位 uid
	uid = utils.GenerateRandID(digits)

	// 如果 uid 重复, 则重新生成 uid, 直到 uid 不重复
	for {
		var user biz.User
		err := r.data.db.Select("id").Where("uid = ?", uid).First(&user).Error
		if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		if user.ID < 1 {
			break
		} else {
			uid = utils.GenerateRandID(digits)
		}
	}
	u.Uid = uid

	// TODO: 默认头像
	u.Avatar = "https://avatars.githubusercontent.com/u/56598311?s=60&v=4"

	err := r.data.db.Model(&biz.User{}).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}
