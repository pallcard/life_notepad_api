// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Chat is the golang structure of table t_chat for DAO operations like Where/Data.
type Chat struct {
	g.Meta     `orm:"table:t_chat, do:true"`
	Id         interface{} //
	SenderId   interface{} // 发送者ID,最新的一条
	ReceiverId interface{} // 接收者ID，逗号分割
	Content    interface{} // 内容，最新的一条
	Link       interface{} // 是否链接 1是 2不是
	Unread     interface{} // 未读 1未读 2已读
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
