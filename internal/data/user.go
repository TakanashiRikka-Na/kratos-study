package data

import (
	"context"
	"encoding/json"
	"github.com/go-kratos/kratos/v2/log"
	"helloworld/internal/biz"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u userRepo) CreateUser(ctx context.Context, user biz.User) error {
	if err := u.data.DB.Table("users").Create(&user).Error; err != nil {
		u.log.Error(err)
		return err
	}
	return nil
}

func (u userRepo) GetUserByID(ctx context.Context, userId string) (user biz.User, err error) {

	result := u.data.Rdb.Get(userId).Val()
	//redis 里没有 从 数据库里取
	if len(result) != 0 {
		user, err = u.UnMarshalUserFromCache(result)
		if err != nil {
			u.log.Error(err)
			return biz.User{}, err
		}
		return
	}
	err = u.data.DB.Table("users").Where("user_id=?", userId).Find(&user).Error
	if err != nil {
		u.log.Error(err)
		return biz.User{}, err
	}
	if err = u.SetUserToCache(user); err != nil {
		u.log.Error(err)
	}
	return
}

func (u userRepo) UpdateUser(ctx context.Context, user biz.User) error {
	err := u.data.DB.Table("users").Where("user_name=?", user.UserName).Updates(&user).Error
	if err != nil {
		u.log.Error(err)
		return err
	}
	return nil
}

func (u userRepo) DeleteUserByID(ctx context.Context, userId string) error {
	err := u.data.DB.Table("users").Where("user_id=?", userId).Delete(&biz.User{}).Error
	if err != nil {
		u.log.Error(err)
		return err
	}
	return nil
}

func (u userRepo) GetUserList(ctx context.Context) (*[]biz.User, error) {
	var users = make([]biz.User, 0)
	err := u.data.DB.Table("users").Find(&users).Error
	if err != nil {
		u.log.Error(err)
		return nil, err
	}
	return &users, nil
}

func (u userRepo) SetUserToCache(user biz.User) (err error) {
	userMarshal, err := json.Marshal(&user)
	if err != nil {
		return
	}
	if err = u.data.Rdb.Set(user.UserID, string(userMarshal), 0).Err(); err != nil {
		return
	}
	return
}

func (u userRepo) DeleteUserCache(user biz.User) error {
	return u.data.Rdb.Del(user.UserID).Err()
}

func (u userRepo) UnMarshalUserFromCache(userMarshal string) (biz.User, error) {
	var user biz.User
	if err := json.Unmarshal([]byte(userMarshal), &user); err != nil {
		return biz.User{}, err
	}
	return user, nil
}
