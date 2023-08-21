package consts

import "time"

const (
	// SendInterval 允许客户端发送聊天消息的间隔时间
	SendInterval = time.Second

	ChatTypeSingle = "single"
	ChatTypeGroup  = "group"
	ChatTypeError  = "error"
)
