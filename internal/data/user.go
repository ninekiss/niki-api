package data

import (
	"context"
	"errors"
	"gorm.io/gorm"
	v1 "niki-api/gen/api/user/v1"
	"niki-api/internal/biz"
	"niki-api/utils"
	"strings"

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
		return nil, biz.ErrUserNotFound
	}

	// 密码解密
	err = utils.ParseRandPassword(user.Password, []byte(u.Password))
	if err != nil {
		// 自定义错误
		err := v1.ErrorCustomerError("密码错误")
		err = err.WithMetadata(map[string]string{
			"code":    "1001",
			"message": "密码错误",
		})
		return nil, err
	}

	return &user, nil
}

func (r *userRepo) Save(ctx context.Context, u *biz.User) (*biz.User, error) {
	// 判断用户名是否重复
	var user biz.User
	err := r.data.db.Select("id").Where("username = ?", u.Username).First(&user).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, biz.ErrUserAlreadyExists
	}

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

	err = r.data.db.Model(&biz.User{}).Create(u).Error
	if err != nil {
		return nil, err
	}
	return u, nil
}

func (r *userRepo) List(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	// 拼接查询条件
	qa := []string{"username LIKE ?", "nickname LIKE ?"}
	args := []interface{}{"%" + req.Username + "%", "%" + req.Nickname + "%"}

	if req.Phone != "" {
		qa = append(qa, "phone = ?")
		args = append(args, req.Phone)
	}
	if req.Email != "" {
		qa = append(qa, "email = ?")
		args = append(args, req.Email)
	}
	query := strings.Join(qa, " AND ")

	tx := r.data.db.Model(&biz.User{}).Where(query, args...)
	// 获取总数
	var total int64
	err := tx.Count(&total).Error
	if err != nil {
		return nil, err
	}
	// 获取分页数据
	var users []*biz.User
	err = tx.Limit(int(req.PageSize)).Offset(int(req.PageSize * (req.Page - 1))).Find(&users).Error
	if err != nil {
		return nil, err
	}

	// 填充数据
	rv := make([]*v1.ListUserReply_Data, 0, len(users))
	for _, user := range users {
		rv = append(rv, &v1.ListUserReply_Data{
			Uid:       user.Uid,
			Username:  user.Username,
			Nickname:  user.Nickname,
			Phone:     user.Phone,
			Email:     user.Email,
			Avatar:    user.Avatar,
			Status:    int32(user.Status),
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
			Gender:    int32(user.Gender),
			Age:       int32(user.Age),
		})
	}
	return &v1.ListUserReply{
		Total:    int32(total),
		Data:     rv,
		Page:     req.Page,
		PageSize: req.PageSize,
	}, nil
}

func (r *userRepo) Update(ctx context.Context, u *biz.User) error {
	_, err := r.GetAll(ctx, []string{u.Uid})
	if err != nil {
		return err
	}
	// 更新数据
	err = r.data.db.Model(&biz.User{}).Select("Gender").Where("uid = ?", u.Uid).Updates(u).Error
	return nil
}

func (r *userRepo) Delete(ctx context.Context, uids []string) error {
	_, err := r.GetAll(ctx, uids)
	if err != nil {
		return err
	}
	// 删除数据
	err = r.data.db.Model(&biz.User{}).Where("uid IN ?", uids).Delete(&biz.User{}).Error
	return nil
}

func (r *userRepo) GetAll(ctx context.Context, uids []string) ([]biz.User, error) {
	var user []biz.User
	err := r.data.db.Model(&biz.User{}).Where("uid IN ?", uids).Find(&user).Error
	if err != nil {
		return nil, biz.ErrUserNotFound
	}
	return user, nil
}
