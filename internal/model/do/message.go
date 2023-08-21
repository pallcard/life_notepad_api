// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Message is the golang structure of table t_message for DAO operations like Where/Data.
type Message struct {
	g.Meta     `orm:"table:t_message, do:true"`
	Id         interface{} //
	SenderId   interface{} // 发送者ID
	ReceiverId interface{} // 接收者ID
	Content    interface{} // 内容
	Link       interface{} // 是否链接 1是 2不是
	Unread     interface{} // 未读 1未读 2已读
	CreatedAt  *gtime.Time //
	UpdatedAt  *gtime.Time //
	DeletedAt  *gtime.Time //
}
