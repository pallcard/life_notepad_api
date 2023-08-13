// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Note is the golang structure of table t_note for DAO operations like Where/Data.
type Note struct {
	g.Meta    `orm:"table:t_note, do:true"`
	Id        interface{} //
	UserId    interface{} // 用户ID
	Content   interface{} // 内容
	Images    interface{} // 图片
	Location  interface{} // 头像
	CreatedAt *gtime.Time //
	UpdatedAt *gtime.Time //
	DeletedAt *gtime.Time //
}
