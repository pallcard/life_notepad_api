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
}
