// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Chat is the golang structure for table chat.
type Chat struct {
	Id         int         `json:"id"         description:""`
	SenderId   int         `json:"senderId"   description:"发送者ID,最新的一条"`
	ReceiverId int         `json:"receiverId" description:"接收者ID，逗号分割"`
	Content    string      `json:"content"    description:"内容，最新的一条"`
	Link       int         `json:"link"       description:"是否链接 1是 2不是"`
	Unread     int         `json:"unread"     description:"未读 1未读 2已读"`
	CreatedAt  *gtime.Time `json:"createdAt"  description:""`
	UpdatedAt  *gtime.Time `json:"updatedAt"  description:""`
	DeletedAt  *gtime.Time `json:"deletedAt"  description:""`
}
