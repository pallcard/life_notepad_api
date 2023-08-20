package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type ChatListReq struct {
	g.Meta   `path:"/chatList" tags:"Chat" method:"post" summary:"You first hello api"`
	UserId   int `v:"required#UserId不能为空"`
	PageNum  int
	PageSize int
}

type ChatListRes struct {
	g.Meta   `mime:"application/json"`
	ChatList []ChatItem
	Total    int
}

type ChatItem struct {
	Id           int
	SenderId     int    //发送者
	SenderAvatar string //发送头像
	NickName     string //发送昵称
	Content      string //内容
	IsLiked      int    //是否链接 1是 2不是
	Unread       int    //未读 1未读 2已读
	CreateTime   string
}

type MessageListReq struct {
	g.Meta     `path:"/messageList" tags:"Chat" method:"post" summary:"You first hello api"`
	ReceiverId int `v:"required#ReceiverId不能为空"` //接收者
	SenderId   int `v:"required#SenderId不能为空"`   //发送者
	PageNum    int
	PageSize   int
}

type MessageListRes struct {
	g.Meta      `mime:"application/json"`
	MessageList []MessageItem
	Total       int
}

type MessageItem struct {
	Id         int
	SenderId   int    //发送者
	ReceiverId int    //接收者
	Content    string //内容
	IsLiked    int    //是否链接 1是 2不是
	Unread     int    //未读 1未读 2已读
	CreateTime string
}
