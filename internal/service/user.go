package service

import (
	"context"
	"fmt"
	pb "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
	"helloworld/internal/util"
)

type UserService struct {
	pb.UnimplementedUserServer
	uc *biz.UserUseCase
}

func NewUserService(uc *biz.UserUseCase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserReply, error) {
	err := s.uc.CreateUser(ctx, biz.User{UserID: req.UserID, UserName: req.UserName, State: req.State})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Msg:  "success",
		Code: "200",
	}, nil
}
func (s *UserService) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserReply, error) {
	fmt.Println(req.UserID)
	err := s.uc.UpdateUser(ctx, biz.User{UserID: req.UserID, UserName: req.UserName, State: req.State})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateUserReply{Msg: "success", Code: "200"}, nil
}
func (s *UserService) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserReply, error) {
	fmt.Println(req.UserID)
	err := s.uc.DeleteUserByID(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}
	return &pb.DeleteUserReply{
		Msg:  "success",
		Code: "200",
	}, nil
}
func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserReply, error) {
	user, err := s.uc.GetUserByID(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Msg:  "success",
		Code: "200",
		User: util.Transform{}.TransformToGet(&user),
	}, nil
}
func (s *UserService) ListUser(ctx context.Context, req *pb.ListUserRequest) (*pb.ListUserReply, error) {
	userList, err := s.uc.GetUserList(ctx)
	if err != nil {
		return nil, err
	}
	var listResult = make([]*pb.ListUserReply_User, 0)
	for _, user := range *userList {
		listResult = append(listResult, util.Transform{}.TransformToList(&user))
	}
	return &pb.ListUserReply{
		User: listResult,
		Msg:  "success",
		Code: "200",
	}, nil
}
