// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package user

import (
	"context"
	v1 "life_notepad_api/api/user/v1"
)

type IUserV1 interface {
	User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	UpdateUser(ctx context.Context, req *v1.UpdateUserReq) (res *v1.UpdateUserRes, err error)
}
