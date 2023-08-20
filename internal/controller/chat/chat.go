package chat

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/container/gset"
	"github.com/gogf/gf/v2/os/gcache"
	"life_notepad_api/api/chat"
)

type Controller struct{}

func NewV1() chat.IChatV1 {
	return &Controller{}
}

var (
	userWsMap = gmap.New(true)       // 使用默认的并发安全Map
	userIds   = gset.NewStrSet(true) // 使用并发安全的Set，用以用户昵称唯一性校验
	cache     = gcache.New()         // 使用特定的缓存对象，不使用全局缓存对象
)
