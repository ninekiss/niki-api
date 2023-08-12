package service

import (
	"context"
	"fmt"

	pb "niki-api/gen/api/user/v1"
)

type UserService struct {
	pb.UnimplementedUserServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	fmt.Println("CreateUser", req.Name, req.Email, req.Password)
	return &pb.CreateUserReply{
		Id: "123",
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	fmt.Println("UpdateUser", req.Id, req.Name, req.Email, req.Password)
	return &pb.UpdateUserReply{
		Id: req.Id,
	}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	fmt.Println("DeleteUser", req.Id)
	return &pb.DeleteUserReply{
		Id: req.Id,
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	fmt.Println("GetUser", req.Id)
	return &pb.GetUserReply{
		Id:       req.Id,
		Name:     "test",
		Email:    "test@test.com",
		Password: "123456",
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	fmt.Println("ListUser", req.Page, req.PageSize, req.Sort, req.Order)
	return &pb.ListUserReply{
		Data: []*pb.ListUserReply_Data{
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
