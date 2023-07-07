package util

import (
	v1 "helloworld/api/helloworld/v1"
	"helloworld/internal/biz"
)

type Transformer interface {
	TransformToGet(user *biz.User) *v1.GetUserReply_User
	TransformToList(user *biz.User) *v1.ListUserReply_User
}

type Transform struct{}

func (t Transform) TransformToGet(user *biz.User) *v1.GetUserReply_User {
	return &v1.GetUserReply_User{UserID: user.UserID, UserName: user.UserName, State: user.State}
}

func (t Transform) TransformToList(user *biz.User) *v1.ListUserReply_User {
	return &v1.ListUserReply_User{UserID: user.UserID, UserName: user.UserName, State: user.State}
}
