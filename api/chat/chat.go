// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package chat

import "github.com/gogf/gf/v2/net/ghttp"

type IChatV1 interface {
	WebSocket(r *ghttp.Request)
	PushChatToUsers(r *ghttp.Request)
	PushChat(r *ghttp.Request)
}
