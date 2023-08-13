package service

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"niki-api/internal/biz"
	"niki-api/utils"

	v1 "niki-api/gen/api/user/v1"
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
	fmt.Println("UpdateUser", req.Id, req.Name, req.Email, req.Password)
	return &v1.UpdateUserReply{
		Id: req.Id,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *v1.DeleteUserRequest) (*v1.DeleteUserReply, error) {
	fmt.Println("DeleteUser", req.Id)
	return &v1.DeleteUserReply{
		Id: req.Id,
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *v1.GetUserRequest) (*v1.GetUserReply, error) {
	fmt.Println("GetUser", req.Id)
	return &v1.GetUserReply{
		Id:       req.Id,
		Name:     "test",
		Email:    "test@test.com",
		Password: "123456",
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *v1.ListUserRequest) (*v1.ListUserReply, error) {
	fmt.Println("ListUser", req.Page, req.PageSize, req.Sort, req.Order)
	return &v1.ListUserReply{
		Data: []*v1.ListUserReply_Data{
			{
				Id:       "1",
				Name:     "test1",
				Email:    "test@test.com",
				Password: "123456",
			},
			{
				Id:       "2",
				Name:     "test2",
				Email:    "test@test.com",
				Password: "123456",
			},
		},
		Page:     1,
		PageSize: 10,
		Total:    2,
	}, nil
}
