package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

func NewUserUseCase(repo UserRepo, logger log.Logger) *UserUseCase {
	return &UserUseCase{repo: repo, log: log.NewHelper(logger)}
}

type UserRepo interface {
	CreateUser(ctx context.Context, user User) error
	GetUserByID(ctx context.Context, userId string) (User, error)
	UpdateUser(ctx context.Context, user User) error
	DeleteUserByID(ctx context.Context, userId string) error
	GetUserList(ctx context.Context) (*[]User, error)
}

type User struct {
	ID       int32
	UserID   string
	UserName string
	State    string
}
type UserUseCase struct {
	repo UserRepo
	log  *log.Helper
}

func (u UserUseCase) CreateUser(ctx context.Context, user User) error {
	//TODO implement me
	return u.repo.CreateUser(ctx, user)
}

func (u UserUseCase) GetUserByID(ctx context.Context, userId string) (User, error) {
	return u.repo.GetUserByID(ctx, userId)
}

func (u UserUseCase) UpdateUser(ctx context.Context, user User) error {
	return u.repo.UpdateUser(ctx, user)
}

func (u UserUseCase) DeleteUserByID(ctx context.Context, userId string) error {
	return u.repo.DeleteUserByID(ctx, userId)
}

func (u UserUseCase) GetUserList(ctx context.Context) (*[]User, error) {
	return u.repo.GetUserList(ctx)
}
