// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// User is the golang structure of table t_user for DAO operations like Where/Data.
type User struct {
	g.Meta      `orm:"table:t_user, do:true"`
	Id          interface{} //
	Email       interface{} // 邮箱
	Password    interface{} // 密码
	Avatar      interface{} // 头像
	NickName    interface{} // 昵称
	Description interface{} // 描述
	CreatedAt   *gtime.Time //
	UpdatedAt   *gtime.Time //
	DeletedAt   *gtime.Time //
}
