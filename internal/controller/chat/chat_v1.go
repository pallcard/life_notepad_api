package chat

import (
	"fmt"
	"github.com/gogf/gf/v2/container/garray"
	"github.com/gogf/gf/v2/encoding/ghtml"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gorilla/websocket"
	"life_notepad_api/internal/consts"
	"life_notepad_api/internal/model"
)

func (c *Controller) WebSocket(r *ghttp.Request) {
	msg := &model.ChatMsg{}

	// 初始化WebSocket请求
	var (
		ws  *ghttp.WebSocket
		err error
	)
	ws, err = r.WebSocket()
	if err != nil {
		g.Log().Error(r.Context(), err)
		return
	}

	userId := r.GetQuery("userId").String()
	if userId == "" {
		userId = r.Request.RemoteAddr
	}

	// 初始化时设置用户昵称为当前链接信息
	userIds.Add(userId)
	userWsMap.Set(userId, ws)

	// 初始化后向所有客户端发送上线消息
	c.writeUserListToClient()

	for {
		// 阻塞读取WS数据
		_, msgByte, err := ws.ReadMessage()
		if err != nil {
			// 如果失败，那么表示断开，这里清除用户信息
			// 为简化演示，这里不实现失败重连机制
			userIds.Remove(userId)
			userWsMap.Remove(userId)
			// 通知所有客户端当前用户已下线
			c.writeUserListToClient()
			break
		}
		// JSON参数解析
		if err := gjson.DecodeTo(msgByte, msg); err != nil {
			c.write(ws, model.ChatMsg{
				Type: "error",
				Data: "消息格式不正确: " + err.Error(),
				From: "",
			})
			continue
		}
		// 数据校验
		//if err := gvalid.CheckStruct(msg, nil); err != nil {
		//	c.write(ws, model.ChatMsg{
		//		Type: "error",
		//		Data: gerror.Current(err).Error(),
		//		From: "",
		//	})
		//	continue
		//}
		msg.From = userId

		// WS操作类型
		switch msg.Type {
		//单发消息
		case "sendTo":
			// 发送间隔检查
			intervalKey := fmt.Sprintf("%p", ws)
			if ok, _ := cache.SetIfNotExist(r.Context(), intervalKey, struct{}{}, consts.SendInterval); !ok {
				c.write(ws, model.ChatMsg{
					Type: "error",
					Data: "您的消息发送得过于频繁，请休息下再重试",
					From: "",
				})
				continue
			}
			// 有消息时，群发消息
			if msg.Data != nil {
				if err = c.writeToUser(ghtml.SpecialChars(msg.From),
					model.ChatMsg{
						Type: "send",
						Data: ghtml.SpecialChars(gconv.String(msg.Data)),
						From: ghtml.SpecialChars(msg.From),
					}); err != nil {
				}
			}
		// 发送消息
		case "send":
			// 发送间隔检查
			intervalKey := fmt.Sprintf("%p", ws)
			if ok, _ := cache.SetIfNotExist(r.Context(), intervalKey, struct{}{}, consts.SendInterval); !ok {
				c.write(ws, model.ChatMsg{
					Type: "error",
					Data: "您的消息发送得过于频繁，请休息下再重试",
					From: "",
				})
				continue
			}
			// 有消息时，群发消息
			if msg.Data != nil {
				if err = c.writeGroup(
					model.ChatMsg{
						Type: "send",
						Data: ghtml.SpecialChars(gconv.String(msg.Data)),
						From: ghtml.SpecialChars(msg.From),
					}); err != nil {
					//g.Log().Error(err)
				}
			}
		}
	}
}

// PushChatToUsers http向指定用户发消息
func (a *Controller) PushChatToUsers(r *ghttp.Request) {
	var name string = "Runoob"
	var (
		err error
	)
	msg := &model.ChatMsg{}
	msg.Data = r.Get("data").String()
	msg.From = r.Get("from").String()
	name = r.Get("name").String()
	fmt.Print(msg)
	if name != "" && userWsMap.Get(name) != nil {
		// 有消息时，群发消息
		if msg.Data != nil {
			if err = a.writeToUser(name,
				model.ChatMsg{
					Type: "send",
					Data: ghtml.SpecialChars(gconv.String(msg.Data)),
					From: ghtml.SpecialChars(msg.From),
				}); err != nil {
				//g.Log().Error(err)
			}
		}
	} else {
		r.Response.Write("{abc:cdb}")
	}

}

// http群发消息
func (a *Controller) PushChat(r *ghttp.Request) {
	fmt.Print(a)
	fmt.Print(r)
	var (
		err error
	)
	msg := &model.ChatMsg{}
	msg.Data = r.Get("name").String()
	msg.From = r.Get("from").String()
	fmt.Print(msg)
	// 有消息时，群发消息
	if msg.Data != nil {
		if err = a.writeGroup(
			model.ChatMsg{
				Type: "send",
				Data: ghtml.SpecialChars(gconv.String(msg.Data)),
				From: ghtml.SpecialChars(msg.From),
			}); err != nil {
			//g.Log().Error(err)
		}
	}
}

// 向客户端写入消息。
// 内部方法不会自动注册到路由中。
func (a *Controller) write(ws *ghttp.WebSocket, msg model.ChatMsg) error {
	msgBytes, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	return ws.WriteMessage(websocket.TextMessage, msgBytes)
}

// 向指定用户名客户端群发消息。
// 内部方法不会自动注册到路由中。
func (a *Controller) writeToUser(user string, msg model.ChatMsg) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	userWsMap.Get(user).(*ghttp.WebSocket).WriteMessage(websocket.TextMessage, []byte(b))
	return nil
}

// 向所有客户端群发消息。
// 内部方法不会自动注册到路由中。
func (a *Controller) writeGroup(msg model.ChatMsg) error {
	b, err := gjson.Encode(msg)
	if err != nil {
		return err
	}
	//users.RLockFunc(func(m map[interface{}]interface{}) {
	userWsMap.RLockFunc(func(m map[interface{}]interface{}) {
		for user := range m {
			fmt.Print(user)
			m[user].(*ghttp.WebSocket).WriteMessage(websocket.TextMessage, []byte(b))
		}
	})
	return nil
}

// 向客户端返回用户列表。
// 内部方法不会自动注册到路由中。
func (a *Controller) writeUserListToClient() error {
	array := garray.NewSortedStrArray()
	userWsMap.Iterator(func(k interface{}, v interface{}) bool {
		array.Add(gconv.String(k))
		return true
	})
	if err := a.writeGroup(model.ChatMsg{
		Type: "list",
		Data: array.Slice(),
		From: "",
	}); err != nil {
		return err
	}
	return nil
}
