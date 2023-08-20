// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package message

import (
	"context"
	v1 "life_notepad_api/api/message/v1"
)

type IMessageV1 interface {
	ChatList(ctx context.Context, req *v1.ChatListReq) (res *v1.ChatListRes, err error)
	MessageList(ctx context.Context, req *v1.MessageListReq) (res *v1.MessageListRes, err error)
}
