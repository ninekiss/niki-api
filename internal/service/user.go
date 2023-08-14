package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	v1 "niki-api/gen/api/user/v1"
	"niki-api/internal/biz"
	"niki-api/utils"
)

type UserService struct {
	v1.UnimplementedUserServer
	uc  *biz.UserUsecase
	log *log.Helper
}

func NewUserService(uc *biz.UserUsecase, logger log.Logger) *UserService {
	return &UserService{
		uc:  uc,
		log: log.NewHelper(log.With(logger, "module", "service/server-service")),
	}
}

func (s *UserService) SignIn(ctx context.Context, req *v1.SignInRequest) (*v1.SignInReply, error) {
	// TODO: 校验验证码
	if req.Captcha != "123456" {
		return nil, v1.ErrorCustomerError("验证码错误")
	}

	rv, err := s.uc.SignIn(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	// 生成 token
	token, err := utils.GenToken(rv.Uid, rv.Username)
	if err != nil {
		return nil, err
	}

	return &v1.SignInReply{
		Uid:      rv.Uid,
		Username: rv.Username,
		Token:    token,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *v1.CreateUserRequest) (*v1.CreateUserReply, error) {
	rv, err := s.uc.CreateUser(ctx, &biz.User{
		Username: req.Username,
		Password: req.Password,
		Email:    req.Email,
	})
	if err != nil {
		return nil, err
	}
	return &v1.CreateUserReply{
		Uid: rv.Uid,
	}, nil
}

func (s *UserService) UpdateUser(ctx context.Context, req *v1.UpdateUserRequest) (*v1.UpdateUserReply, error) {
	fmt.Println(req.Gender, byte(req.Gender))
	err := s.uc.UpdateUser(ctx, &biz.User{
		Uid:      req.Uid,
		Username: req.Username,
		Nickname: req.Nickname,
		Email:    req.Email,
		Phone:    req.Phone,
		Age:      byte(req.Age),
		Gender:   byte(req.Gender),
		Avatar:   req.Avatar,
		Status:   byte(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return &v1.UpdateUserReply{}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	err := s.uc.DeleteUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	return &v1.DeleteUserReply{}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	users, err := s.uc.GetUser(ctx, req.Uid)
	if err != nil {
		return nil, err
	}
	// 转换
	var rv v1.GetUserReply
	for _, user := range users {
		rv.Uid = user.Uid
		rv.Username = user.Username
		rv.Nickname = user.Nickname
		rv.Email = user.Email
		rv.Phone = user.Phone
		rv.Age = int32(user.Age)
		rv.Gender = int32(user.Gender)
		rv.Avatar = user.Avatar
		rv.Status = int32(user.Status)
		rv.CreatedAt = user.CreatedAt.Format("2006-01-02 15:04:05")
		rv.UpdatedAt = user.UpdatedAt.Format("2006-01-02 15:04:05")
	}
	return &rv, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	rv, err := s.uc.ListUser(ctx, req)
	if err != nil {
		return nil, err
	}
	return rv, nil
}
