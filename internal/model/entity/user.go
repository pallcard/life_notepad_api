// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure for table user.
type User struct {
	Id          int         `json:"id"          description:""`
	Email       string      `json:"email"       description:"邮箱"`
	Password    string      `json:"password"    description:"密码"`
	Avatar      string      `json:"avatar"      description:"头像"`
	NickName    string      `json:"nickName"    description:"昵称"`
	Description string      `json:"description" description:"描述"`
	CreatedAt   *gtime.Time `json:"createdAt"   description:""`
	UpdatedAt   *gtime.Time `json:"updatedAt"   description:""`
	DeletedAt   *gtime.Time `json:"deletedAt"   description:""`
}
