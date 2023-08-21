package v1

import "github.com/gogf/gf/v2/frame/g"

type WebSocketReq struct {
	g.Meta `path:"/webSocket" tags:"Chat" summary:"You first hello api"`
}

type WebSocketRes struct {
	g.Meta `mime:"application/json"`
}

type PushChatToUsersReq struct {
	g.Meta `path:"/pushChatToUsers" tags:"Chat" summary:"You first hello api"`
}

type PushChatToUsersRes struct {
	g.Meta `mime:"application/json"`
}

type PushChatReq struct {
	g.Meta `path:"/pushChatToUsers" tags:"Chat" summary:"You first hello api"`
}

type PushChatRes struct {
	g.Meta `mime:"application/json"`
}
